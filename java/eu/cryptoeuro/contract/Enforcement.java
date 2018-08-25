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
public class Enforcement extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b50604051608080610e83833981016040908152815160208301519183015160609093015160008054600160a060020a031916600160a060020a0384161790559092906100636401000000006100a6810204565b60028054600160a060020a03948516600160a060020a0319918216179091556003805493851693821693909317909255600480549190931691161790555061017e565b6100b960016401000000006100e5810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100e357600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561014c57600080fd5b505af1158015610160573d6000803e3d6000fd5b505050506040513d602081101561017657600080fd5b505192915050565b610cf68061018d6000396000f3006080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100c9578063516c4b84146100f95780635dab24201461013757806372cfc9dc1461014c57806373d4a13a14610161578063788649ea1461017657806385a0f282146101a457806390f28e74146101b9578063b10725e8146101e7578063b9b0330f14610215578063f26c159f14610243578063f3fef3a314610271578063fb55a055146102a2575b600080fd5b3480156100d557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff600435166102d0565b005b34801561010557600080fd5b5061010e6103ff565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561014357600080fd5b5061010e61041b565b34801561015857600080fd5b5061010e610437565b34801561016d57600080fd5b5061010e610453565b34801561018257600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661046f565b3480156101b057600080fd5b5061010e610517565b3480156101c557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610533565b3480156101f357600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661059e565b34801561022157600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610609565b34801561024f57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff600435166106ab565b34801561027d57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516602435610734565b3480156102ae57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661082b565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561036c57600080fd5b505af1158015610380573d6000803e3d6000fd5b505050506040513d602081101561039657600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146103b857600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff16331461049357600080fd5b6104c7817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb6104c18261095a565b16610a39565b6040805160008152905173ffffffffffffffffffffffffffffffffffffffff8316917fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee919081900360200190a250565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff16331461055757600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff1633146105c257600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60035473ffffffffffffffffffffffffffffffffffffffff16331461062d57600080fd5b8061063781610b06565b1561064157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561066357600080fd5b50600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff1633146106cf57600080fd5b6106e48160046106de8461095a565b17610a39565b6040805160018152905173ffffffffffffffffffffffffffffffffffffffff8316917fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee919081900360200190a250565b60025473ffffffffffffffffffffffffffffffffffffffff16331461075857600080fd5b60045473ffffffffffffffffffffffffffffffffffffffff1661077a81610b06565b1561078457600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156107a657600080fd5b6107b08383610b1c565b6004546107d39073ffffffffffffffffffffffffffffffffffffffff1683610b47565b60045460408051848152905173ffffffffffffffffffffffffffffffffffffffff928316928616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156108c757600080fd5b505af11580156108db573d6000803e3d6000fd5b505050506040513d60208110156108f157600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461091357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c0100000000000000000000000085021660248401525160009273ffffffffffffffffffffffffffffffffffffffff9092169163295f36d791604480830192602092919082900301818787803b158015610a0757600080fd5b505af1158015610a1b573d6000803e3d6000fd5b505050506040513d6020811015610a3157600080fd5b505192915050565b60018054604080517f461b09c000000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008602166024840152604483018490525173ffffffffffffffffffffffffffffffffffffffff9091169163461b09c091606480830192600092919082900301818387803b158015610aea57600080fd5b505af1158015610afe573d6000803e3d6000fd5b505050505050565b6000600280610b148461095a565b161492915050565b6000610b2783610b6f565b905081811015610b3657600080fd5b610b4283838303610c19565b505050565b6000610b5283610b6f565b9050818101811115610b6357600080fd5b610b4283838301610c19565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008502166024820152905160009273ffffffffffffffffffffffffffffffffffffffff169163295f36d791604480830192602092919082900301818787803b158015610a0757600080fd5b600154604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c01000000000000000000000000860216602482015260448101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163461b09c09160648082019260009290919082900301818387803b158015610aea57600080fd00a165627a7a72305820f3997157b03ff7d1aa485e858588bd2f11df145f94ffa3b56b901002f94197a50029";

    public static final String FUNC_SWITCHDATA = "switchData";

    public static final String FUNC_CRYPTOFIAT = "cryptoFiat";

    public static final String FUNC_ACCOUNT = "account";

    public static final String FUNC_LAWENFORCER = "lawEnforcer";

    public static final String FUNC_DATA = "data";

    public static final String FUNC_UNFREEZEACCOUNT = "unfreezeAccount";

    public static final String FUNC_ACCOUNTDESIGNATOR = "accountDesignator";

    public static final String FUNC_APPOINTACCOUNTDESIGNATOR = "appointAccountDesignator";

    public static final String FUNC_APPOINTLAWENFORCER = "appointLawEnforcer";

    public static final String FUNC_DESIGNATEACCOUNT = "designateAccount";

    public static final String FUNC_FREEZEACCOUNT = "freezeAccount";

    public static final String FUNC_WITHDRAW = "withdraw";

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

    protected Enforcement(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Enforcement(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
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

    public RemoteCall<String> account() {
        final Function function = new Function(FUNC_ACCOUNT, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> lawEnforcer() {
        final Function function = new Function(FUNC_LAWENFORCER, 
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

    public RemoteCall<TransactionReceipt> unfreezeAccount(String target) {
        final Function function = new Function(
                FUNC_UNFREEZEACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(target)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> accountDesignator() {
        final Function function = new Function(FUNC_ACCOUNTDESIGNATOR, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<TransactionReceipt> appointAccountDesignator(String next) {
        final Function function = new Function(
                FUNC_APPOINTACCOUNTDESIGNATOR, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> appointLawEnforcer(String next) {
        final Function function = new Function(
                FUNC_APPOINTLAWENFORCER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> designateAccount(String _account) {
        final Function function = new Function(
                FUNC_DESIGNATEACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_account)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> freezeAccount(String target) {
        final Function function = new Function(
                FUNC_FREEZEACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(target)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> withdraw(String from, BigInteger amount) {
        final Function function = new Function(
                FUNC_WITHDRAW, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(from), 
                new org.web3j.abi.datatypes.generated.Uint256(amount)), 
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

    public static RemoteCall<Enforcement> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat, String _lawEnforcer, String _enforcementAccountDesignator, String _enforcementAccount) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat), 
                new org.web3j.abi.datatypes.Address(_lawEnforcer), 
                new org.web3j.abi.datatypes.Address(_enforcementAccountDesignator), 
                new org.web3j.abi.datatypes.Address(_enforcementAccount)));
        return deployRemoteCall(Enforcement.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static RemoteCall<Enforcement> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat, String _lawEnforcer, String _enforcementAccountDesignator, String _enforcementAccount) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat), 
                new org.web3j.abi.datatypes.Address(_lawEnforcer), 
                new org.web3j.abi.datatypes.Address(_enforcementAccountDesignator), 
                new org.web3j.abi.datatypes.Address(_enforcementAccount)));
        return deployRemoteCall(Enforcement.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
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

    public static Enforcement load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new Enforcement(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static Enforcement load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Enforcement(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
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
