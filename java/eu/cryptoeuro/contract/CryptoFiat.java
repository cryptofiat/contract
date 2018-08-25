package eu.cryptoeuro.contract;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
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
public class CryptoFiat extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b5060008054600160a060020a0319908116331782556003805460018101825592527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9091018054909116301790556105468061006d6000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009257806313c01368146100c55780633fad74ad14610106578063474da79a1461012d5780635db4380d14610145578063874c3473146101735780639afd453c146101a1578063e814861e146101b6575b600080fd5b34801561009e57600080fd5b506100c360043573ffffffffffffffffffffffffffffffffffffffff602435166101f8565b005b3480156100d157600080fd5b506100dd6004356103f5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561011257600080fd5b5061011b61041d565b60408051918252519081900360200190f35b34801561013957600080fd5b506100dd600435610423565b34801561015157600080fd5b506100c373ffffffffffffffffffffffffffffffffffffffff60043516610458565b34801561017f57600080fd5b5061011b73ffffffffffffffffffffffffffffffffffffffff600435166104c3565b3480156101ad57600080fd5b506100dd6104d5565b3480156101c257600080fd5b506101e473ffffffffffffffffffffffffffffffffffffffff600435166104f1565b604080519115158252519081900360200190f35b60008083151561020757600080fd5b60008481526001602052604090205473ffffffffffffffffffffffffffffffffffffffff9081169250831682141561023e57600080fd5b60005473ffffffffffffffffffffffffffffffffffffffff1633148061027957503373ffffffffffffffffffffffffffffffffffffffff8316145b905080151561028757600080fd5b610290836104f1565b1561029a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff82811660009081526002602090815260408083208390558783526001909152902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001691851691821790551561032a5773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090208490555b6040805173ffffffffffffffffffffffffffffffffffffffff808516825285166020820152815186927fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c928290030190a25050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60035490565b600380548290811061043157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005473ffffffffffffffffffffffffffffffffffffffff16331461047c57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604081205411905600a165627a7a72305820aac86c9c80b30ddab803841a057ca96a9862b0c0ee28f2c70646ba6b307f21500029";

    public static final String FUNC_UPGRADE = "upgrade";

    public static final String FUNC_CONTRACTADDRESS = "contractAddress";

    public static final String FUNC_CONTRACTSLENGTH = "contractsLength";

    public static final String FUNC_CONTRACTS = "contracts";

    public static final String FUNC_APPOINTMASTERACCOUNT = "appointMasterAccount";

    public static final String FUNC_CONTRACTID = "contractId";

    public static final String FUNC_MASTERACCOUNT = "masterAccount";

    public static final String FUNC_CONTRACTACTIVE = "contractActive";

    public static final Event CONTRACTUPGRADED_EVENT = new Event("ContractUpgraded", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>(true) {}, new TypeReference<Address>() {}, new TypeReference<Address>() {}));
    ;

    protected CryptoFiat(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected CryptoFiat(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<TransactionReceipt> upgrade(BigInteger id, String next) {
        final Function function = new Function(
                FUNC_UPGRADE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(id), 
                new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> contractAddress(BigInteger param0) {
        final Function function = new Function(FUNC_CONTRACTADDRESS, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<BigInteger> contractsLength() {
        final Function function = new Function(FUNC_CONTRACTSLENGTH, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<String> contracts(BigInteger param0) {
        final Function function = new Function(FUNC_CONTRACTS, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<TransactionReceipt> appointMasterAccount(String next) {
        final Function function = new Function(
                FUNC_APPOINTMASTERACCOUNT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(next)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<BigInteger> contractId(String param0) {
        final Function function = new Function(FUNC_CONTRACTID, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<String> masterAccount() {
        final Function function = new Function(FUNC_MASTERACCOUNT, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<Boolean> contractActive(String addr) {
        final Function function = new Function(FUNC_CONTRACTACTIVE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(addr)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public static RemoteCall<CryptoFiat> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(CryptoFiat.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    public static RemoteCall<CryptoFiat> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(CryptoFiat.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    public List<ContractUpgradedEventResponse> getContractUpgradedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(CONTRACTUPGRADED_EVENT, transactionReceipt);
        ArrayList<ContractUpgradedEventResponse> responses = new ArrayList<ContractUpgradedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            ContractUpgradedEventResponse typedResponse = new ContractUpgradedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.id = (BigInteger) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.previous = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.next = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<ContractUpgradedEventResponse> contractUpgradedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, ContractUpgradedEventResponse>() {
            @Override
            public ContractUpgradedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(CONTRACTUPGRADED_EVENT, log);
                ContractUpgradedEventResponse typedResponse = new ContractUpgradedEventResponse();
                typedResponse.log = log;
                typedResponse.id = (BigInteger) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.previous = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.next = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<ContractUpgradedEventResponse> contractUpgradedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(CONTRACTUPGRADED_EVENT));
        return contractUpgradedEventObservable(filter);
    }

    public static CryptoFiat load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new CryptoFiat(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static CryptoFiat load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new CryptoFiat(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static class ContractUpgradedEventResponse {
        public Log log;

        public BigInteger id;

        public String previous;

        public String next;
    }
}
