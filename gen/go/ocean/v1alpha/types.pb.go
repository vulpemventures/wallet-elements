// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: ocean/v1alpha/types.proto

package oceanv1alpha

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TxEventType int32

const (
	TxEventType_TX_EVENT_TYPE_UNSPECIFIED TxEventType = 0
	// Tx broadcasted.
	TxEventType_TX_EVENT_TYPE_BROADCASTED TxEventType = 1
	// Tx unconfirmed.
	TxEventType_TX_EVENT_TYPE_UNCONFIRMED TxEventType = 2
	// Tx confirmed.
	TxEventType_TX_EVENT_TYPE_CONFIRMED TxEventType = 3
)

// Enum value maps for TxEventType.
var (
	TxEventType_name = map[int32]string{
		0: "TX_EVENT_TYPE_UNSPECIFIED",
		1: "TX_EVENT_TYPE_BROADCASTED",
		2: "TX_EVENT_TYPE_UNCONFIRMED",
		3: "TX_EVENT_TYPE_CONFIRMED",
	}
	TxEventType_value = map[string]int32{
		"TX_EVENT_TYPE_UNSPECIFIED": 0,
		"TX_EVENT_TYPE_BROADCASTED": 1,
		"TX_EVENT_TYPE_UNCONFIRMED": 2,
		"TX_EVENT_TYPE_CONFIRMED":   3,
	}
)

func (x TxEventType) Enum() *TxEventType {
	p := new(TxEventType)
	*p = x
	return p
}

func (x TxEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_ocean_v1alpha_types_proto_enumTypes[0].Descriptor()
}

func (TxEventType) Type() protoreflect.EnumType {
	return &file_ocean_v1alpha_types_proto_enumTypes[0]
}

func (x TxEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxEventType.Descriptor instead.
func (TxEventType) EnumDescriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{0}
}

type UtxoEventType int32

const (
	UtxoEventType_UTXO_EVENT_TYPE_UNSPECIFIED UtxoEventType = 0
	UtxoEventType_UTXO_EVENT_TYPE_LOCKED      UtxoEventType = 1
	UtxoEventType_UTXO_EVENT_TYPE_UNLOCKED    UtxoEventType = 2
	UtxoEventType_UTXO_EVENT_TYPE_SPENT       UtxoEventType = 3
	UtxoEventType_UTXO_EVENT_TYPE_NEW         UtxoEventType = 4
)

// Enum value maps for UtxoEventType.
var (
	UtxoEventType_name = map[int32]string{
		0: "UTXO_EVENT_TYPE_UNSPECIFIED",
		1: "UTXO_EVENT_TYPE_LOCKED",
		2: "UTXO_EVENT_TYPE_UNLOCKED",
		3: "UTXO_EVENT_TYPE_SPENT",
		4: "UTXO_EVENT_TYPE_NEW",
	}
	UtxoEventType_value = map[string]int32{
		"UTXO_EVENT_TYPE_UNSPECIFIED": 0,
		"UTXO_EVENT_TYPE_LOCKED":      1,
		"UTXO_EVENT_TYPE_UNLOCKED":    2,
		"UTXO_EVENT_TYPE_SPENT":       3,
		"UTXO_EVENT_TYPE_NEW":         4,
	}
)

func (x UtxoEventType) Enum() *UtxoEventType {
	p := new(UtxoEventType)
	*p = x
	return p
}

func (x UtxoEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UtxoEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_ocean_v1alpha_types_proto_enumTypes[1].Descriptor()
}

func (UtxoEventType) Type() protoreflect.EnumType {
	return &file_ocean_v1alpha_types_proto_enumTypes[1]
}

func (x UtxoEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UtxoEventType.Descriptor instead.
func (UtxoEventType) EnumDescriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{1}
}

type WebhookEventType int32

const (
	WebhookEventType_WEBHOOK_EVENT_TYPE_UNSPECIFIED WebhookEventType = 0
	// Receive notification about transactions.
	WebhookEventType_WEBHOOK_EVENT_TYPE_TRANSACTION WebhookEventType = 1
	// Receive notifications about utxos.
	WebhookEventType_WEBHOOK_EVENT_TYPE_UNSPENTS WebhookEventType = 2
)

// Enum value maps for WebhookEventType.
var (
	WebhookEventType_name = map[int32]string{
		0: "WEBHOOK_EVENT_TYPE_UNSPECIFIED",
		1: "WEBHOOK_EVENT_TYPE_TRANSACTION",
		2: "WEBHOOK_EVENT_TYPE_UNSPENTS",
	}
	WebhookEventType_value = map[string]int32{
		"WEBHOOK_EVENT_TYPE_UNSPECIFIED": 0,
		"WEBHOOK_EVENT_TYPE_TRANSACTION": 1,
		"WEBHOOK_EVENT_TYPE_UNSPENTS":    2,
	}
)

func (x WebhookEventType) Enum() *WebhookEventType {
	p := new(WebhookEventType)
	*p = x
	return p
}

func (x WebhookEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebhookEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_ocean_v1alpha_types_proto_enumTypes[2].Descriptor()
}

func (WebhookEventType) Type() protoreflect.EnumType {
	return &file_ocean_v1alpha_types_proto_enumTypes[2]
}

func (x WebhookEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WebhookEventType.Descriptor instead.
func (WebhookEventType) EnumDescriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{2}
}

type Template_Format int32

const (
	Template_FORMAT_UNSPECIFIED Template_Format = 0
	Template_FORMAT_DESCRIPTOR  Template_Format = 1
	Template_FORMAT_MINISCRIPT  Template_Format = 2
	Template_FORMAT_IONIO       Template_Format = 3
	Template_FORMAT_RAW         Template_Format = 4
)

// Enum value maps for Template_Format.
var (
	Template_Format_name = map[int32]string{
		0: "FORMAT_UNSPECIFIED",
		1: "FORMAT_DESCRIPTOR",
		2: "FORMAT_MINISCRIPT",
		3: "FORMAT_IONIO",
		4: "FORMAT_RAW",
	}
	Template_Format_value = map[string]int32{
		"FORMAT_UNSPECIFIED": 0,
		"FORMAT_DESCRIPTOR":  1,
		"FORMAT_MINISCRIPT":  2,
		"FORMAT_IONIO":       3,
		"FORMAT_RAW":         4,
	}
)

func (x Template_Format) Enum() *Template_Format {
	p := new(Template_Format)
	*p = x
	return p
}

func (x Template_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Template_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_ocean_v1alpha_types_proto_enumTypes[3].Descriptor()
}

func (Template_Format) Type() protoreflect.EnumType {
	return &file_ocean_v1alpha_types_proto_enumTypes[3]
}

func (x Template_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Template_Format.Descriptor instead.
func (Template_Format) EnumDescriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{7, 0}
}

type AccountKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the account.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the account to be updated.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountKey) Reset() {
	*x = AccountKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountKey) ProtoMessage() {}

func (x *AccountKey) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountKey.ProtoReflect.Descriptor instead.
func (*AccountKey) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{0}
}

func (x *AccountKey) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account key.
	AccountKey *AccountKey `protobuf:"bytes,1,opt,name=account_key,json=accountKey,proto3" json:"account_key,omitempty"`
	// Derivation path.
	DerivationPath string `protobuf:"bytes,2,opt,name=derivation_path,json=derivationPath,proto3" json:"derivation_path,omitempty"`
	// XPub.
	Xpub string `protobuf:"bytes,3,opt,name=xpub,proto3" json:"xpub,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{1}
}

func (x *AccountInfo) GetAccountKey() *AccountKey {
	if x != nil {
		return x.AccountKey
	}
	return nil
}

func (x *AccountInfo) GetDerivationPath() string {
	if x != nil {
		return x.DerivationPath
	}
	return ""
}

func (x *AccountInfo) GetXpub() string {
	if x != nil {
		return x.Xpub
	}
	return ""
}

type BalanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The balance of the wallet (excludes locked utxos)
	TotalBalance uint64 `protobuf:"varint,1,opt,name=total_balance,json=totalBalance,proto3" json:"total_balance,omitempty"`
	// The confirmed balance of a wallet (includes utxos with 1+ confirmations)
	ConfirmedBalance uint64 `protobuf:"varint,2,opt,name=confirmed_balance,json=confirmedBalance,proto3" json:"confirmed_balance,omitempty"`
	// The unconfirmed balance of a wallet (includes utxos with no confirmations)
	UnconfirmedBalance uint64 `protobuf:"varint,3,opt,name=unconfirmed_balance,json=unconfirmedBalance,proto3" json:"unconfirmed_balance,omitempty"`
}

func (x *BalanceInfo) Reset() {
	*x = BalanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceInfo) ProtoMessage() {}

func (x *BalanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceInfo.ProtoReflect.Descriptor instead.
func (*BalanceInfo) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceInfo) GetTotalBalance() uint64 {
	if x != nil {
		return x.TotalBalance
	}
	return 0
}

func (x *BalanceInfo) GetConfirmedBalance() uint64 {
	if x != nil {
		return x.ConfirmedBalance
	}
	return 0
}

func (x *BalanceInfo) GetUnconfirmedBalance() uint64 {
	if x != nil {
		return x.UnconfirmedBalance
	}
	return 0
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Previous output txid.
	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	// Previous tx output index.
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{3}
}

func (x *Input) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Input) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asset hash.
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// Sent amount.
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// Address to send funds to.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{4}
}

func (x *Output) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *Output) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Output) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Utxos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// account key
	AccountKey *AccountKey `protobuf:"bytes,1,opt,name=account_key,json=accountKey,proto3" json:"account_key,omitempty"`
	// list of utxos
	Utxos []*Utxo `protobuf:"bytes,2,rep,name=utxos,proto3" json:"utxos,omitempty"`
}

func (x *Utxos) Reset() {
	*x = Utxos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utxos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utxos) ProtoMessage() {}

func (x *Utxos) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utxos.ProtoReflect.Descriptor instead.
func (*Utxos) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{5}
}

func (x *Utxos) GetAccountKey() *AccountKey {
	if x != nil {
		return x.AccountKey
	}
	return nil
}

func (x *Utxos) GetUtxos() []*Utxo {
	if x != nil {
		return x.Utxos
	}
	return nil
}

type Utxo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Txid of the Uuxo
	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	// Output index
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// Asset
	Asset string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	// Value
	Value uint64 `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	// Script
	Script []byte `protobuf:"bytes,5,opt,name=script,proto3" json:"script,omitempty"`
	// Confirmation status
	IsConfirmed bool `protobuf:"varint,6,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	// Whether locked by a coin selection
	IsLocked bool `protobuf:"varint,7,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty"`
}

func (x *Utxo) Reset() {
	*x = Utxo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utxo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utxo) ProtoMessage() {}

func (x *Utxo) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utxo.ProtoReflect.Descriptor instead.
func (*Utxo) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{6}
}

func (x *Utxo) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Utxo) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Utxo) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *Utxo) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Utxo) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *Utxo) GetIsConfirmed() bool {
	if x != nil {
		return x.IsConfirmed
	}
	return false
}

func (x *Utxo) GetIsLocked() bool {
	if x != nil {
		return x.IsLocked
	}
	return false
}

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format Template_Format `protobuf:"varint,1,opt,name=format,proto3,enum=ocean.v1alpha.Template_Format" json:"format,omitempty"`
	Value  string          `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_v1alpha_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_v1alpha_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_ocean_v1alpha_types_proto_rawDescGZIP(), []int{7}
}

func (x *Template) GetFormat() Template_Format {
	if x != nil {
		return x.Format
	}
	return Template_FORMAT_UNSPECIFIED
}

func (x *Template) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_ocean_v1alpha_types_proto protoreflect.FileDescriptor

var file_ocean_v1alpha_types_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x63, 0x65,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x70, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x78, 0x70, 0x75, 0x62, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x75, 0x6e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x50, 0x0a, 0x06, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a,
	0x05, 0x55, 0x74, 0x78, 0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x75, 0x74, 0x78, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x55, 0x74, 0x78, 0x6f, 0x52, 0x05, 0x75, 0x74, 0x78, 0x6f, 0x73, 0x22, 0xb4, 0x01,
	0x0a, 0x04, 0x55, 0x74, 0x78, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x70, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x4f, 0x52,
	0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43,
	0x52, 0x49, 0x50, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x49, 0x4f, 0x4e, 0x49, 0x4f, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x52, 0x41, 0x57, 0x10,
	0x04, 0x2a, 0x87, 0x01, 0x0a, 0x0b, 0x54, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x58, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1d, 0x0a, 0x19, 0x54, 0x58, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x54, 0x58, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1b,
	0x0a, 0x17, 0x54, 0x58, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x9e, 0x01, 0x0a, 0x0d,
	0x55, 0x74, 0x78, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x1b, 0x55, 0x54, 0x58, 0x4f, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a,
	0x0a, 0x16, 0x55, 0x54, 0x58, 0x4f, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x54,
	0x58, 0x4f, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x54, 0x58, 0x4f,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x58, 0x4f, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x04, 0x2a, 0x7b, 0x0a, 0x10,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x1e, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x57, 0x45, 0x42, 0x48,
	0x4f, 0x4f, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x02, 0x42, 0xd7, 0x01, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x75, 0x6c, 0x70, 0x65, 0x6d,
	0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x3b, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x0d, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x19, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0e, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocean_v1alpha_types_proto_rawDescOnce sync.Once
	file_ocean_v1alpha_types_proto_rawDescData = file_ocean_v1alpha_types_proto_rawDesc
)

func file_ocean_v1alpha_types_proto_rawDescGZIP() []byte {
	file_ocean_v1alpha_types_proto_rawDescOnce.Do(func() {
		file_ocean_v1alpha_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocean_v1alpha_types_proto_rawDescData)
	})
	return file_ocean_v1alpha_types_proto_rawDescData
}

var file_ocean_v1alpha_types_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_ocean_v1alpha_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ocean_v1alpha_types_proto_goTypes = []interface{}{
	(TxEventType)(0),      // 0: ocean.v1alpha.TxEventType
	(UtxoEventType)(0),    // 1: ocean.v1alpha.UtxoEventType
	(WebhookEventType)(0), // 2: ocean.v1alpha.WebhookEventType
	(Template_Format)(0),  // 3: ocean.v1alpha.Template.Format
	(*AccountKey)(nil),    // 4: ocean.v1alpha.AccountKey
	(*AccountInfo)(nil),   // 5: ocean.v1alpha.AccountInfo
	(*BalanceInfo)(nil),   // 6: ocean.v1alpha.BalanceInfo
	(*Input)(nil),         // 7: ocean.v1alpha.Input
	(*Output)(nil),        // 8: ocean.v1alpha.Output
	(*Utxos)(nil),         // 9: ocean.v1alpha.Utxos
	(*Utxo)(nil),          // 10: ocean.v1alpha.Utxo
	(*Template)(nil),      // 11: ocean.v1alpha.Template
}
var file_ocean_v1alpha_types_proto_depIdxs = []int32{
	4,  // 0: ocean.v1alpha.AccountInfo.account_key:type_name -> ocean.v1alpha.AccountKey
	4,  // 1: ocean.v1alpha.Utxos.account_key:type_name -> ocean.v1alpha.AccountKey
	10, // 2: ocean.v1alpha.Utxos.utxos:type_name -> ocean.v1alpha.Utxo
	3,  // 3: ocean.v1alpha.Template.format:type_name -> ocean.v1alpha.Template.Format
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_ocean_v1alpha_types_proto_init() }
func file_ocean_v1alpha_types_proto_init() {
	if File_ocean_v1alpha_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocean_v1alpha_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utxos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utxo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ocean_v1alpha_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocean_v1alpha_types_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocean_v1alpha_types_proto_goTypes,
		DependencyIndexes: file_ocean_v1alpha_types_proto_depIdxs,
		EnumInfos:         file_ocean_v1alpha_types_proto_enumTypes,
		MessageInfos:      file_ocean_v1alpha_types_proto_msgTypes,
	}.Build()
	File_ocean_v1alpha_types_proto = out.File
	file_ocean_v1alpha_types_proto_rawDesc = nil
	file_ocean_v1alpha_types_proto_goTypes = nil
	file_ocean_v1alpha_types_proto_depIdxs = nil
}
