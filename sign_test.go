package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"testing"
)

func TestSignTimeStamp(t *testing.T) {
	client, err := ethclient.Dial("https://stardust.metis.io/?owner=588")
	if err != nil {
		t.Fatal(err)
	}

	caller, err := NewSigtest(common.HexToAddress("0x2dC4C689e6cb27be674A9fe279091Ada1634D0B2"), client)
	if err != nil {
		t.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("db96ee03632eb01fed47f061d2ba383ad8231844d51fc94d9a938d2476e7fd53")
	if err != nil {
		t.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(588))

	timestamp, sig, err := SignTimeStamp()
	if err != nil {
		t.Fatal(err)
	}

	timestampInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := caller.Verify(auth, big.NewInt(timestampInt), sig)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tx.Hash())
}
