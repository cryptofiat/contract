package eu.cryptoeuro.contract;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Bytes32;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 3.5.0.
 */
public class Data extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b506040516020806103d8833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610386806100526000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663295f36d7811461005b578063461b09c014610088578063516c4b84146100a8575b600080fd5b34801561006757600080fd5b506100766004356024356100e6565b60408051918252519081900360200190f35b34801561009457600080fd5b506100a66004356024356044356101bc565b005b3480156100b457600080fd5b506100bd61033e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60408051602080820185905281830184905282518083038401815260609092019283905281516000936001938593909282918401908083835b6020831061015c57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161011f565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01801990921691161790526040805192909401829003909120865285019590955292909201600020549695505050505050565b60008054604080517fe814861e000000000000000000000000000000000000000000000000000000008152336004820152905173ffffffffffffffffffffffffffffffffffffffff9092169263e814861e926024808401936020939083900390910190829087803b15801561023057600080fd5b505af1158015610244573d6000803e3d6000fd5b505050506040513d602081101561025a57600080fd5b5051151561026757600080fd5b60408051602080820186905281830185905282518083038401815260609092019283905281518493600193600093909282918401908083835b602083106102dd57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016102a0565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0180199092169116179052604080519290940182900390912086528501959095529290920160002093909355505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820c85e3b9a9ac41897e74ae4932b2edfb8f97cb13764ba3b49075d56a9acc242930029";

    public static final String FUNC_GET = "get";

    public static final String FUNC_SET = "set";

    public static final String FUNC_CRYPTOFIAT = "cryptoFiat";

    protected Data(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Data(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<byte[]> get(BigInteger bucket, byte[] key) {
        final Function function = new Function(FUNC_GET, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(bucket), 
                new org.web3j.abi.datatypes.generated.Bytes32(key)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bytes32>() {}));
        return executeRemoteCallSingleValueReturn(function, byte[].class);
    }

    public RemoteCall<TransactionReceipt> set(BigInteger bucket, byte[] key, byte[] value) {
        final Function function = new Function(
                FUNC_SET, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(bucket), 
                new org.web3j.abi.datatypes.generated.Bytes32(key), 
                new org.web3j.abi.datatypes.generated.Bytes32(value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> cryptoFiat() {
        final Function function = new Function(FUNC_CRYPTOFIAT, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public static RemoteCall<Data> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat)));
        return deployRemoteCall(Data.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static RemoteCall<Data> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _cryptoFiat) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_cryptoFiat)));
        return deployRemoteCall(Data.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static Data load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new Data(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static Data load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Data(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }
}
