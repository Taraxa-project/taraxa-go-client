package dpos_sc_client

import (
	"log"

	dpos_interface "taraxa-dpos-sc-client/dpos_sc_client/dpos_interface"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DposScClient contains variables needed for communication with taraxa dpos smart contract
type DposScClient struct {
	dposInterface *dpos_interface.TaraxaDposInterface
}

func NewDposScClient(rpcNodeURL string) (*DposScClient, error) {
	ethClient, err := ethclient.Dial(rpcNodeURL)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Precompiled dpos contract has the same address in all taraxa networks
	dposScAddress := common.HexToAddress("0x00000000000000000000000000000000000000FE")

	dposScClient := new(DposScClient)
	dposScClient.dposInterface, err = dpos_interface.NewTaraxaDposInterface(dposScAddress, ethClient)
	if err != nil {
		log.Print(err)
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
