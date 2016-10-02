rmdir /s /q .bin
mkdir .bin

solc -o ./.bin --bin --abi --asm --optimize --optimize-runs 1000000 --gas ./CryptoFiat.sol >> ./.bin/_gas.log