// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: types.proto

package gen

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

type AccountKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the account.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the account to be updated.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountKey) Reset() {
	*x = AccountKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountKey) ProtoMessage() {}

func (x *AccountKey) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
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
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *AccountKey) GetId() int64 {
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

type Unspents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// account key
	AccountKey *AccountKey `protobuf:"bytes,1,opt,name=account_key,json=accountKey,proto3" json:"account_key,omitempty"`
	// list of unspents
	Unspents []*Unspent `protobuf:"bytes,2,rep,name=unspents,proto3" json:"unspents,omitempty"`
}

func (x *Unspents) Reset() {
	*x = Unspents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unspents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unspents) ProtoMessage() {}

func (x *Unspents) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unspents.ProtoReflect.Descriptor instead.
func (*Unspents) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Unspents) GetAccountKey() *AccountKey {
	if x != nil {
		return x.AccountKey
	}
	return nil
}

func (x *Unspents) GetUnspents() []*Unspent {
	if x != nil {
		return x.Unspents
	}
	return nil
}

type Unspent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tx id of the unspent
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// Output index
	Index int64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// Asset
	Asset []byte `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	// Asset commitment, empty if asset is not confidential
	AssetCommitment []byte `protobuf:"bytes,4,opt,name=asset_commitment,json=assetCommitment,proto3" json:"asset_commitment,omitempty"`
	// Value
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	// Value commitment, empty if value is not confidential
	ValueCommitment []byte `protobuf:"bytes,6,opt,name=value_commitment,json=valueCommitment,proto3" json:"value_commitment,omitempty"`
	// Script
	Script []byte `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	// Nonce
	Nonce []byte `protobuf:"bytes,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Range proof
	RangeProof []byte `protobuf:"bytes,9,opt,name=rangeProof,proto3" json:"rangeProof,omitempty"`
	// Surjection proof
	SurjectionProof []byte `protobuf:"bytes,10,opt,name=surjectionProof,proto3" json:"surjectionProof,omitempty"`
}

func (x *Unspent) Reset() {
	*x = Unspent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unspent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unspent) ProtoMessage() {}

func (x *Unspent) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unspent.ProtoReflect.Descriptor instead.
func (*Unspent) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *Unspent) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *Unspent) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Unspent) GetAsset() []byte {
	if x != nil {
		return x.Asset
	}
	return nil
}

func (x *Unspent) GetAssetCommitment() []byte {
	if x != nil {
		return x.AssetCommitment
	}
	return nil
}

func (x *Unspent) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Unspent) GetValueCommitment() []byte {
	if x != nil {
		return x.ValueCommitment
	}
	return nil
}

func (x *Unspent) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *Unspent) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Unspent) GetRangeProof() []byte {
	if x != nil {
		return x.RangeProof
	}
	return nil
}

func (x *Unspent) GetSurjectionProof() []byte {
	if x != nil {
		return x.SurjectionProof
	}
	return nil
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a,
	0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x5e, 0x0a, 0x08, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x08, 0x75, 0x6e, 0x73,
	0x70, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x55, 0x6e,
	0x73, 0x70, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xae, 0x02, 0x0a, 0x07, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x74,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x75, 0x72, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0f, 0x73, 0x75, 0x72, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x75, 0x6c, 0x70, 0x65, 0x6d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_types_proto_goTypes = []interface{}{
	(*AccountKey)(nil), // 0: AccountKey
	(*Unspents)(nil),   // 1: Unspents
	(*Unspent)(nil),    // 2: Unspent
}
var file_types_proto_depIdxs = []int32{
	0, // 0: Unspents.account_key:type_name -> AccountKey
	2, // 1: Unspents.unspents:type_name -> Unspent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unspents); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unspent); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
