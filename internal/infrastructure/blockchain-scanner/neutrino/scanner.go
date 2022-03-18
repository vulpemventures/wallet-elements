package neutrino_scanner

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	log "github.com/sirupsen/logrus"
	"github.com/vulpemventures/go-elements/confidential"
	"github.com/vulpemventures/go-elements/elementsutil"
	"github.com/vulpemventures/neutrino-elements/pkg/blockservice"
	"github.com/vulpemventures/neutrino-elements/pkg/repository"
	"github.com/vulpemventures/neutrino-elements/pkg/scanner"
	"github.com/vulpemventures/ocean/internal/core/domain"
)

type scannerService struct {
	accountName  string
	svc          scanner.ScannerService
	blindingKeys map[string][]byte
	chTxs        chan *domain.Transaction
	chUtxos      chan []*domain.Utxo
	lock         *sync.RWMutex

	log  func(format string, a ...interface{})
	warn func(err error, format string, a ...interface{})
}

func newScannerSvc(
	accountName string,
	filtersDb repository.FilterRepository,
	headersDb repository.BlockHeaderRepository,
	blockSvc blockservice.BlockService, genesisHash *chainhash.Hash,
) *scannerService {
	logFn := func(format string, a ...interface{}) {
		format = fmt.Sprintf("scanner: %s", format)
		log.Debugf(format, a...)
	}
	warnFn := func(err error, format string, a ...interface{}) {
		format = fmt.Sprintf("scanner: %s", format)
		log.WithError(err).Warnf(format, a...)
	}
	scannerSvc := &scannerService{
		accountName:  accountName,
		svc:          scanner.New(filtersDb, headersDb, blockSvc, genesisHash),
		blindingKeys: make(map[string][]byte),
		chTxs:        make(chan *domain.Transaction, 10),
		chUtxos:      make(chan []*domain.Utxo, 10),
		lock:         &sync.RWMutex{},
		log:          logFn,
		warn:         warnFn,
	}
	chReports, _ := scannerSvc.svc.Start()
	go scannerSvc.listenToReports(chReports)
	return scannerSvc
}

func (s *scannerService) stop() {
	s.svc.Stop()
	close(s.chTxs)
	close(s.chUtxos)
}

func (s *scannerService) watchAddresses(addressesInfo []domain.AddressInfo) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, info := range addressesInfo {
		// Prevent duplicates
		if _, ok := s.blindingKeys[info.Script]; ok {
			continue
		}

		s.blindingKeys[info.Script] = info.BlindingKey
		item, _ := scanner.NewScriptWatchItemFromAddress(info.Address)
		s.svc.Watch(scanner.WithWatchItem(item))
		s.log(
			"start watching address %s for account %s",
			info.DerivationPath, s.accountName,
		)
	}
}

func (s *scannerService) listenToReports(chReports <-chan scanner.Report) {
	s.log("start listening to incoming reports from node")
	for r := range chReports {
		tx := r.Transaction
		txid := tx.TxHash().String()
		txHex, _ := tx.ToHex()

		s.log("received report for tx %s", txid)

		s.chTxs <- &domain.Transaction{
			TxID:  txid,
			TxHex: txHex,
			Accounts: map[string]struct{}{
				s.accountName: {},
			},
			BlockHash:   r.BlockHash.String(),
			BlockHeight: r.BlockHeight,
		}

		spentUtxos := make([]*domain.Utxo, 0, len(tx.Inputs))
		for _, in := range tx.Inputs {
			spentUtxos = append(spentUtxos, &domain.Utxo{
				UtxoKey: domain.UtxoKey{
					TxID: elementsutil.TxIDFromBytes(in.Hash),
					VOut: in.Index,
				},
				Spent: true,
			})
		}
		s.chUtxos <- spentUtxos

		newUtxos := make([]*domain.Utxo, 0)
		for i, out := range tx.Outputs {
			if len(out.Script) == 0 {
				continue
			}

			script := hex.EncodeToString(out.Script)
			blindingKey, ok := s.getBlindingKey(script)
			if !ok {
				continue
			}

			revealed, err := confidential.UnblindOutputWithKey(out, blindingKey)
			if err != nil {
				s.warn(err, "failed to unblind utxo with given blinding key")
				continue
			}

			newUtxos = append(newUtxos, &domain.Utxo{
				UtxoKey: domain.UtxoKey{
					TxID: txid,
					VOut: uint32(i),
				},
				Value:           revealed.Value,
				Asset:           assetFromBytes(revealed.Asset),
				ValueCommitment: out.Value,
				AssetCommitment: out.Asset,
				ValueBlinder:    revealed.ValueBlindingFactor,
				AssetBlinder:    revealed.AssetBlindingFactor,
				Script:          out.Script,
				Nonce:           out.Nonce,
				RangeProof:      out.RangeProof,
				SurjectionProof: out.SurjectionProof,
				AccountName:     s.accountName,
				Confirmed:       true,
			})
		}

		if len(newUtxos) > 0 {
			s.chUtxos <- newUtxos
		}
	}
}

func (s *scannerService) getBlindingKey(script string) ([]byte, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	key, ok := s.blindingKeys[script]
	return key, ok
}

func assetFromBytes(buf []byte) string {
	return hex.EncodeToString(elementsutil.ReverseBytes(buf))
}