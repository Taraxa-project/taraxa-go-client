package taraxa_client

import (
	"errors"
	"math/big"

	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Network uint64

const (
	Undefined Network = iota
	Mainnet
	Testnet
	Devnet
)

// TaraxaClient contains variables and methods needed for communication with taraxa dpos smart contract and RPC
type TaraxaClient struct {
	chainID            *big.Int
	DposContractClient *dpos_contract_client.DposContractClient
	EthRpcClient       *ethclient.Client
	// TODO: add taraxaRpcClient
}

func NewTaraxaClient(chainID *big.Int, url string, dposContractAddress common.Address) (*TaraxaClient, error) {
	taraxaClient := new(TaraxaClient)
	taraxaClient.chainID = chainID

	var err error
	taraxaClient.EthRpcClient, err = ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	taraxaClient.DposContractClient, err = dpos_contract_client.NewDposContractClient(taraxaClient.EthRpcClient, dposContractAddress, taraxaClient.chainID)
	if err != nil {
		return nil, err
	}

	return taraxaClient, nil
}

// Creates new taraxa client based on default values set for different taraxa networks
func NewDefaultTaraxaClient(network Network) (*TaraxaClient, error) {
	var networkRpcUrl string
	var chainID *big.Int
	switch network {
	case Mainnet:
		networkRpcUrl = "https://rpc.mainnet.taraxa.io"
		chainID = big.NewInt(841)
		break
	case Testnet:
		networkRpcUrl = "https://rpc.testnet.taraxa.io"
		chainID = big.NewInt(842)
		break
	case Devnet:
		networkRpcUrl = "https://rpc.devnet.taraxa.io"
		chainID = big.NewInt(843)
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	// Precompiled dpos contract has the same address in all taraxa networks
	dposContractAddress := common.HexToAddress("0x00000000000000000000000000000000000000FE")

	return NewTaraxaClient(chainID, networkRpcUrl, dposContractAddress)
}
