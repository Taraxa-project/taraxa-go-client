package main

import (
	"log"
	"math/big"
	dpos_sc_client "taraxa-dpos-sc-client/dpos_sc_client"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	dposClient, err := dpos_sc_client.NewDposScClient(dpos_sc_client.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	totalEligibleVotesCount, err := dposClient.GetTotalEligibleVotesCount()
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("GetTotalEligibleVotesCount: %d\n\n", totalEligibleVotesCount)
	}

	validators, err := dposClient.GetValidators()
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("GetValidators count: %d\n\n", len(validators))
	}

	if len(validators) != 0 {
		validator, err := dposClient.GetValidator(validators[0].Account)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetValidator for %s: %s\n\n", validators[0].Account, validator.Description)
		}

		isValidatorEligible, err := dposClient.IsValidatorEligible(validators[0].Account)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("IsValidatorEligible for %s: %t\n\n", validators[0].Account, isValidatorEligible)
		}

		validatorEligibleVotesCount, err := dposClient.GetValidatorEligibleVotesCount(validators[0].Account)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetValidatorEligibleVotesCount for %s: %d\n\n", validators[0].Account, validatorEligibleVotesCount)
		}

		ownerValidators, err := dposClient.GetOwnerValidators(validators[0].Info.Owner)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetOwnerValidators count for %s: %d\n\n", validators[0].Info.Owner, len(ownerValidators))
		}

		defaultDevnetDelegator := common.HexToAddress("0x7e4aa664f71de4e9d0b4a6473d796372639bdcde")
		devnet_delegations, err := dposClient.GetDelegations(defaultDevnetDelegator)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetDelegations count for %s: %d\n\n", defaultDevnetDelegator, len(devnet_delegations))
		}

		defaultTestnetDelegator := common.HexToAddress("0x76870407332398322576505f3c5423d0a71af296")
		testnet_delegations, err := dposClient.GetDelegations(defaultTestnetDelegator)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetDelegations count for %s: %d\n\n", defaultTestnetDelegator, len(testnet_delegations))
		}

		devnet_undelegations, err := dposClient.GetUndelegations(defaultDevnetDelegator)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetUndelegations count for %s: %d\n\n", defaultDevnetDelegator, len(devnet_undelegations))
		}

		testnet_undelegations, err := dposClient.GetUndelegations(defaultTestnetDelegator)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("GetUndelegations count for %s: %d\n\n", defaultTestnetDelegator, len(testnet_undelegations))
		}

		// TODO: add here valid private key
		privateKey := "0000000000000000000000000000000000000000000000000000000000000000"

		transactor, err := dposClient.NewTransactor(privateKey)
		if err != nil {
			log.Fatal(err)
		}

		thousandTara, ok := big.NewInt(0).SetString("100000000000000000000", 10)
		if !ok {
			log.Fatal("Unable to create big.Int for delegeation")
		}

		delegate_tx, err := dposClient.Delegate(transactor, thousandTara, validators[0].Account)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Delegate %d to %s, tx hash: %s, tx nonce: %d\n\n", delegate_tx.Value(), validators[0].Account, delegate_tx.Hash(), delegate_tx.Nonce())
		}

		undelegate_tx, err := dposClient.Undelegate(transactor, thousandTara, validators[0].Account)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Undelegate %d to %s, tx hash: %s, tx nonce: %d\n\n", thousandTara, validators[0].Account, undelegate_tx.Hash(), undelegate_tx.Nonce())
		}

		confirmUndelegateTx, err := dposClient.ConfirmUndelegate(transactor, validators[0].Account)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("ConfirmUndelegate to %s, tx hash: %s, tx nonce: %d\n\n", validators[0].Account, confirmUndelegateTx.Hash(), confirmUndelegateTx.Nonce())
		}
	}
}
