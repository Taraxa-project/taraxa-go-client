# taraxa-dpos-sc-client
Taraxa dpos interface contract client 


#### Prerequisites
##### solc
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

##### abigen
```
go install github.com/ethereum/go-ethereum@latest
cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v<X.Y.Z>/
make
make devtools
```

#### Create SC ABI
run
```
solc --abi --overwrite --optimize taraxa-evm/taraxa/state/dpos/solidity/dpos_contract_interface.sol --output-dir client/abi
```

#### Create SC go class
run
```
abigen --abi=client/abi/DposInterface.abi --pkg=taraxaDposClient --out=client/taraxa_dpos_client.go
```

### !!! To work with latest dpos contract interface
Update taraxa-evm submodule and re-generate abi & client go implementation by running:
```
git submodule update --init --recursive
solc --abi --overwrite --optimize taraxa-evm/taraxa/state/dpos/solidity/dpos_contract_interface.sol --output-dir .
abigen --abi=DposInterface.abi --pkg=taraxaDposClient --out=dpos_contract_interface.go
```


