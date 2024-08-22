# taraxa-go-client
Taraxa rpc and dpos contract go client

# Usage
```
package main

import (
	"log"
	"math/big"

	"github.com/Taraxa-project/taraxa-go-client/taraxa_client"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
  taraxaClient, err := taraxa_client.NewTaraxaClient(taraxa_client.Testnet)
  if err != nil {
    log.Fatal(err)
  }

  // Get all valdiators from dpos contract
  validators, err := taraxaClient.DposContractClient.GetValidators()
  if err != nil {
    log.Fatal(err)
  }
  // Process validators


  // Delegate to validator
  // TODO: add here valid private key
  privateKey := "0000000000000000000000000000000000000000000000000000000000000000"

  transactor, err := taraxaClient.DposContractClient.NewTransactor(privateKey)
  if err != nil {
    log.Fatal(err)
  }

  thousandTara, ok := big.NewInt(0).SetString("1000000000000000000000", 10)
  delegate_tx, err := taraxaClient.DposContractClient.Delegate(transactor, thousandTara, validators[0].Account)
  if err != nil {
    log.Fatal(err)
  }
}
```

See demo.go for more examples


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
solc --abi --overwrite --optimize taraxa-evm/taraxa/state/contracts/dpos/solidity/dpos_contract_interface.sol --output-dir taraxa_client/dpos_contract_client/dpos_interface/
```

#### Create SC go class
run
```
abigen --abi=taraxa_client/dpos_contract_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=taraxa_client/dpos_contract_client/dpos_interface/dpos_interface.go
```

### !!! To work with latest dpos contract interface
Update taraxa-evm submodule and re-generate abi & client go implementation by running:
```
git submodule update --init --recursive

solc --abi --overwrite --optimize taraxa-evm/taraxa/state/contracts/dpos/solidity/dpos_contract_interface.sol --output-dir taraxa_client/dpos_contract_client/dpos_interface/ 

abigen --abi=taraxa_client/dpos_contract_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=taraxa_client/dpos_contract_client/dpos_interface/dpos_interface.go
```


