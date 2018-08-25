package eu.cryptoeuro.contract;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import rx.Observable;
import rx.functions.Func1;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 3.5.0.
 */
public class Reserve extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b50604051604080610d5283398101604052805160209091015160008054600160a060020a031916600160a060020a03841617905561005564010000000061007b810204565b60028054600160a060020a031916600160a060020a039290921691909117905550610153565b61008e60016401000000006100ba810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100b857600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561012157600080fd5b505af1158015610135573d6000803e3d6000fd5b505050506040513d602081101561014b57600080fd5b505192915050565b610bf0806101626000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302946804811461009d57806318160ddd146100db5780633363375c14610102578063516c4b841461013257806373d4a13a1461014757806398e52f9a1461015c578063b921e16314610174578063ddf05f591461018c578063fb55a055146101ba575b600080fd5b3480156100a957600080fd5b506100b26101e8565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100e757600080fd5b506100f0610204565b60408051918252519081900360200190f35b34801561010e57600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff60043516610213565b005b34801561013e57600080fd5b506100b2610342565b34801561015357600080fd5b506100b261035e565b34801561016857600080fd5b5061013060043561037a565b34801561018057600080fd5b506101306004356104da565b34801561019857600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff60043516610623565b3480156101c657600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff6004351661068e565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b600061020e6107bd565b905090565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156102af57600080fd5b505af11580156102c3573d6000803e3d6000fd5b505050506040513d60208110156102d957600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146102fb57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025460009073ffffffffffffffffffffffffffffffffffffffff1633146103a157600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff1660006103c58261086a565b90506001808216146103d657600080fd5b600481811614156103e657600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561040857600080fd5b6104106107bd565b92508383101561041f57600080fd5b838303925061042d83610949565b6002546104509073ffffffffffffffffffffffffffffffffffffffff16856109e4565b60025460408051868152905160009273ffffffffffffffffffffffffffffffffffffffff16917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a36040805184815290517ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e99181900360200190a150505050565b60025460009073ffffffffffffffffffffffffffffffffffffffff16331461050157600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff1661052381610a0f565b1561052d57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561054f57600080fd5b610557610204565b915082820182111561056857600080fd5b9082019061057582610949565b6002546105989073ffffffffffffffffffffffffffffffffffffffff1684610a25565b60025460408051858152905173ffffffffffffffffffffffffffffffffffffffff909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a36040805183815290517ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e99181900360200190a1505050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461064757600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561072a57600080fd5b505af115801561073e573d6000803e3d6000fd5b505050506040513d602081101561075457600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461077657600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600560048201526000602482018190529151919273ffffffffffffffffffffffffffffffffffffffff169163295f36d79160448082019260209290919082900301818787803b15801561083957600080fd5b505af115801561084d573d6000803e3d6000fd5b505050506040513d602081101561086357600080fd5b5051905090565b60018054604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c0100000000000000000000000085021660248401525160009273ffffffffffffffffffffffffffffffffffffffff9092169163295f36d791604480830192602092919082900301818787803b15801561091757600080fd5b505af115801561092b573d6000803e3d6000fd5b505050506040513d602081101561094157600080fd5b505192915050565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526005600482015260006024820181905260448201859052915173ffffffffffffffffffffffffffffffffffffffff9093169263461b09c09260648084019391929182900301818387803b1580156109c957600080fd5b505af11580156109dd573d6000803e3d6000fd5b5050505050565b60006109ef83610a4d565b9050818110156109fe57600080fd5b610a0a83838303610af7565b505050565b6000600280610a1d8461086a565b161492915050565b6000610a3083610a4d565b9050818101811115610a4157600080fd5b610a0a83838301610af7565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008502166024820152905160009273ffffffffffffffffffffffffffffffffffffffff169163295f36d791604480830192602092919082900301818787803b15801561091757600080fd5b600154604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c01000000000000000000000000860216602482015260448101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163461b09c09160648082019260009290919082900301818387803b158015610ba857600080fd5b505af1158015610bbc573d6000803e3d6000fd5b5050505050505600a165627a7a7230582003b8b01f830aade202fe18f5413bbcdbe48259989e2d898f44e87e62c9c1ba650029";

    public static final String FUNC_RESERVEBANK = "reserveBank";

    public static final String FUNC_TOTALSUPPLY = "totalSupply";

    public static final String FUNC_SWITCHDATA = "switchData";

    public static final String FUNC_CRYPTOFIAT = "cryptoFiat";

    public static final String FUNC_DATA = "data";

    public static final String FUNC_DECREASESUPPLY = "decreaseSupply";

    public static final String FUNC_INCREASESUPPLY = "increaseSupply";

    public static final String FUNC_APPOINTRESERVEBANK = "appointReserveBank";

    public static final String FUNC_SWITCHCRYPTOFIAT = "switchCryptoFiat";

    public static final Event TRANSFER_EVENT = new Event("Transfer", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Uint256>() {}));
    ;

    public static final Event ACCOUNTAPPROVED_EVENT = new Event("AccountApproved", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    public static final Event ACCOUNTCLOSED_EVENT = new Event("AccountClosed", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    public static final Event ACCOUNTFREEZE_EVENT = new Event("AccountFreeze", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Bool>() {}));
    ;

    public static final Event SUPPLYCHANGED_EVENT = new Event("SupplyChanged", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
    ;

    protected Reserve(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Reserve(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<String> reserveBank() {
        final Function function = new Function(FUNC_RESERVEBANK, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<BigInteger> totalSupply() {
        final Function function = new Function(FUNC_TOTALSUPPLY, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<TransactionReceipt> switchData(String next) {
        final Function function = new Function(
                FUNC_SWITCHDATA, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> cryptoFiat() {
        final Function function = new Function(FUNC_CRYPTOFIAT, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> data() {
        final Function function = new Function(FUNC_DATA, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<TransactionReceipt> decreaseSupply(BigInteger amount) {
        final Function function = new Function(
                FUNC_DECREASESUPPLY, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(amount)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> increaseSupply(BigInteger amount) {
        final Function function = new Function(
                FUNC_INCREASESUPPLY, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(amount)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> appointReserveBank(String next) {
        final Function function = new Function(
                FUNC_APPOINTRESERVEBANK, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> switchCryptoFiat(String next) {
        final Function function = new Function(
                FUNC_SWITCHCRYPTOFIAT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public static RemoteCall<Reserve> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat, String _reserveBank) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat), 
                new org.web3j.abi.datatypes.Address(_reserveBank)));
        return deployRemoteCall(Reserve.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static RemoteCall<Reserve> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat, String _reserveBank) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat), 
                new org.web3j.abi.datatypes.Address(_reserveBank)));
        return deployRemoteCall(Reserve.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public List<TransferEventResponse> getTransferEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(TRANSFER_EVENT, transactionReceipt);
        ArrayList<TransferEventResponse> responses = new ArrayList<TransferEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            TransferEventResponse typedResponse = new TransferEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.destination = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse.amount = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<TransferEventResponse> transferEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, TransferEventResponse>() {
            @Override
            public TransferEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(TRANSFER_EVENT, log);
                TransferEventResponse typedResponse = new TransferEventResponse();
                typedResponse.log = log;
                typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.destination = (String) eventValues.getIndexedValues().get(1).getValue();
                typedResponse.amount = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<TransferEventResponse> transferEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(TRANSFER_EVENT));
        return transferEventObservable(filter);
    }

    public List<AccountApprovedEventResponse> getAccountApprovedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(ACCOUNTAPPROVED_EVENT, transactionReceipt);
        ArrayList<AccountApprovedEventResponse> responses = new ArrayList<AccountApprovedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            AccountApprovedEventResponse typedResponse = new AccountApprovedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<AccountApprovedEventResponse> accountApprovedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, AccountApprovedEventResponse>() {
            @Override
            public AccountApprovedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(ACCOUNTAPPROVED_EVENT, log);
                AccountApprovedEventResponse typedResponse = new AccountApprovedEventResponse();
                typedResponse.log = log;
                typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<AccountApprovedEventResponse> accountApprovedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(ACCOUNTAPPROVED_EVENT));
        return accountApprovedEventObservable(filter);
    }

    public List<AccountClosedEventResponse> getAccountClosedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(ACCOUNTCLOSED_EVENT, transactionReceipt);
        ArrayList<AccountClosedEventResponse> responses = new ArrayList<AccountClosedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            AccountClosedEventResponse typedResponse = new AccountClosedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<AccountClosedEventResponse> accountClosedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, AccountClosedEventResponse>() {
            @Override
            public AccountClosedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(ACCOUNTCLOSED_EVENT, log);
                AccountClosedEventResponse typedResponse = new AccountClosedEventResponse();
                typedResponse.log = log;
                typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<AccountClosedEventResponse> accountClosedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(ACCOUNTCLOSED_EVENT));
        return accountClosedEventObservable(filter);
    }

    public List<AccountFreezeEventResponse> getAccountFreezeEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(ACCOUNTFREEZE_EVENT, transactionReceipt);
        ArrayList<AccountFreezeEventResponse> responses = new ArrayList<AccountFreezeEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            AccountFreezeEventResponse typedResponse = new AccountFreezeEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.frozen = (Boolean) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<AccountFreezeEventResponse> accountFreezeEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, AccountFreezeEventResponse>() {
            @Override
            public AccountFreezeEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(ACCOUNTFREEZE_EVENT, log);
                AccountFreezeEventResponse typedResponse = new AccountFreezeEventResponse();
                typedResponse.log = log;
                typedResponse.source = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.frozen = (Boolean) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<AccountFreezeEventResponse> accountFreezeEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(ACCOUNTFREEZE_EVENT));
        return accountFreezeEventObservable(filter);
    }

    public List<SupplyChangedEventResponse> getSupplyChangedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(SUPPLYCHANGED_EVENT, transactionReceipt);
        ArrayList<SupplyChangedEventResponse> responses = new ArrayList<SupplyChangedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            SupplyChangedEventResponse typedResponse = new SupplyChangedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.totalSupply = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<SupplyChangedEventResponse> supplyChangedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, SupplyChangedEventResponse>() {
            @Override
            public SupplyChangedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(SUPPLYCHANGED_EVENT, log);
                SupplyChangedEventResponse typedResponse = new SupplyChangedEventResponse();
                typedResponse.log = log;
                typedResponse.totalSupply = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<SupplyChangedEventResponse> supplyChangedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(SUPPLYCHANGED_EVENT));
        return supplyChangedEventObservable(filter);
    }

    public static Reserve load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new Reserve(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static Reserve load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Reserve(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static class TransferEventResponse {
        public Log log;

        public String source;

        public String destination;

        public BigInteger amount;
    }

    public static class AccountApprovedEventResponse {
        public Log log;

        public String source;
    }

    public static class AccountClosedEventResponse {
        public Log log;

        public String source;
    }

    public static class AccountFreezeEventResponse {
        public Log log;

        public String source;

        public Boolean frozen;
    }

    public static class SupplyChangedEventResponse {
        public Log log;

        public BigInteger totalSupply;
    }
}
