// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package _go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// Returns fee cost of one byte in satoshi's.
	Fee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeReply, error)
	// Updates fee cost of one byte.
	UpdateFee(ctx context.Context, in *UpdateFeeRequest, opts ...grpc.CallOption) (*UpdateFeeReply, error)
	// Signs transaction hex.
	SignTransaction(ctx context.Context, in *SignTransactionRequest, opts ...grpc.CallOption) (*SignTransactionReply, error)
	// Broadcast signed transaction.
	BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*BroadcastTransactionReply, error)
	// Creates unsigned, non-blinded psbt based on provided inputs and outputs.
	CreatePsbt(ctx context.Context, in *CreatePsbtRequest, opts ...grpc.CallOption) (*CreatePsbtReply, error)
	// Blind unblinded psbt.
	BlindPsbt(ctx context.Context, in *BlindPsbtRequest, opts ...grpc.CallOption) (*BlindPsbtReply, error)
	// Sign unsigned psbt.
	SignPsbt(ctx context.Context, in *SignPsbtRequest, opts ...grpc.CallOption) (*SignPsbtReply, error)
	// Mint asset.
	Mint(ctx context.Context, in *MintRequest, opts ...grpc.CallOption) (*MintReply, error)
	// ReMint asset.
	ReMint(ctx context.Context, in *ReMintRequest, opts ...grpc.CallOption) (*ReMintReply, error)
	// Burn minted assets belongs to list of addresses.
	Burn(ctx context.Context, in *BurnRequest, opts ...grpc.CallOption) (*BurnReply, error)
	// Transfer asset to one or more receivers.
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error)
	//
	// Returns information needed for claim pegin to move coins to the side-chain.
	// The user should send coins from their Bitcoin wallet to the main-chain address returned.
	PegInAddress(ctx context.Context, in *PegInAddressRequest, opts ...grpc.CallOption) (*PegInAddressReply, error)
	//
	// Claim coins from the main chain by creating a pegin transaction with the necessary metadata after the corresponding Bitcoin transaction.
	// Note that the transaction will not be relayed unless it is buried at least 102 blocks deep.
	ClaimPegIn(ctx context.Context, in *ClaimPegInRequest, opts ...grpc.CallOption) (*ClaimPegInReply, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Fee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeReply, error) {
	out := new(FeeReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/Fee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UpdateFee(ctx context.Context, in *UpdateFeeRequest, opts ...grpc.CallOption) (*UpdateFeeReply, error) {
	out := new(UpdateFeeReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/UpdateFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) SignTransaction(ctx context.Context, in *SignTransactionRequest, opts ...grpc.CallOption) (*SignTransactionReply, error) {
	out := new(SignTransactionReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/SignTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*BroadcastTransactionReply, error) {
	out := new(BroadcastTransactionReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/BroadcastTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreatePsbt(ctx context.Context, in *CreatePsbtRequest, opts ...grpc.CallOption) (*CreatePsbtReply, error) {
	out := new(CreatePsbtReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/CreatePsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) BlindPsbt(ctx context.Context, in *BlindPsbtRequest, opts ...grpc.CallOption) (*BlindPsbtReply, error) {
	out := new(BlindPsbtReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/BlindPsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) SignPsbt(ctx context.Context, in *SignPsbtRequest, opts ...grpc.CallOption) (*SignPsbtReply, error) {
	out := new(SignPsbtReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/SignPsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Mint(ctx context.Context, in *MintRequest, opts ...grpc.CallOption) (*MintReply, error) {
	out := new(MintReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/Mint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ReMint(ctx context.Context, in *ReMintRequest, opts ...grpc.CallOption) (*ReMintReply, error) {
	out := new(ReMintReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ReMint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Burn(ctx context.Context, in *BurnRequest, opts ...grpc.CallOption) (*BurnReply, error) {
	out := new(BurnReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/Burn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error) {
	out := new(TransferReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) PegInAddress(ctx context.Context, in *PegInAddressRequest, opts ...grpc.CallOption) (*PegInAddressReply, error) {
	out := new(PegInAddressReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/PegInAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ClaimPegIn(ctx context.Context, in *ClaimPegInRequest, opts ...grpc.CallOption) (*ClaimPegInReply, error) {
	out := new(ClaimPegInReply)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ClaimPegIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations should embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	// Returns fee cost of one byte in satoshi's.
	Fee(context.Context, *FeeRequest) (*FeeReply, error)
	// Updates fee cost of one byte.
	UpdateFee(context.Context, *UpdateFeeRequest) (*UpdateFeeReply, error)
	// Signs transaction hex.
	SignTransaction(context.Context, *SignTransactionRequest) (*SignTransactionReply, error)
	// Broadcast signed transaction.
	BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*BroadcastTransactionReply, error)
	// Creates unsigned, non-blinded psbt based on provided inputs and outputs.
	CreatePsbt(context.Context, *CreatePsbtRequest) (*CreatePsbtReply, error)
	// Blind unblinded psbt.
	BlindPsbt(context.Context, *BlindPsbtRequest) (*BlindPsbtReply, error)
	// Sign unsigned psbt.
	SignPsbt(context.Context, *SignPsbtRequest) (*SignPsbtReply, error)
	// Mint asset.
	Mint(context.Context, *MintRequest) (*MintReply, error)
	// ReMint asset.
	ReMint(context.Context, *ReMintRequest) (*ReMintReply, error)
	// Burn minted assets belongs to list of addresses.
	Burn(context.Context, *BurnRequest) (*BurnReply, error)
	// Transfer asset to one or more receivers.
	Transfer(context.Context, *TransferRequest) (*TransferReply, error)
	//
	// Returns information needed for claim pegin to move coins to the side-chain.
	// The user should send coins from their Bitcoin wallet to the main-chain address returned.
	PegInAddress(context.Context, *PegInAddressRequest) (*PegInAddressReply, error)
	//
	// Claim coins from the main chain by creating a pegin transaction with the necessary metadata after the corresponding Bitcoin transaction.
	// Note that the transaction will not be relayed unless it is buried at least 102 blocks deep.
	ClaimPegIn(context.Context, *ClaimPegInRequest) (*ClaimPegInReply, error)
}

// UnimplementedTransactionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) Fee(context.Context, *FeeRequest) (*FeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fee not implemented")
}
func (UnimplementedTransactionServiceServer) UpdateFee(context.Context, *UpdateFeeRequest) (*UpdateFeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFee not implemented")
}
func (UnimplementedTransactionServiceServer) SignTransaction(context.Context, *SignTransactionRequest) (*SignTransactionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*BroadcastTransactionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) CreatePsbt(context.Context, *CreatePsbtRequest) (*CreatePsbtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePsbt not implemented")
}
func (UnimplementedTransactionServiceServer) BlindPsbt(context.Context, *BlindPsbtRequest) (*BlindPsbtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlindPsbt not implemented")
}
func (UnimplementedTransactionServiceServer) SignPsbt(context.Context, *SignPsbtRequest) (*SignPsbtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignPsbt not implemented")
}
func (UnimplementedTransactionServiceServer) Mint(context.Context, *MintRequest) (*MintReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (UnimplementedTransactionServiceServer) ReMint(context.Context, *ReMintRequest) (*ReMintReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReMint not implemented")
}
func (UnimplementedTransactionServiceServer) Burn(context.Context, *BurnRequest) (*BurnReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burn not implemented")
}
func (UnimplementedTransactionServiceServer) Transfer(context.Context, *TransferRequest) (*TransferReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTransactionServiceServer) PegInAddress(context.Context, *PegInAddressRequest) (*PegInAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PegInAddress not implemented")
}
func (UnimplementedTransactionServiceServer) ClaimPegIn(context.Context, *ClaimPegInRequest) (*ClaimPegInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimPegIn not implemented")
}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_Fee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Fee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/Fee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Fee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_UpdateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).UpdateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/UpdateFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UpdateFee(ctx, req.(*UpdateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_SignTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SignTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/SignTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SignTransaction(ctx, req.(*SignTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_BroadcastTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).BroadcastTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/BroadcastTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).BroadcastTransaction(ctx, req.(*BroadcastTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreatePsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreatePsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/CreatePsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreatePsbt(ctx, req.(*CreatePsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_BlindPsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlindPsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).BlindPsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/BlindPsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).BlindPsbt(ctx, req.(*BlindPsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_SignPsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignPsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SignPsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/SignPsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SignPsbt(ctx, req.(*SignPsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/Mint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Mint(ctx, req.(*MintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ReMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReMintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ReMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ReMint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ReMint(ctx, req.(*ReMintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/Burn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Burn(ctx, req.(*BurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_PegInAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PegInAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).PegInAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/PegInAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).PegInAddress(ctx, req.(*PegInAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ClaimPegIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimPegInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ClaimPegIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ClaimPegIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ClaimPegIn(ctx, req.(*ClaimPegInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fee",
			Handler:    _TransactionService_Fee_Handler,
		},
		{
			MethodName: "UpdateFee",
			Handler:    _TransactionService_UpdateFee_Handler,
		},
		{
			MethodName: "SignTransaction",
			Handler:    _TransactionService_SignTransaction_Handler,
		},
		{
			MethodName: "BroadcastTransaction",
			Handler:    _TransactionService_BroadcastTransaction_Handler,
		},
		{
			MethodName: "CreatePsbt",
			Handler:    _TransactionService_CreatePsbt_Handler,
		},
		{
			MethodName: "BlindPsbt",
			Handler:    _TransactionService_BlindPsbt_Handler,
		},
		{
			MethodName: "SignPsbt",
			Handler:    _TransactionService_SignPsbt_Handler,
		},
		{
			MethodName: "Mint",
			Handler:    _TransactionService_Mint_Handler,
		},
		{
			MethodName: "ReMint",
			Handler:    _TransactionService_ReMint_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _TransactionService_Burn_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _TransactionService_Transfer_Handler,
		},
		{
			MethodName: "PegInAddress",
			Handler:    _TransactionService_PegInAddress_Handler,
		},
		{
			MethodName: "ClaimPegIn",
			Handler:    _TransactionService_ClaimPegIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}