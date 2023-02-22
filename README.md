# taraxa-go-client
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
solc --abi --overwrite --optimize taraxa-evm/taraxa/state/dpos/solidity/dpos_contract_interface.sol --output-dir dpos_sc_client/dpos_interface/
```

#### Create SC go class
run
```
abigen --abi=dpos_sc_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=dpos_sc_client/dpos_interface/dpos_interface.go
```

### !!! To work with latest dpos contract interface
Update taraxa-evm submodule and re-generate abi & client go implementation by running:
```
git submodule update --init --recursive
solc --abi --overwrite --optimize taraxa-evm/taraxa/state/dpos/solidity/dpos_contract_interface.sol --output-dir dpos_sc_client/dpos_interface/
abigen --abi=dpos_sc_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=dpos_sc_client/dpos_interface/dpos_interface.go
```


