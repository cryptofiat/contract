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
public class Delegation extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b506040516020806112f3833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6111b88061013b6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305bafaa481146100875780633363375c146100ff578063516c4b841461012d57806373d4a13a1461016b578063e218e6d214610180578063ed2a2d6414610219578063fb55a05514610259575b600080fd5b34801561009357600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100fd9583359536956044949193909101919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff16935061028792505050565b005b34801561010b57600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff6004351661045d565b34801561013957600080fd5b5061014261058c565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017757600080fd5b506101426105a8565b34801561018c57600080fd5b50604080516020600460843581810135601f81018490048402850184019095528484526100fd948235946024803573ffffffffffffffffffffffffffffffffffffffff169560443595606435953695919460a49490939101919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff1693506105c492505050565b34801561022557600080fd5b5061024773ffffffffffffffffffffffffffffffffffffffff60043516610833565b60408051918252519081900360200190f35b34801561026557600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff60043516610844565b6000610291611130565b8261029b81610973565b156102a557600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156102c757600080fd5b600092505b85831015610455576102de8584610989565b91506102ed8260200151610bc8565b6102fa8260400151610c1c565b8151602083015161030a90610c56565b1061031457600080fd5b61032682602001518360000151610d32565b61033e82602001518360800151846060015101610df7565b61035082604001518360600151610e1d565b816040015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84606001516040518082815260200191505060405180910390a360008260800151111561044a576103dc848360800151610e1d565b8373ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84608001516040518082815260200191505060405180910390a35b6001909201916102cc565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156104f957600080fd5b505af115801561050d573d6000803e3d6000fd5b505050506040513d602081101561052357600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461054557600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6000856105d081610973565b156105da57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156105fc57600080fd5b8261060681610973565b1561061057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561063257600080fd5b6040805160208082018c90526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8c160282840152605482018a905260748083018a905283518084039091018152609490920192839052815161070b93918291908401908083835b602083106106d857805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161069b565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902086610e45565b925061071683610bc8565b8861072084610c56565b1061072a57600080fd5b610734838a610d32565b61074083878901610df7565b61074a8888610e1d565b8773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef896040518082815260200191505060405180910390a36000861115610828576107c28487610e1d565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef886040518082815260200191505060405180910390a35b505050505050505050565b600061083e82610c56565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156108e057600080fd5b505af11580156108f4573d6000803e3d6000fd5b505050506040513d602081101561090a57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461092c57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060028061098184610f28565b161492915050565b610991611130565b60008060008060008060008060006109a7611130565b8b60b5029950898d019c5060208d0151985060348d0151975060548d0151965060748d0151955060948d0151945060b48d0151935060ff60b58e0151169250601b8360ff1610156109f957601b830192505b6040805160208082018c90526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8c160282840152605482018a905260748083018a905283518084039091018152609490920192839052815191929182918401908083835b60208310610a9c57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610a5f565b5181517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209485036101000a0190811690199190911617905260408051949092018490039093208e87528151600080825281860180855283905260ff8b1682850152606082018d9052608082018c905292519198506001965060a080820196507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083019450908290030191865af1158015610b5c573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015173ffffffffffffffffffffffffffffffffffffffff90811660208501528a1690830152506060810187905260808101869052995050505050505050505092915050565b806000610bd482610f28565b9050600180821614610be557600080fd5b60048181161415610bf557600080fd5b73ffffffffffffffffffffffffffffffffffffffff82161515610c1757600080fd5b505050565b80610c2681610973565b15610c3057600080fd5b73ffffffffffffffffffffffffffffffffffffffff81161515610c5257600080fd5b5050565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600360048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008502166024820152905160009273ffffffffffffffffffffffffffffffffffffffff169163295f36d791604480830192602092919082900301818787803b158015610d0057600080fd5b505af1158015610d14573d6000803e3d6000fd5b505050506040513d6020811015610d2a57600080fd5b505192915050565b600154604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600360048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c01000000000000000000000000860216602482015260448101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163461b09c09160648082019260009290919082900301818387803b158015610de357600080fd5b505af1158015610455573d6000803e3d6000fd5b6000610e0283610fd5565b905081811015610e1157600080fd5b610c178383830361107f565b6000610e2883610fd5565b9050818101811115610e3957600080fd5b610c178383830161107f565b60008060008084516041141515610e5b57600080fd5b50505060208201516040830151604184015160ff16601b811015610e7d57601b015b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a08085019491937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840193928390039091019190865af1158015610ef5573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00151979650505050505050565b60018054604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048101939093527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c0100000000000000000000000085021660248401525160009273ffffffffffffffffffffffffffffffffffffffff9092169163295f36d791604480830192602092919082900301818787803b158015610d0057600080fd5b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c010000000000000000000000008502166024820152905160009273ffffffffffffffffffffffffffffffffffffffff169163295f36d791604480830192602092919082900301818787803b158015610d0057600080fd5b600154604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600260048201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000006c01000000000000000000000000860216602482015260448101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163461b09c09160648082019260009290919082900301818387803b158015610de357600080fd5b60a06040519081016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815250905600a165627a7a72305820583608dc7dfc1e801032f7df6c4b5fb18d81928b0b06383afdb9a8d4e5b31ea80029";

    public static final String FUNC_MULTITRANSFER = "multitransfer";

    public static final String FUNC_SWITCHDATA = "switchData";

    public static final String FUNC_CRYPTOFIAT = "cryptoFiat";

    public static final String FUNC_DATA = "data";

    public static final String FUNC_TRANSFER = "transfer";

    public static final String FUNC_NONCEOF = "nonceOf";

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

    protected Delegation(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Delegation(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<TransactionReceipt> multitransfer(BigInteger count, byte[] transfers, String delegate) {
        final Function function = new Function(
                FUNC_MULTITRANSFER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(count), 
                new org.web3j.abi.datatypes.DynamicBytes(transfers), 
                new org.web3j.abi.datatypes.Address(delegate)), 
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

    public RemoteCall<TransactionReceipt> transfer(BigInteger nonce, String destination, BigInteger amount, BigInteger fee, byte[] signature, String delegate) {
        final Function function = new Function(
                FUNC_TRANSFER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(nonce), 
                new org.web3j.abi.datatypes.Address(destination), 
                new org.web3j.abi.datatypes.generated.Uint256(amount), 
                new org.web3j.abi.datatypes.generated.Uint256(fee), 
                new org.web3j.abi.datatypes.DynamicBytes(signature), 
                new org.web3j.abi.datatypes.Address(delegate)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<BigInteger> nonceOf(String account) {
        final Function function = new Function(FUNC_NONCEOF, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<TransactionReceipt> switchCryptoFiat(String next) {
        final Function function = new Function(
                FUNC_SWITCHCRYPTOFIAT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public static RemoteCall<Delegation> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat)));
        return deployRemoteCall(Delegation.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static RemoteCall<Delegation> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat)));
        return deployRemoteCall(Delegation.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
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

    public static Delegation load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new Delegation(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static Delegation load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Delegation(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
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
