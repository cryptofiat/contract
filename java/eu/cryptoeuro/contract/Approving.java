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
public class Approving extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b506040516040806109c583398101604052805160209091015160008054600160a060020a031916600160a060020a03841617905561005564010000000061007b810204565b60028054600160a060020a031916600160a060020a039290921691909117905550610153565b61008e60016401000000006100ba810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100b857600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561012157600080fd5b505af1158015610135573d6000803e3d6000fd5b505050506040513d602081101561014b57600080fd5b505192915050565b610863806101626000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663071a8b53811461009d57806307a385e6146100f45780633363375c14610132578063516c4b841461016057806373d4a13a14610175578063c8b091091461018a578063dd336b94146101b8578063f89f4e77146101e6578063fb55a05514610214575b600080fd5b3480156100a957600080fd5b50604080516020600480358082013583810280860185019096528085526100f2953695939460249493850192918291850190849080828437509497506102429650505050505050565b005b34801561010057600080fd5b5061010961027a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561013e57600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610296565b34801561016c57600080fd5b506101096103c5565b34801561018157600080fd5b506101096103e1565b34801561019657600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff600435166103fd565b3480156101c457600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610468565b3480156101f257600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff600435166104e5565b34801561022057600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff6004351661055c565b60005b81518110156102765761026e828281518110151561025f57fe5b906020019060200201516104e5565b600101610245565b5050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561033257600080fd5b505af1158015610346573d6000803e3d6000fd5b505050506040513d602081101561035c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461037e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff16331461042157600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff16331461048c57600080fd5b6104a181600261049b8461068b565b1761076a565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de90600090a250565b60025473ffffffffffffffffffffffffffffffffffffffff16331461050957600080fd5b61051881600161049b8461068b565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573090600090a250565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156105f857600080fd5b505af115801561060c573d6000803e3d6000fd5b505050506040513d602081101561062257600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461064457600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c0100000000000000000000000085021660248401525160009273ffffffffffffffffffffffffffffffffffffffff9092169163295f36d791604480830192602092919082900301818787803b15801561073857600080fd5b505af115801561074c573d6000803e3d6000fd5b505050506040513d602081101561076257600080fd5b505192915050565b60018054604080517f461b09c000000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008602166024840152604483018490525173ffffffffffffffffffffffffffffffffffffffff9091169163461b09c091606480830192600092919082900301818387803b15801561081b57600080fd5b505af115801561082f573d6000803e3d6000fd5b5050505050505600a165627a7a723058208b010ce29ef6eb69162af7640cc40fae8715b0dbd3ea24a3ce8e0ce36931e4e60029";

    public static final String FUNC_APPROVEACCOUNTS = "approveAccounts";

    public static final String FUNC_ACCOUNTAPPROVER = "accountApprover";

    public static final String FUNC_SWITCHDATA = "switchData";

    public static final String FUNC_CRYPTOFIAT = "cryptoFiat";

    public static final String FUNC_DATA = "data";

    public static final String FUNC_APPOINTACCOUNTAPPROVER = "appointAccountApprover";

    public static final String FUNC_CLOSEACCOUNT = "closeAccount";

    public static final String FUNC_APPROVEACCOUNT = "approveAccount";

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

    protected Approving(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Approving(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<TransactionReceipt> approveAccounts(List<String> accounts) {
        final Function function = new Function(
                FUNC_APPROVEACCOUNTS, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.Address>(
                        org.web3j.abi.Utils.typeMap(accounts, org.web3j.abi.datatypes.Address.class))), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> accountApprover() {
        final Function function = new Function(FUNC_ACCOUNTAPPROVER, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
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

    public RemoteCall<TransactionReceipt> appointAccountApprover(String next) {
        final Function function = new Function(
                FUNC_APPOINTACCOUNTAPPROVER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> closeAccount(String account) {
        final Function function = new Function(
                FUNC_CLOSEACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> approveAccount(String account) {
        final Function function = new Function(
                FUNC_APPROVEACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
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

    public static RemoteCall<Approving> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat, String _accountApprover) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat), 
                new org.web3j.abi.datatypes.Address(_accountApprover)));
        return deployRemoteCall(Approving.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static RemoteCall<Approving> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat, String _accountApprover) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat), 
                new org.web3j.abi.datatypes.Address(_accountApprover)));
        return deployRemoteCall(Approving.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
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

    public static Approving load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new Approving(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static Approving load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Approving(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
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
