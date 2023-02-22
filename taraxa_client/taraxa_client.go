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

func NewTaraxaClient(network Network) (*TaraxaClient, error) {
	taraxaClient := new(TaraxaClient)

	var networkRpcUrl string
	switch network {
	case Mainnet:
		networkRpcUrl = "https://rpc.mainnet.taraxa.io"
		taraxaClient.chainID = big.NewInt(841)
		break
	case Testnet:
		networkRpcUrl = "https://rpc.testnet.taraxa.io"
		taraxaClient.chainID = big.NewInt(842)
		break
	case Devnet:
		networkRpcUrl = "https://rpc.devnet.taraxa.io"
		taraxaClient.chainID = big.NewInt(843)
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	var err error
	taraxaClient.EthRpcClient, err = ethclient.Dial(networkRpcUrl)
	if err != nil {
		return nil, err
	}

	// Precompiled dpos contract has the same address in all taraxa networks
	dposContractAddress := common.HexToAddress("0x00000000000000000000000000000000000000FE")

	taraxaClient.DposContractClient, err = dpos_contract_client.NewDposContractClient(taraxaClient.EthRpcClient, dposContractAddress, taraxaClient.chainID)
	if err != nil {
		return nil, err
	}

	return taraxaClient, nil
}
