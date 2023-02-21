package main

import (
	"log"
	dpos_sc_client "taraxa-dpos-sc-client/dpos_sc_client"
)

func main() {
	dposClient, err := dpos_sc_client.NewDposScClient("https://rpc.testnet.taraxa.io")
	if err != nil {
		log.Fatal(err)
	}

	totalEligibleVotesCount, err := dposClient.GetTotalEligibleVotesCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetTotalEligibleVotesCount: %d\n\n", totalEligibleVotesCount)

	validators, err := dposClient.GetValidators()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetValidators count: %d\n\n", len(validators))

	if len(validators) != 0 {
		validator, err := dposClient.GetValidator(validators[0].Account)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("GetValidator for %s: %s\n\n", validators[0].Account, validator.Description)

		validatorEligibleVotesCount, err := dposClient.GetValidatorEligibleVotesCount(validators[0].Account)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("GetValidatorEligibleVotesCount for %s: %d\n\n", validators[0].Account, validatorEligibleVotesCount)

		ownerValidators, err := dposClient.GetOwnerValidators(validators[0].Info.Owner)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("GetOwnerValidators count for %s: %d\n\n", validators[0].Info.Owner, len(ownerValidators))
	}

	// client, err := ethclient.Dial("https://rpc.devnet.taraxa.io")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(300000) // in units
	// auth.GasPrice = gasPrice

	// address := common.HexToAddress("0x00000000000000000000000000000000000000FE")
	// instance, err := taraxa_dpos_interface.NewTaraxaDposInterface(address, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(instance)

	// opts := bind.CallOpts{}
	// num, err := instance.GetTotalEligibleVotesCount(&opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetTotalEligibleVotesCount: %d", num)
}
