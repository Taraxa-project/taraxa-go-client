package dpos_sc_client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	dpos_interface "taraxa-dpos-sc-client/dpos_sc_client/dpos_interface"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Network uint64

const (
	Undefined Network = iota
	Mainnet
	Testnet
	Devnet
)

// DposScClient contains variables needed for communication with taraxa dpos smart contract
type DposScClient struct {
	dposInterface *dpos_interface.TaraxaDposInterface
	ethClient     *ethclient.Client
	chainID       *big.Int
}

type Transactor struct {
	transactOpts *bind.TransactOpts
	address      common.Address
}

func NewDposScClient(network Network) (*DposScClient, error) {
	dposScClient := new(DposScClient)

	var networkRpcUrl string
	switch network {
	case Mainnet:
		networkRpcUrl = "https://rpc.mainnet.taraxa.io"
		dposScClient.chainID = big.NewInt(841)
		break
	case Testnet:
		networkRpcUrl = "https://rpc.testnet.taraxa.io"
		dposScClient.chainID = big.NewInt(842)
		break
	case Devnet:
		networkRpcUrl = "https://rpc.devnet.taraxa.io"
		dposScClient.chainID = big.NewInt(843)
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	var err error
	dposScClient.ethClient, err = ethclient.Dial(networkRpcUrl)
	if err != nil {
		return nil, err
	}

	// Precompiled dpos contract has the same address in all taraxa networks
	dposScAddress := common.HexToAddress("0x00000000000000000000000000000000000000FE")

	dposScClient.dposInterface, err = dpos_interface.NewTaraxaDposInterface(dposScAddress, dposScClient.ethClient)
	if err != nil {
		return nil, err
	}

	return dposScClient, nil
}

func (dposScClient *DposScClient) GetTotalEligibleVotesCount() (uint64, error) {
	return dposScClient.dposInterface.GetTotalEligibleVotesCount(&bind.CallOpts{})
}

func (dposScClient *DposScClient) GetValidator(validator common.Address) (dpos_interface.DposInterfaceValidatorBasicInfo, error) {
	return dposScClient.dposInterface.GetValidator(&bind.CallOpts{}, validator)
}

func (dposScClient *DposScClient) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return dposScClient.dposInterface.GetValidatorEligibleVotesCount(&bind.CallOpts{}, validator)
}

func (dposScClient *DposScClient) IsValidatorEligible(validator common.Address) (bool, error) {
	return dposScClient.dposInterface.IsValidatorEligible(&bind.CallOpts{}, validator)
}

func (dposScClient *DposScClient) GetValidators() ([]dpos_interface.DposInterfaceValidatorData, error) {
	var validators []dpos_interface.DposInterfaceValidatorData

	for batch := uint32(0); ; batch++ {
		validatorsBatch, err := dposScClient.dposInterface.GetValidators(&bind.CallOpts{}, batch)
		if err != nil {
			return nil, err
		}

		if len(validatorsBatch.Validators) != 0 {
			validators = append(validators, validatorsBatch.Validators...)
		}

		if validatorsBatch.End == true {
			break
		}
	}

	return validators, nil
}

func (dposScClient *DposScClient) GetOwnerValidators(owner common.Address) ([]dpos_interface.DposInterfaceValidatorData, error) {
	var validators []dpos_interface.DposInterfaceValidatorData

	for batch := uint32(0); ; batch++ {
		validatorsBatch, err := dposScClient.dposInterface.GetValidatorsFor(&bind.CallOpts{}, owner, batch)
		if err != nil {
			return nil, err
		}

		if len(validatorsBatch.Validators) != 0 {
			validators = append(validators, validatorsBatch.Validators...)
		}

		if validatorsBatch.End == true {
			break
		}
	}

	return validators, nil
}

func (dposScClient *DposScClient) GetDelegations(delegator common.Address) ([]dpos_interface.DposInterfaceDelegationData, error) {
	var delegations []dpos_interface.DposInterfaceDelegationData

	for batch := uint32(0); ; batch++ {
		delegationsBatch, err := dposScClient.dposInterface.GetDelegations(&bind.CallOpts{}, delegator, batch)
		if err != nil {
			return nil, err
		}

		if len(delegationsBatch.Delegations) != 0 {
			delegations = append(delegations, delegationsBatch.Delegations...)
		}

		if delegationsBatch.End == true {
			break
		}
	}

	return delegations, nil
}

func (dposScClient *DposScClient) GetUndelegations(delegator common.Address) ([]dpos_interface.DposInterfaceUndelegationData, error) {
	var undelegations []dpos_interface.DposInterfaceUndelegationData

	for batch := uint32(0); ; batch++ {
		undelegationsBatch, err := dposScClient.dposInterface.GetUndelegations(&bind.CallOpts{}, delegator, batch)
		if err != nil {
			return nil, err
		}

		if len(undelegationsBatch.Undelegations) != 0 {
			undelegations = append(undelegations, undelegationsBatch.Undelegations...)
		}

		if undelegationsBatch.End == true {
			break
		}
	}

	return undelegations, nil
}

func (dposScClient *DposScClient) NewTransactor(privateKeyStr string) (*Transactor, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, dposScClient.chainID)
	if err != nil {
		return nil, err
	}

	transactor := new(Transactor)
	transactor.address = crypto.PubkeyToAddress(*publicKeyECDSA)
	transactor.transactOpts = new(bind.TransactOpts)
	*transactor.transactOpts = *transactOpts

	return transactor, nil
}

func (dposScClient *DposScClient) createNewTransactOpts(transactor *Transactor) (*bind.TransactOpts, error) {
	nonce, err := dposScClient.ethClient.PendingNonceAt(context.Background(), transactor.address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := dposScClient.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	transactOpts := new(bind.TransactOpts)
	*transactOpts = *transactor.transactOpts

	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasLimit = uint64(300000) // in units
	transactOpts.GasPrice = gasPrice

	return transactOpts, nil
}

func (dposScClient *DposScClient) Delegate(transactor *Transactor, amount *big.Int, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	transactOpts.Value = amount // in wei
	return dposScClient.dposInterface.Delegate(transactOpts, validator)
}

func (dposScClient *DposScClient) Undelegate(transactor *Transactor, amount *big.Int, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.Undelegate(transactOpts, validator, amount)
}

func (dposScClient *DposScClient) ConfirmUndelegate(transactor *Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.ConfirmUndelegate(transactOpts, validator)
}

func (dposScClient *DposScClient) CancelUndelegate(transactor *Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.CancelUndelegate(transactOpts, validator)
}

func (dposScClient *DposScClient) RedelegateUndelegate(transactor *Transactor, amount *big.Int, validatorFrom common.Address, validatorTo common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.ReDelegate(transactOpts, validatorFrom, validatorTo, amount)
}

func (dposScClient *DposScClient) ClaimRewards(transactor *Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.ClaimRewards(transactOpts, validator)
}

func (dposScClient *DposScClient) ClaimCommissionRewards(transactor *Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.ClaimCommissionRewards(transactOpts, validator)
}

func (dposScClient *DposScClient) RegisterValidator(transactor *Transactor, validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.RegisterValidator(transactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

func (dposScClient *DposScClient) SetValidatorInfo(transactor *Transactor, validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.SetValidatorInfo(transactOpts, validator, description, endpoint)
}

func (dposScClient *DposScClient) SetCommission(transactor *Transactor, validator common.Address, commission uint16) (*types.Transaction, error) {
	transactOpts, err := dposScClient.createNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return dposScClient.dposInterface.SetCommission(transactOpts, validator, commission)
}
