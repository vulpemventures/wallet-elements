# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from ocean.v1alpha import transaction_pb2 as ocean_dot_v1alpha_dot_transaction__pb2


class TransactionServiceStub(object):
    """
    TransactionService is used to craft and sign various kind's of transactions.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SelectUtxos = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/SelectUtxos',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.SelectUtxosRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.SelectUtxosResponse.FromString,
                )
        self.EstimateFees = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/EstimateFees',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.EstimateFeesRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.EstimateFeesResponse.FromString,
                )
        self.SignTransaction = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/SignTransaction',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.SignTransactionRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.SignTransactionResponse.FromString,
                )
        self.BroadcastTransaction = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/BroadcastTransaction',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.BroadcastTransactionRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.BroadcastTransactionResponse.FromString,
                )
        self.CreatePset = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/CreatePset',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.CreatePsetRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.CreatePsetResponse.FromString,
                )
        self.UpdatePset = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/UpdatePset',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.UpdatePsetRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.UpdatePsetResponse.FromString,
                )
        self.BlindPset = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/BlindPset',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.BlindPsetRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.BlindPsetResponse.FromString,
                )
        self.SignPset = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/SignPset',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.SignPsetRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.SignPsetResponse.FromString,
                )
        self.Mint = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/Mint',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.MintRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.MintResponse.FromString,
                )
        self.Remint = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/Remint',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.RemintRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.RemintResponse.FromString,
                )
        self.Burn = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/Burn',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.BurnRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.BurnResponse.FromString,
                )
        self.Transfer = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/Transfer',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.TransferRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.TransferResponse.FromString,
                )
        self.PegInAddress = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/PegInAddress',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.PegInAddressRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.PegInAddressResponse.FromString,
                )
        self.ClaimPegIn = channel.unary_unary(
                '/ocean.v1alpha.TransactionService/ClaimPegIn',
                request_serializer=ocean_dot_v1alpha_dot_transaction__pb2.ClaimPegInRequest.SerializeToString,
                response_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.ClaimPegInResponse.FromString,
                )


class TransactionServiceServicer(object):
    """
    TransactionService is used to craft and sign various kind's of transactions.
    """

    def SelectUtxos(self, request, context):
        """SelectUtxos returns a selction of utxos, to be used in another 
        transaction, for provided target amount and strategy.
        Selected utxos are locked for predefined amount of time to prevent 
        double-spending them.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def EstimateFees(self, request, context):
        """EstimateFees returns the fee amount to pay for a tx containing the given 
        inputs and outputs.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SignTransaction(self, request, context):
        """SignTransaction signs a raw transaction in hex format.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BroadcastTransaction(self, request, context):
        """BroadcastTransaction broadacats a raw transaction in hex format.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreatePset(self, request, context):
        """CreatePset returns an unsigned pset for given inputs and outputs.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdatePset(self, request, context):
        """UpdatePset adds the given inputs and outputs to the partial transaction.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BlindPset(self, request, context):
        """BlindPset updates the given pset with required ins and outs blinded.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SignPset(self, request, context):
        """SignPset updates the given pset adding the required signatures.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Mint(self, request, context):
        """Mint returns a transaction that issues a new asset.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Remint(self, request, context):
        """Remint returns a transaction that re-issues an existing asset.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Burn(self, request, context):
        """Burn returns a transaction that burns some funds.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Transfer(self, request, context):
        """Transfer returns a transaction to send funds to some receiver.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PegInAddress(self, request, context):
        """PegInAddress returns what's necessary to peg funds of the Bitcoin 
        main-chain and have them available on the Liquid side-chain.
        Bitcoin funds must be sent to the main-chain address while the claim
        output script must be used to redeem the LBTC ones.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ClaimPegIn(self, request, context):
        """ClaimPegIn returns a transaction to claim funds pegged on the Bitcoin 
        main-chain to have them available on the Liquid side-chain.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TransactionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SelectUtxos': grpc.unary_unary_rpc_method_handler(
                    servicer.SelectUtxos,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.SelectUtxosRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.SelectUtxosResponse.SerializeToString,
            ),
            'EstimateFees': grpc.unary_unary_rpc_method_handler(
                    servicer.EstimateFees,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.EstimateFeesRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.EstimateFeesResponse.SerializeToString,
            ),
            'SignTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.SignTransaction,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.SignTransactionRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.SignTransactionResponse.SerializeToString,
            ),
            'BroadcastTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.BroadcastTransaction,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.BroadcastTransactionRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.BroadcastTransactionResponse.SerializeToString,
            ),
            'CreatePset': grpc.unary_unary_rpc_method_handler(
                    servicer.CreatePset,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.CreatePsetRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.CreatePsetResponse.SerializeToString,
            ),
            'UpdatePset': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdatePset,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.UpdatePsetRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.UpdatePsetResponse.SerializeToString,
            ),
            'BlindPset': grpc.unary_unary_rpc_method_handler(
                    servicer.BlindPset,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.BlindPsetRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.BlindPsetResponse.SerializeToString,
            ),
            'SignPset': grpc.unary_unary_rpc_method_handler(
                    servicer.SignPset,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.SignPsetRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.SignPsetResponse.SerializeToString,
            ),
            'Mint': grpc.unary_unary_rpc_method_handler(
                    servicer.Mint,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.MintRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.MintResponse.SerializeToString,
            ),
            'Remint': grpc.unary_unary_rpc_method_handler(
                    servicer.Remint,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.RemintRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.RemintResponse.SerializeToString,
            ),
            'Burn': grpc.unary_unary_rpc_method_handler(
                    servicer.Burn,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.BurnRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.BurnResponse.SerializeToString,
            ),
            'Transfer': grpc.unary_unary_rpc_method_handler(
                    servicer.Transfer,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.TransferRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.TransferResponse.SerializeToString,
            ),
            'PegInAddress': grpc.unary_unary_rpc_method_handler(
                    servicer.PegInAddress,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.PegInAddressRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.PegInAddressResponse.SerializeToString,
            ),
            'ClaimPegIn': grpc.unary_unary_rpc_method_handler(
                    servicer.ClaimPegIn,
                    request_deserializer=ocean_dot_v1alpha_dot_transaction__pb2.ClaimPegInRequest.FromString,
                    response_serializer=ocean_dot_v1alpha_dot_transaction__pb2.ClaimPegInResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ocean.v1alpha.TransactionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TransactionService(object):
    """
    TransactionService is used to craft and sign various kind's of transactions.
    """

    @staticmethod
    def SelectUtxos(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/SelectUtxos',
            ocean_dot_v1alpha_dot_transaction__pb2.SelectUtxosRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.SelectUtxosResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def EstimateFees(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/EstimateFees',
            ocean_dot_v1alpha_dot_transaction__pb2.EstimateFeesRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.EstimateFeesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SignTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/SignTransaction',
            ocean_dot_v1alpha_dot_transaction__pb2.SignTransactionRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.SignTransactionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BroadcastTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/BroadcastTransaction',
            ocean_dot_v1alpha_dot_transaction__pb2.BroadcastTransactionRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.BroadcastTransactionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreatePset(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/CreatePset',
            ocean_dot_v1alpha_dot_transaction__pb2.CreatePsetRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.CreatePsetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdatePset(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/UpdatePset',
            ocean_dot_v1alpha_dot_transaction__pb2.UpdatePsetRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.UpdatePsetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BlindPset(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/BlindPset',
            ocean_dot_v1alpha_dot_transaction__pb2.BlindPsetRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.BlindPsetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SignPset(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/SignPset',
            ocean_dot_v1alpha_dot_transaction__pb2.SignPsetRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.SignPsetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Mint(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/Mint',
            ocean_dot_v1alpha_dot_transaction__pb2.MintRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.MintResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Remint(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/Remint',
            ocean_dot_v1alpha_dot_transaction__pb2.RemintRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.RemintResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Burn(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/Burn',
            ocean_dot_v1alpha_dot_transaction__pb2.BurnRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.BurnResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Transfer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/Transfer',
            ocean_dot_v1alpha_dot_transaction__pb2.TransferRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.TransferResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PegInAddress(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/PegInAddress',
            ocean_dot_v1alpha_dot_transaction__pb2.PegInAddressRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.PegInAddressResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ClaimPegIn(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ocean.v1alpha.TransactionService/ClaimPegIn',
            ocean_dot_v1alpha_dot_transaction__pb2.ClaimPegInRequest.SerializeToString,
            ocean_dot_v1alpha_dot_transaction__pb2.ClaimPegInResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)