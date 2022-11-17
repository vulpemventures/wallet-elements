// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account WHERE name = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, deleteAccount, name)
	return err
}

const deleteAccountScripts = `-- name: DeleteAccountScripts :exec
DELETE FROM account_script_info WHERE fk_account_name = $1
`

func (q *Queries) DeleteAccountScripts(ctx context.Context, fkAccountName string) error {
	_, err := q.db.Exec(ctx, deleteAccountScripts, fkAccountName)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT name, index, xpub, derivation_path, next_external_index, next_internal_index, fk_wallet_id FROM account WHERE name = $1
`

func (q *Queries) GetAccount(ctx context.Context, name string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, name)
	var i Account
	err := row.Scan(
		&i.Name,
		&i.Index,
		&i.Xpub,
		&i.DerivationPath,
		&i.NextExternalIndex,
		&i.NextInternalIndex,
		&i.FkWalletID,
	)
	return i, err
}

const getWalletAccountsAndScripts = `-- name: GetWalletAccountsAndScripts :many
SELECT w.id as walletId,w.encrypted_mnemonic,w.password_hash,w.birthday_block_height,w.root_path,w.network_name,w.next_account_index, a.name,a.index,a.xpub,a.derivation_path as account_derivation_path,a.next_external_index,a.next_internal_index,a.fk_wallet_id,asi.script,asi.derivation_path as script_derivation_path,asi.fk_account_name FROM
wallet w LEFT JOIN account a ON w.id = a.fk_wallet_id
LEFT JOIN account_script_info asi on a.name = asi.fk_account_name
WHERE w.id = $1
`

type GetWalletAccountsAndScriptsRow struct {
	Walletid              string
	EncryptedMnemonic     []byte
	PasswordHash          []byte
	BirthdayBlockHeight   int32
	RootPath              string
	NetworkName           string
	NextAccountIndex      int32
	Name                  sql.NullString
	Index                 sql.NullInt32
	Xpub                  sql.NullString
	AccountDerivationPath sql.NullString
	NextExternalIndex     sql.NullInt32
	NextInternalIndex     sql.NullInt32
	FkWalletID            sql.NullString
	Script                sql.NullString
	ScriptDerivationPath  sql.NullString
	FkAccountName         sql.NullString
}

func (q *Queries) GetWalletAccountsAndScripts(ctx context.Context, id string) ([]GetWalletAccountsAndScriptsRow, error) {
	rows, err := q.db.Query(ctx, getWalletAccountsAndScripts, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWalletAccountsAndScriptsRow
	for rows.Next() {
		var i GetWalletAccountsAndScriptsRow
		if err := rows.Scan(
			&i.Walletid,
			&i.EncryptedMnemonic,
			&i.PasswordHash,
			&i.BirthdayBlockHeight,
			&i.RootPath,
			&i.NetworkName,
			&i.NextAccountIndex,
			&i.Name,
			&i.Index,
			&i.Xpub,
			&i.AccountDerivationPath,
			&i.NextExternalIndex,
			&i.NextInternalIndex,
			&i.FkWalletID,
			&i.Script,
			&i.ScriptDerivationPath,
			&i.FkAccountName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAccount = `-- name: InsertAccount :one
INSERT INTO account(name,index,xpub,derivation_path,next_external_index,next_internal_index,fk_wallet_id)
VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING name, index, xpub, derivation_path, next_external_index, next_internal_index, fk_wallet_id
`

type InsertAccountParams struct {
	Name              string
	Index             int32
	Xpub              string
	DerivationPath    string
	NextExternalIndex int32
	NextInternalIndex int32
	FkWalletID        string
}

func (q *Queries) InsertAccount(ctx context.Context, arg InsertAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, insertAccount,
		arg.Name,
		arg.Index,
		arg.Xpub,
		arg.DerivationPath,
		arg.NextExternalIndex,
		arg.NextInternalIndex,
		arg.FkWalletID,
	)
	var i Account
	err := row.Scan(
		&i.Name,
		&i.Index,
		&i.Xpub,
		&i.DerivationPath,
		&i.NextExternalIndex,
		&i.NextInternalIndex,
		&i.FkWalletID,
	)
	return i, err
}

type InsertAccountScriptsParams struct {
	Script         string
	DerivationPath string
	FkAccountName  string
}

const insertUtxo = `-- name: InsertUtxo :one
INSERT INTO utxo(tx_id,vout,value,asset,value_commitment,asset_commitment,value_blinder,asset_blinder,script,nonce,range_proof,surjection_proof,account_name,lock_timestamp)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, nonce, range_proof, surjection_proof, account_name, lock_timestamp
`

type InsertUtxoParams struct {
	TxID            string
	Vout            int32
	Value           int32
	Asset           string
	ValueCommitment []byte
	AssetCommitment []byte
	ValueBlinder    []byte
	AssetBlinder    []byte
	Script          []byte
	Nonce           []byte
	RangeProof      []byte
	SurjectionProof []byte
	AccountName     string
	LockTimestamp   time.Time
}

// UTXO
func (q *Queries) InsertUtxo(ctx context.Context, arg InsertUtxoParams) (Utxo, error) {
	row := q.db.QueryRow(ctx, insertUtxo,
		arg.TxID,
		arg.Vout,
		arg.Value,
		arg.Asset,
		arg.ValueCommitment,
		arg.AssetCommitment,
		arg.ValueBlinder,
		arg.AssetBlinder,
		arg.Script,
		arg.Nonce,
		arg.RangeProof,
		arg.SurjectionProof,
		arg.AccountName,
		arg.LockTimestamp,
	)
	var i Utxo
	err := row.Scan(
		&i.ID,
		&i.TxID,
		&i.Vout,
		&i.Value,
		&i.Asset,
		&i.ValueCommitment,
		&i.AssetCommitment,
		&i.ValueBlinder,
		&i.AssetBlinder,
		&i.Script,
		&i.Nonce,
		&i.RangeProof,
		&i.SurjectionProof,
		&i.AccountName,
		&i.LockTimestamp,
	)
	return i, err
}

const insertWallet = `-- name: InsertWallet :one
INSERT INTO wallet(id, encrypted_mnemonic,password_hash,birthday_block_height,root_path,network_name,next_account_index)
VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, encrypted_mnemonic, password_hash, birthday_block_height, root_path, network_name, next_account_index
`

type InsertWalletParams struct {
	ID                  string
	EncryptedMnemonic   []byte
	PasswordHash        []byte
	BirthdayBlockHeight int32
	RootPath            string
	NetworkName         string
	NextAccountIndex    int32
}

// WALLET & ACCOUNT
func (q *Queries) InsertWallet(ctx context.Context, arg InsertWalletParams) (Wallet, error) {
	row := q.db.QueryRow(ctx, insertWallet,
		arg.ID,
		arg.EncryptedMnemonic,
		arg.PasswordHash,
		arg.BirthdayBlockHeight,
		arg.RootPath,
		arg.NetworkName,
		arg.NextAccountIndex,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.EncryptedMnemonic,
		&i.PasswordHash,
		&i.BirthdayBlockHeight,
		&i.RootPath,
		&i.NetworkName,
		&i.NextAccountIndex,
	)
	return i, err
}

const updateAccountIndexes = `-- name: UpdateAccountIndexes :one
UPDATE account SET next_external_index = $1, next_internal_index = $2 WHERE name = $3 RETURNING name, index, xpub, derivation_path, next_external_index, next_internal_index, fk_wallet_id
`

type UpdateAccountIndexesParams struct {
	NextExternalIndex int32
	NextInternalIndex int32
	Name              string
}

func (q *Queries) UpdateAccountIndexes(ctx context.Context, arg UpdateAccountIndexesParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccountIndexes, arg.NextExternalIndex, arg.NextInternalIndex, arg.Name)
	var i Account
	err := row.Scan(
		&i.Name,
		&i.Index,
		&i.Xpub,
		&i.DerivationPath,
		&i.NextExternalIndex,
		&i.NextInternalIndex,
		&i.FkWalletID,
	)
	return i, err
}

const updateWallet = `-- name: UpdateWallet :one
UPDATE wallet SET encrypted_mnemonic = $2, password_hash = $3, birthday_block_height = $4, root_path = $5, network_name = $6, next_account_index = $7 WHERE id = $1 RETURNING id, encrypted_mnemonic, password_hash, birthday_block_height, root_path, network_name, next_account_index
`

type UpdateWalletParams struct {
	ID                  string
	EncryptedMnemonic   []byte
	PasswordHash        []byte
	BirthdayBlockHeight int32
	RootPath            string
	NetworkName         string
	NextAccountIndex    int32
}

func (q *Queries) UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error) {
	row := q.db.QueryRow(ctx, updateWallet,
		arg.ID,
		arg.EncryptedMnemonic,
		arg.PasswordHash,
		arg.BirthdayBlockHeight,
		arg.RootPath,
		arg.NetworkName,
		arg.NextAccountIndex,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.EncryptedMnemonic,
		&i.PasswordHash,
		&i.BirthdayBlockHeight,
		&i.RootPath,
		&i.NetworkName,
		&i.NextAccountIndex,
	)
	return i, err
}