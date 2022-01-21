# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ocean/v1alpha/transaction.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ocean.v1alpha import types_pb2 as ocean_dot_v1alpha_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1focean/v1alpha/transaction.proto\x12\rocean.v1alpha\x1a\x19ocean/v1alpha/types.proto\"\xb8\x02\n\x12SelectUtxosRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12!\n\x0ctarget_asset\x18\x02 \x01(\tR\x0btargetAsset\x12#\n\rtarget_amount\x18\x03 \x01(\x04R\x0ctargetAmount\x12\x46\n\x08strategy\x18\x04 \x01(\x0e\x32*.ocean.v1alpha.SelectUtxosRequest.StrategyR\x08strategy\"V\n\x08Strategy\x12\x18\n\x14STRATEGY_UNSPECIFIED\x10\x00\x12\x19\n\x15STRATEGY_BRANCH_BOUND\x10\x01\x12\x15\n\x11STRATEGY_FRAGMENT\x10\x02\"X\n\x13SelectUtxosResponse\x12)\n\x05utxos\x18\x01 \x03(\x0b\x32\x13.ocean.v1alpha.UtxoR\x05utxos\x12\x16\n\x06\x63hange\x18\x02 \x01(\x04R\x06\x63hange\"t\n\x13\x45stimateFeesRequest\x12,\n\x06inputs\x18\x01 \x03(\x0b\x32\x14.ocean.v1alpha.InputR\x06inputs\x12/\n\x07outputs\x18\x02 \x03(\x0b\x32\x15.ocean.v1alpha.OutputR\x07outputs\"5\n\x14\x45stimateFeesResponse\x12\x1d\n\nfee_amount\x18\x01 \x01(\x04R\tfeeAmount\"/\n\x16SignTransactionRequest\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"0\n\x17SignTransactionResponse\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"4\n\x1b\x42roadcastTransactionRequest\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"2\n\x1c\x42roadcastTransactionResponse\x12\x12\n\x04txid\x18\x01 \x01(\tR\x04txid\"r\n\x11\x43reatePsetRequest\x12,\n\x06inputs\x18\x01 \x03(\x0b\x32\x14.ocean.v1alpha.InputR\x06inputs\x12/\n\x07outputs\x18\x02 \x03(\x0b\x32\x15.ocean.v1alpha.OutputR\x07outputs\"(\n\x12\x43reatePsetResponse\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\"\x86\x01\n\x11UpdatePsetRequest\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\x12,\n\x06inputs\x18\x02 \x03(\x0b\x32\x14.ocean.v1alpha.InputR\x06inputs\x12/\n\x07outputs\x18\x03 \x03(\x0b\x32\x15.ocean.v1alpha.OutputR\x07outputs\"(\n\x12UpdatePsetResponse\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\"&\n\x10\x42lindPsetRequest\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\"\'\n\x11\x42lindPsetResponse\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\"%\n\x0fSignPsetRequest\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\"&\n\x10SignPsetResponse\x12\x12\n\x04pset\x18\x01 \x01(\tR\x04pset\"\xf4\x01\n\x0bMintRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12!\n\x0c\x61sset_amount\x18\x02 \x01(\x04R\x0b\x61ssetAmount\x12!\n\x0ctoken_amount\x18\x03 \x01(\x04R\x0btokenAmount\x12\x1d\n\nasset_name\x18\x04 \x01(\tR\tassetName\x12!\n\x0c\x61sset_ticker\x18\x05 \x01(\tR\x0b\x61ssetTicker\x12!\n\x0c\x61sset_domain\x18\x06 \x01(\tR\x0b\x61ssetDomain\"%\n\x0cMintResponse\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"y\n\rRemintRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12\x14\n\x05\x61sset\x18\x02 \x01(\tR\x05\x61sset\x12\x16\n\x06\x61mount\x18\x03 \x01(\x04R\x06\x61mount\"\'\n\x0eRemintResponse\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"~\n\x0b\x42urnRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12\x33\n\treceivers\x18\x02 \x03(\x0b\x32\x15.ocean.v1alpha.OutputR\treceivers\"%\n\x0c\x42urnResponse\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"\x82\x01\n\x0fTransferRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12\x33\n\treceivers\x18\x02 \x03(\x0b\x32\x15.ocean.v1alpha.OutputR\treceivers\")\n\x10TransferResponse\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex\"\x15\n\x13PegInAddressRequest\"\xa3\x01\n\x14PegInAddressResponse\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12,\n\x12main_chain_address\x18\x02 \x01(\tR\x10mainChainAddress\x12!\n\x0c\x63laim_script\x18\x03 \x01(\tR\x0b\x63laimScript\"w\n\x11\x43laimPegInRequest\x12\x1d\n\nbitcoin_tx\x18\x01 \x01(\tR\tbitcoinTx\x12 \n\x0ctx_out_proof\x18\x02 \x01(\tR\ntxOutProof\x12!\n\x0c\x63laim_script\x18\x03 \x01(\tR\x0b\x63laimScript\"+\n\x12\x43laimPegInResponse\x12\x15\n\x06tx_hex\x18\x01 \x01(\tR\x05txHex2\x9b\t\n\x12TransactionService\x12T\n\x0bSelectUtxos\x12!.ocean.v1alpha.SelectUtxosRequest\x1a\".ocean.v1alpha.SelectUtxosResponse\x12W\n\x0c\x45stimateFees\x12\".ocean.v1alpha.EstimateFeesRequest\x1a#.ocean.v1alpha.EstimateFeesResponse\x12`\n\x0fSignTransaction\x12%.ocean.v1alpha.SignTransactionRequest\x1a&.ocean.v1alpha.SignTransactionResponse\x12o\n\x14\x42roadcastTransaction\x12*.ocean.v1alpha.BroadcastTransactionRequest\x1a+.ocean.v1alpha.BroadcastTransactionResponse\x12Q\n\nCreatePset\x12 .ocean.v1alpha.CreatePsetRequest\x1a!.ocean.v1alpha.CreatePsetResponse\x12Q\n\nUpdatePset\x12 .ocean.v1alpha.UpdatePsetRequest\x1a!.ocean.v1alpha.UpdatePsetResponse\x12N\n\tBlindPset\x12\x1f.ocean.v1alpha.BlindPsetRequest\x1a .ocean.v1alpha.BlindPsetResponse\x12K\n\x08SignPset\x12\x1e.ocean.v1alpha.SignPsetRequest\x1a\x1f.ocean.v1alpha.SignPsetResponse\x12?\n\x04Mint\x12\x1a.ocean.v1alpha.MintRequest\x1a\x1b.ocean.v1alpha.MintResponse\x12\x45\n\x06Remint\x12\x1c.ocean.v1alpha.RemintRequest\x1a\x1d.ocean.v1alpha.RemintResponse\x12?\n\x04\x42urn\x12\x1a.ocean.v1alpha.BurnRequest\x1a\x1b.ocean.v1alpha.BurnResponse\x12K\n\x08Transfer\x12\x1e.ocean.v1alpha.TransferRequest\x1a\x1f.ocean.v1alpha.TransferResponse\x12W\n\x0cPegInAddress\x12\".ocean.v1alpha.PegInAddressRequest\x1a#.ocean.v1alpha.PegInAddressResponse\x12Q\n\nClaimPegIn\x12 .ocean.v1alpha.ClaimPegInRequest\x1a!.ocean.v1alpha.ClaimPegInResponseB\xdd\x01\n\x11\x63om.ocean.v1alphaB\x10TransactionProtoP\x01Zagithub.com/vulpemventures/ocean/api-spec/protobuf/ocean/v1alpha/gen/go/ocean/v1alpha;oceanv1alpha\xa2\x02\x03OXX\xaa\x02\rOcean.V1alpha\xca\x02\rOcean\\V1alpha\xe2\x02\x19Ocean\\V1alpha\\GPBMetadata\xea\x02\x0eOcean::V1alphab\x06proto3')



_SELECTUTXOSREQUEST = DESCRIPTOR.message_types_by_name['SelectUtxosRequest']
_SELECTUTXOSRESPONSE = DESCRIPTOR.message_types_by_name['SelectUtxosResponse']
_ESTIMATEFEESREQUEST = DESCRIPTOR.message_types_by_name['EstimateFeesRequest']
_ESTIMATEFEESRESPONSE = DESCRIPTOR.message_types_by_name['EstimateFeesResponse']
_SIGNTRANSACTIONREQUEST = DESCRIPTOR.message_types_by_name['SignTransactionRequest']
_SIGNTRANSACTIONRESPONSE = DESCRIPTOR.message_types_by_name['SignTransactionResponse']
_BROADCASTTRANSACTIONREQUEST = DESCRIPTOR.message_types_by_name['BroadcastTransactionRequest']
_BROADCASTTRANSACTIONRESPONSE = DESCRIPTOR.message_types_by_name['BroadcastTransactionResponse']
_CREATEPSETREQUEST = DESCRIPTOR.message_types_by_name['CreatePsetRequest']
_CREATEPSETRESPONSE = DESCRIPTOR.message_types_by_name['CreatePsetResponse']
_UPDATEPSETREQUEST = DESCRIPTOR.message_types_by_name['UpdatePsetRequest']
_UPDATEPSETRESPONSE = DESCRIPTOR.message_types_by_name['UpdatePsetResponse']
_BLINDPSETREQUEST = DESCRIPTOR.message_types_by_name['BlindPsetRequest']
_BLINDPSETRESPONSE = DESCRIPTOR.message_types_by_name['BlindPsetResponse']
_SIGNPSETREQUEST = DESCRIPTOR.message_types_by_name['SignPsetRequest']
_SIGNPSETRESPONSE = DESCRIPTOR.message_types_by_name['SignPsetResponse']
_MINTREQUEST = DESCRIPTOR.message_types_by_name['MintRequest']
_MINTRESPONSE = DESCRIPTOR.message_types_by_name['MintResponse']
_REMINTREQUEST = DESCRIPTOR.message_types_by_name['RemintRequest']
_REMINTRESPONSE = DESCRIPTOR.message_types_by_name['RemintResponse']
_BURNREQUEST = DESCRIPTOR.message_types_by_name['BurnRequest']
_BURNRESPONSE = DESCRIPTOR.message_types_by_name['BurnResponse']
_TRANSFERREQUEST = DESCRIPTOR.message_types_by_name['TransferRequest']
_TRANSFERRESPONSE = DESCRIPTOR.message_types_by_name['TransferResponse']
_PEGINADDRESSREQUEST = DESCRIPTOR.message_types_by_name['PegInAddressRequest']
_PEGINADDRESSRESPONSE = DESCRIPTOR.message_types_by_name['PegInAddressResponse']
_CLAIMPEGINREQUEST = DESCRIPTOR.message_types_by_name['ClaimPegInRequest']
_CLAIMPEGINRESPONSE = DESCRIPTOR.message_types_by_name['ClaimPegInResponse']
_SELECTUTXOSREQUEST_STRATEGY = _SELECTUTXOSREQUEST.enum_types_by_name['Strategy']
SelectUtxosRequest = _reflection.GeneratedProtocolMessageType('SelectUtxosRequest', (_message.Message,), {
  'DESCRIPTOR' : _SELECTUTXOSREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.SelectUtxosRequest)
  })
_sym_db.RegisterMessage(SelectUtxosRequest)

SelectUtxosResponse = _reflection.GeneratedProtocolMessageType('SelectUtxosResponse', (_message.Message,), {
  'DESCRIPTOR' : _SELECTUTXOSRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.SelectUtxosResponse)
  })
_sym_db.RegisterMessage(SelectUtxosResponse)

EstimateFeesRequest = _reflection.GeneratedProtocolMessageType('EstimateFeesRequest', (_message.Message,), {
  'DESCRIPTOR' : _ESTIMATEFEESREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.EstimateFeesRequest)
  })
_sym_db.RegisterMessage(EstimateFeesRequest)

EstimateFeesResponse = _reflection.GeneratedProtocolMessageType('EstimateFeesResponse', (_message.Message,), {
  'DESCRIPTOR' : _ESTIMATEFEESRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.EstimateFeesResponse)
  })
_sym_db.RegisterMessage(EstimateFeesResponse)

SignTransactionRequest = _reflection.GeneratedProtocolMessageType('SignTransactionRequest', (_message.Message,), {
  'DESCRIPTOR' : _SIGNTRANSACTIONREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.SignTransactionRequest)
  })
_sym_db.RegisterMessage(SignTransactionRequest)

SignTransactionResponse = _reflection.GeneratedProtocolMessageType('SignTransactionResponse', (_message.Message,), {
  'DESCRIPTOR' : _SIGNTRANSACTIONRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.SignTransactionResponse)
  })
_sym_db.RegisterMessage(SignTransactionResponse)

BroadcastTransactionRequest = _reflection.GeneratedProtocolMessageType('BroadcastTransactionRequest', (_message.Message,), {
  'DESCRIPTOR' : _BROADCASTTRANSACTIONREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BroadcastTransactionRequest)
  })
_sym_db.RegisterMessage(BroadcastTransactionRequest)

BroadcastTransactionResponse = _reflection.GeneratedProtocolMessageType('BroadcastTransactionResponse', (_message.Message,), {
  'DESCRIPTOR' : _BROADCASTTRANSACTIONRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BroadcastTransactionResponse)
  })
_sym_db.RegisterMessage(BroadcastTransactionResponse)

CreatePsetRequest = _reflection.GeneratedProtocolMessageType('CreatePsetRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPSETREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.CreatePsetRequest)
  })
_sym_db.RegisterMessage(CreatePsetRequest)

CreatePsetResponse = _reflection.GeneratedProtocolMessageType('CreatePsetResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPSETRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.CreatePsetResponse)
  })
_sym_db.RegisterMessage(CreatePsetResponse)

UpdatePsetRequest = _reflection.GeneratedProtocolMessageType('UpdatePsetRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPSETREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.UpdatePsetRequest)
  })
_sym_db.RegisterMessage(UpdatePsetRequest)

UpdatePsetResponse = _reflection.GeneratedProtocolMessageType('UpdatePsetResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPSETRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.UpdatePsetResponse)
  })
_sym_db.RegisterMessage(UpdatePsetResponse)

BlindPsetRequest = _reflection.GeneratedProtocolMessageType('BlindPsetRequest', (_message.Message,), {
  'DESCRIPTOR' : _BLINDPSETREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BlindPsetRequest)
  })
_sym_db.RegisterMessage(BlindPsetRequest)

BlindPsetResponse = _reflection.GeneratedProtocolMessageType('BlindPsetResponse', (_message.Message,), {
  'DESCRIPTOR' : _BLINDPSETRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BlindPsetResponse)
  })
_sym_db.RegisterMessage(BlindPsetResponse)

SignPsetRequest = _reflection.GeneratedProtocolMessageType('SignPsetRequest', (_message.Message,), {
  'DESCRIPTOR' : _SIGNPSETREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.SignPsetRequest)
  })
_sym_db.RegisterMessage(SignPsetRequest)

SignPsetResponse = _reflection.GeneratedProtocolMessageType('SignPsetResponse', (_message.Message,), {
  'DESCRIPTOR' : _SIGNPSETRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.SignPsetResponse)
  })
_sym_db.RegisterMessage(SignPsetResponse)

MintRequest = _reflection.GeneratedProtocolMessageType('MintRequest', (_message.Message,), {
  'DESCRIPTOR' : _MINTREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.MintRequest)
  })
_sym_db.RegisterMessage(MintRequest)

MintResponse = _reflection.GeneratedProtocolMessageType('MintResponse', (_message.Message,), {
  'DESCRIPTOR' : _MINTRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.MintResponse)
  })
_sym_db.RegisterMessage(MintResponse)

RemintRequest = _reflection.GeneratedProtocolMessageType('RemintRequest', (_message.Message,), {
  'DESCRIPTOR' : _REMINTREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.RemintRequest)
  })
_sym_db.RegisterMessage(RemintRequest)

RemintResponse = _reflection.GeneratedProtocolMessageType('RemintResponse', (_message.Message,), {
  'DESCRIPTOR' : _REMINTRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.RemintResponse)
  })
_sym_db.RegisterMessage(RemintResponse)

BurnRequest = _reflection.GeneratedProtocolMessageType('BurnRequest', (_message.Message,), {
  'DESCRIPTOR' : _BURNREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BurnRequest)
  })
_sym_db.RegisterMessage(BurnRequest)

BurnResponse = _reflection.GeneratedProtocolMessageType('BurnResponse', (_message.Message,), {
  'DESCRIPTOR' : _BURNRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BurnResponse)
  })
_sym_db.RegisterMessage(BurnResponse)

TransferRequest = _reflection.GeneratedProtocolMessageType('TransferRequest', (_message.Message,), {
  'DESCRIPTOR' : _TRANSFERREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.TransferRequest)
  })
_sym_db.RegisterMessage(TransferRequest)

TransferResponse = _reflection.GeneratedProtocolMessageType('TransferResponse', (_message.Message,), {
  'DESCRIPTOR' : _TRANSFERRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.TransferResponse)
  })
_sym_db.RegisterMessage(TransferResponse)

PegInAddressRequest = _reflection.GeneratedProtocolMessageType('PegInAddressRequest', (_message.Message,), {
  'DESCRIPTOR' : _PEGINADDRESSREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.PegInAddressRequest)
  })
_sym_db.RegisterMessage(PegInAddressRequest)

PegInAddressResponse = _reflection.GeneratedProtocolMessageType('PegInAddressResponse', (_message.Message,), {
  'DESCRIPTOR' : _PEGINADDRESSRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.PegInAddressResponse)
  })
_sym_db.RegisterMessage(PegInAddressResponse)

ClaimPegInRequest = _reflection.GeneratedProtocolMessageType('ClaimPegInRequest', (_message.Message,), {
  'DESCRIPTOR' : _CLAIMPEGINREQUEST,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.ClaimPegInRequest)
  })
_sym_db.RegisterMessage(ClaimPegInRequest)

ClaimPegInResponse = _reflection.GeneratedProtocolMessageType('ClaimPegInResponse', (_message.Message,), {
  'DESCRIPTOR' : _CLAIMPEGINRESPONSE,
  '__module__' : 'ocean.v1alpha.transaction_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.ClaimPegInResponse)
  })
_sym_db.RegisterMessage(ClaimPegInResponse)

_TRANSACTIONSERVICE = DESCRIPTOR.services_by_name['TransactionService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.ocean.v1alphaB\020TransactionProtoP\001Zagithub.com/vulpemventures/ocean/api-spec/protobuf/ocean/v1alpha/gen/go/ocean/v1alpha;oceanv1alpha\242\002\003OXX\252\002\rOcean.V1alpha\312\002\rOcean\\V1alpha\342\002\031Ocean\\V1alpha\\GPBMetadata\352\002\016Ocean::V1alpha'
  _SELECTUTXOSREQUEST._serialized_start=78
  _SELECTUTXOSREQUEST._serialized_end=390
  _SELECTUTXOSREQUEST_STRATEGY._serialized_start=304
  _SELECTUTXOSREQUEST_STRATEGY._serialized_end=390
  _SELECTUTXOSRESPONSE._serialized_start=392
  _SELECTUTXOSRESPONSE._serialized_end=480
  _ESTIMATEFEESREQUEST._serialized_start=482
  _ESTIMATEFEESREQUEST._serialized_end=598
  _ESTIMATEFEESRESPONSE._serialized_start=600
  _ESTIMATEFEESRESPONSE._serialized_end=653
  _SIGNTRANSACTIONREQUEST._serialized_start=655
  _SIGNTRANSACTIONREQUEST._serialized_end=702
  _SIGNTRANSACTIONRESPONSE._serialized_start=704
  _SIGNTRANSACTIONRESPONSE._serialized_end=752
  _BROADCASTTRANSACTIONREQUEST._serialized_start=754
  _BROADCASTTRANSACTIONREQUEST._serialized_end=806
  _BROADCASTTRANSACTIONRESPONSE._serialized_start=808
  _BROADCASTTRANSACTIONRESPONSE._serialized_end=858
  _CREATEPSETREQUEST._serialized_start=860
  _CREATEPSETREQUEST._serialized_end=974
  _CREATEPSETRESPONSE._serialized_start=976
  _CREATEPSETRESPONSE._serialized_end=1016
  _UPDATEPSETREQUEST._serialized_start=1019
  _UPDATEPSETREQUEST._serialized_end=1153
  _UPDATEPSETRESPONSE._serialized_start=1155
  _UPDATEPSETRESPONSE._serialized_end=1195
  _BLINDPSETREQUEST._serialized_start=1197
  _BLINDPSETREQUEST._serialized_end=1235
  _BLINDPSETRESPONSE._serialized_start=1237
  _BLINDPSETRESPONSE._serialized_end=1276
  _SIGNPSETREQUEST._serialized_start=1278
  _SIGNPSETREQUEST._serialized_end=1315
  _SIGNPSETRESPONSE._serialized_start=1317
  _SIGNPSETRESPONSE._serialized_end=1355
  _MINTREQUEST._serialized_start=1358
  _MINTREQUEST._serialized_end=1602
  _MINTRESPONSE._serialized_start=1604
  _MINTRESPONSE._serialized_end=1641
  _REMINTREQUEST._serialized_start=1643
  _REMINTREQUEST._serialized_end=1764
  _REMINTRESPONSE._serialized_start=1766
  _REMINTRESPONSE._serialized_end=1805
  _BURNREQUEST._serialized_start=1807
  _BURNREQUEST._serialized_end=1933
  _BURNRESPONSE._serialized_start=1935
  _BURNRESPONSE._serialized_end=1972
  _TRANSFERREQUEST._serialized_start=1975
  _TRANSFERREQUEST._serialized_end=2105
  _TRANSFERRESPONSE._serialized_start=2107
  _TRANSFERRESPONSE._serialized_end=2148
  _PEGINADDRESSREQUEST._serialized_start=2150
  _PEGINADDRESSREQUEST._serialized_end=2171
  _PEGINADDRESSRESPONSE._serialized_start=2174
  _PEGINADDRESSRESPONSE._serialized_end=2337
  _CLAIMPEGINREQUEST._serialized_start=2339
  _CLAIMPEGINREQUEST._serialized_end=2458
  _CLAIMPEGINRESPONSE._serialized_start=2460
  _CLAIMPEGINRESPONSE._serialized_end=2503
  _TRANSACTIONSERVICE._serialized_start=2506
  _TRANSACTIONSERVICE._serialized_end=3685
# @@protoc_insertion_point(module_scope)
