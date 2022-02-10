# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ocean/v1alpha/notification.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ocean.v1alpha import types_pb2 as ocean_dot_v1alpha_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n ocean/v1alpha/notification.proto\x12\rocean.v1alpha\x1a\x19ocean/v1alpha/types.proto\"]\n\x1fTransactionNotificationsRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\"\xef\x01\n TransactionNotificationsResponse\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12\x12\n\x04txid\x18\x02 \x01(\tR\x04txid\x12\x39\n\nevent_type\x18\x03 \x01(\x0e\x32\x1a.ocean.v1alpha.TxEventTypeR\teventType\x12@\n\rblock_details\x18\x04 \x01(\x0b\x32\x1b.ocean.v1alpha.BlockDetailsR\x0c\x62lockDetails\"X\n\x0c\x42lockDetails\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\x12\x16\n\x06height\x18\x02 \x01(\x05R\x06height\x12\x1c\n\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\"W\n\x19UtxosNotificationsRequest\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\"\xbe\x01\n\x1aUtxosNotificationsResponse\x12:\n\x0b\x61\x63\x63ount_key\x18\x01 \x01(\x0b\x32\x19.ocean.v1alpha.AccountKeyR\naccountKey\x12\'\n\x04utxo\x18\x02 \x01(\x0b\x32\x13.ocean.v1alpha.UtxoR\x04utxo\x12;\n\nevent_type\x18\x03 \x01(\x0e\x32\x1c.ocean.v1alpha.UtxoEventTypeR\teventType\"\x87\x01\n\x11\x41\x64\x64WebhookRequest\x12\x1a\n\x08\x65ndpoint\x18\x01 \x01(\tR\x08\x65ndpoint\x12>\n\nevent_type\x18\x02 \x01(\x0e\x32\x1f.ocean.v1alpha.WebhookEventTypeR\teventType\x12\x16\n\x06secret\x18\x03 \x01(\tR\x06secret\"$\n\x12\x41\x64\x64WebhookResponse\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"&\n\x14RemoveWebhookRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\x17\n\x15RemoveWebhookResponse\"U\n\x13ListWebhooksRequest\x12>\n\nevent_type\x18\x01 \x01(\x0e\x32\x1f.ocean.v1alpha.WebhookEventTypeR\teventType\"U\n\x14ListWebhooksResponse\x12=\n\x0cwebhook_info\x18\x01 \x03(\x0b\x32\x1a.ocean.v1alpha.WebhookInfoR\x0bwebhookInfo\"X\n\x0bWebhookInfo\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n\x08\x65ndpoint\x18\x02 \x01(\tR\x08\x65ndpoint\x12\x1d\n\nis_secured\x18\x03 \x01(\x08R\tisSecured2\x8f\x04\n\x13NotificationService\x12}\n\x18TransactionNotifications\x12..ocean.v1alpha.TransactionNotificationsRequest\x1a/.ocean.v1alpha.TransactionNotificationsResponse0\x01\x12k\n\x12UtxosNotifications\x12(.ocean.v1alpha.UtxosNotificationsRequest\x1a).ocean.v1alpha.UtxosNotificationsResponse0\x01\x12S\n\nAddWebhook\x12 .ocean.v1alpha.AddWebhookRequest\x1a!.ocean.v1alpha.AddWebhookResponse\"\x00\x12\\\n\rRemoveWebhook\x12#.ocean.v1alpha.RemoveWebhookRequest\x1a$.ocean.v1alpha.RemoveWebhookResponse\"\x00\x12Y\n\x0cListWebhooks\x12\".ocean.v1alpha.ListWebhooksRequest\x1a#.ocean.v1alpha.ListWebhooksResponse\"\x00\x42\xde\x01\n\x11\x63om.ocean.v1alphaB\x11NotificationProtoP\x01Zagithub.com/vulpemventures/ocean/api-spec/protobuf/ocean/v1alpha/gen/go/ocean/v1alpha;oceanv1alpha\xa2\x02\x03OXX\xaa\x02\rOcean.V1alpha\xca\x02\rOcean\\V1alpha\xe2\x02\x19Ocean\\V1alpha\\GPBMetadata\xea\x02\x0eOcean::V1alphab\x06proto3')



_TRANSACTIONNOTIFICATIONSREQUEST = DESCRIPTOR.message_types_by_name['TransactionNotificationsRequest']
_TRANSACTIONNOTIFICATIONSRESPONSE = DESCRIPTOR.message_types_by_name['TransactionNotificationsResponse']
_BLOCKDETAILS = DESCRIPTOR.message_types_by_name['BlockDetails']
_UTXOSNOTIFICATIONSREQUEST = DESCRIPTOR.message_types_by_name['UtxosNotificationsRequest']
_UTXOSNOTIFICATIONSRESPONSE = DESCRIPTOR.message_types_by_name['UtxosNotificationsResponse']
_ADDWEBHOOKREQUEST = DESCRIPTOR.message_types_by_name['AddWebhookRequest']
_ADDWEBHOOKRESPONSE = DESCRIPTOR.message_types_by_name['AddWebhookResponse']
_REMOVEWEBHOOKREQUEST = DESCRIPTOR.message_types_by_name['RemoveWebhookRequest']
_REMOVEWEBHOOKRESPONSE = DESCRIPTOR.message_types_by_name['RemoveWebhookResponse']
_LISTWEBHOOKSREQUEST = DESCRIPTOR.message_types_by_name['ListWebhooksRequest']
_LISTWEBHOOKSRESPONSE = DESCRIPTOR.message_types_by_name['ListWebhooksResponse']
_WEBHOOKINFO = DESCRIPTOR.message_types_by_name['WebhookInfo']
TransactionNotificationsRequest = _reflection.GeneratedProtocolMessageType('TransactionNotificationsRequest', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTIONNOTIFICATIONSREQUEST,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.TransactionNotificationsRequest)
  })
_sym_db.RegisterMessage(TransactionNotificationsRequest)

TransactionNotificationsResponse = _reflection.GeneratedProtocolMessageType('TransactionNotificationsResponse', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTIONNOTIFICATIONSRESPONSE,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.TransactionNotificationsResponse)
  })
_sym_db.RegisterMessage(TransactionNotificationsResponse)

BlockDetails = _reflection.GeneratedProtocolMessageType('BlockDetails', (_message.Message,), {
  'DESCRIPTOR' : _BLOCKDETAILS,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.BlockDetails)
  })
_sym_db.RegisterMessage(BlockDetails)

UtxosNotificationsRequest = _reflection.GeneratedProtocolMessageType('UtxosNotificationsRequest', (_message.Message,), {
  'DESCRIPTOR' : _UTXOSNOTIFICATIONSREQUEST,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.UtxosNotificationsRequest)
  })
_sym_db.RegisterMessage(UtxosNotificationsRequest)

UtxosNotificationsResponse = _reflection.GeneratedProtocolMessageType('UtxosNotificationsResponse', (_message.Message,), {
  'DESCRIPTOR' : _UTXOSNOTIFICATIONSRESPONSE,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.UtxosNotificationsResponse)
  })
_sym_db.RegisterMessage(UtxosNotificationsResponse)

AddWebhookRequest = _reflection.GeneratedProtocolMessageType('AddWebhookRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDWEBHOOKREQUEST,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.AddWebhookRequest)
  })
_sym_db.RegisterMessage(AddWebhookRequest)

AddWebhookResponse = _reflection.GeneratedProtocolMessageType('AddWebhookResponse', (_message.Message,), {
  'DESCRIPTOR' : _ADDWEBHOOKRESPONSE,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.AddWebhookResponse)
  })
_sym_db.RegisterMessage(AddWebhookResponse)

RemoveWebhookRequest = _reflection.GeneratedProtocolMessageType('RemoveWebhookRequest', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEWEBHOOKREQUEST,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.RemoveWebhookRequest)
  })
_sym_db.RegisterMessage(RemoveWebhookRequest)

RemoveWebhookResponse = _reflection.GeneratedProtocolMessageType('RemoveWebhookResponse', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEWEBHOOKRESPONSE,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.RemoveWebhookResponse)
  })
_sym_db.RegisterMessage(RemoveWebhookResponse)

ListWebhooksRequest = _reflection.GeneratedProtocolMessageType('ListWebhooksRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTWEBHOOKSREQUEST,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.ListWebhooksRequest)
  })
_sym_db.RegisterMessage(ListWebhooksRequest)

ListWebhooksResponse = _reflection.GeneratedProtocolMessageType('ListWebhooksResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTWEBHOOKSRESPONSE,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.ListWebhooksResponse)
  })
_sym_db.RegisterMessage(ListWebhooksResponse)

WebhookInfo = _reflection.GeneratedProtocolMessageType('WebhookInfo', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKINFO,
  '__module__' : 'ocean.v1alpha.notification_pb2'
  # @@protoc_insertion_point(class_scope:ocean.v1alpha.WebhookInfo)
  })
_sym_db.RegisterMessage(WebhookInfo)

_NOTIFICATIONSERVICE = DESCRIPTOR.services_by_name['NotificationService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.ocean.v1alphaB\021NotificationProtoP\001Zagithub.com/vulpemventures/ocean/api-spec/protobuf/ocean/v1alpha/gen/go/ocean/v1alpha;oceanv1alpha\242\002\003OXX\252\002\rOcean.V1alpha\312\002\rOcean\\V1alpha\342\002\031Ocean\\V1alpha\\GPBMetadata\352\002\016Ocean::V1alpha'
  _TRANSACTIONNOTIFICATIONSREQUEST._serialized_start=78
  _TRANSACTIONNOTIFICATIONSREQUEST._serialized_end=171
  _TRANSACTIONNOTIFICATIONSRESPONSE._serialized_start=174
  _TRANSACTIONNOTIFICATIONSRESPONSE._serialized_end=413
  _BLOCKDETAILS._serialized_start=415
  _BLOCKDETAILS._serialized_end=503
  _UTXOSNOTIFICATIONSREQUEST._serialized_start=505
  _UTXOSNOTIFICATIONSREQUEST._serialized_end=592
  _UTXOSNOTIFICATIONSRESPONSE._serialized_start=595
  _UTXOSNOTIFICATIONSRESPONSE._serialized_end=785
  _ADDWEBHOOKREQUEST._serialized_start=788
  _ADDWEBHOOKREQUEST._serialized_end=923
  _ADDWEBHOOKRESPONSE._serialized_start=925
  _ADDWEBHOOKRESPONSE._serialized_end=961
  _REMOVEWEBHOOKREQUEST._serialized_start=963
  _REMOVEWEBHOOKREQUEST._serialized_end=1001
  _REMOVEWEBHOOKRESPONSE._serialized_start=1003
  _REMOVEWEBHOOKRESPONSE._serialized_end=1026
  _LISTWEBHOOKSREQUEST._serialized_start=1028
  _LISTWEBHOOKSREQUEST._serialized_end=1113
  _LISTWEBHOOKSRESPONSE._serialized_start=1115
  _LISTWEBHOOKSRESPONSE._serialized_end=1200
  _WEBHOOKINFO._serialized_start=1202
  _WEBHOOKINFO._serialized_end=1290
  _NOTIFICATIONSERVICE._serialized_start=1293
  _NOTIFICATIONSERVICE._serialized_end=1820
# @@protoc_insertion_point(module_scope)
