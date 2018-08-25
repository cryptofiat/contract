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
public class AccountRecovery extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b50604051602080610c3a833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b610aff8061013b6000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632d1c5ff8811461007c5780633363375c146100b2578063516c4b84146100e057806373d4a13a1461011e578063f1375f3814610133578063fb55a05514610161575b600080fd5b34801561008857600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff6004358116906024351661018f565b005b3480156100be57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610322565b3480156100ec57600080fd5b506100f5610451565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012a57600080fd5b506100f561046d565b34801561013f57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610489565b34801561016d57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610496565b600082600061019d826105c5565b90506001808216146101ae57600080fd5b600481811614156101be57600080fd5b73ffffffffffffffffffffffffffffffffffffffff821615156101e057600080fd5b836101ea816106a4565b156101f457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561021657600080fd5b61021f866106ba565b73ffffffffffffffffffffffffffffffffffffffff16331461024057600080fd5b61025586600261024f896105c5565b176107a5565b60405173ffffffffffffffffffffffffffffffffffffffff8716907fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de90600090a261029f86610872565b93506102ab868561091c565b6102b58585610947565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040518082815260200191505060405180910390a3505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156103be57600080fd5b505af11580156103d2573d6000803e3d6000fd5b505050506040513d60208110156103e857600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461040a57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b610493338261096f565b50565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561053257600080fd5b505af1158015610546573d6000803e3d6000fd5b505050506040513d602081101561055c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461057e57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c0100000000000000000000000085021660248401525160009273ffffffffffffffffffffffffffffffffffffffff9092169163295f36d791604480830192602092919082900301818787803b15801561067257600080fd5b505af1158015610686573d6000803e3d6000fd5b505050506040513d602081101561069c57600080fd5b505192915050565b60006002806106b2846105c5565b161492915050565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526004818101527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008502166024820152905160009273ffffffffffffffffffffffffffffffffffffffff169163295f36d791604480830192602092919082900301818787803b15801561076357600080fd5b505af1158015610777573d6000803e3d6000fd5b505050506040513d602081101561078d57600080fd5b50516c01000000000000000000000000900492915050565b60018054604080517f461b09c000000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008602166024840152604483018490525173ffffffffffffffffffffffffffffffffffffffff9091169163461b09c091606480830192600092919082900301818387803b15801561085657600080fd5b505af115801561086a573d6000803e3d6000fd5b505050505050565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008502166024820152905160009273ffffffffffffffffffffffffffffffffffffffff169163295f36d791604480830192602092919082900301818787803b15801561067257600080fd5b600061092783610872565b90508181101561093657600080fd5b61094283838303610a22565b505050565b600061095283610872565b905081810181111561096357600080fd5b61094283838301610a22565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526004818101527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c01000000000000000000000000808702821660248401528502166044820152905173ffffffffffffffffffffffffffffffffffffffff9092169163461b09c09160648082019260009290919082900301818387803b15801561085657600080fd5b600154604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c01000000000000000000000000860216602482015260448101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163461b09c09160648082019260009290919082900301818387803b15801561085657600080fd00a165627a7a7230582092b71e579c7ae7e81b07bc0f0d7e3c6784edb9fbf4047b0d9693ed241abb8d5c0029";

    public static final String FUNC_RECOVERBALANCE = "recoverBalance";

    public static final String FUNC_SWITCHDATA = "switchData";

    public static final String FUNC_CRYPTOFIAT = "cryptoFiat";

    public static final String FUNC_DATA = "data";

    public static final String FUNC_DESIGNATERECOVERYACCOUNT = "designateRecoveryAccount";

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

    protected AccountRecovery(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected AccountRecovery(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<TransactionReceipt> recoverBalance(String from, String into) {
        final Function function = new Function(
                FUNC_RECOVERBALANCE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(from), 
                new org.web3j.abi.datatypes.Address(into)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
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

    public RemoteCall<TransactionReceipt> designateRecoveryAccount(String recoveryAccount) {
        final Function function = new Function(
                FUNC_DESIGNATERECOVERYACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(recoveryAccount)), 
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

    public static RemoteCall<AccountRecovery> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat)));
        return deployRemoteCall(AccountRecovery.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static RemoteCall<AccountRecovery> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat)));
        return deployRemoteCall(AccountRecovery.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
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

    public static AccountRecovery load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new AccountRecovery(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static AccountRecovery load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new AccountRecovery(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
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
