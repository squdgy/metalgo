// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"log"
	"time"

	"github.com/MetalBlockchain/metalgo/api/info"
	"github.com/MetalBlockchain/metalgo/genesis"
	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/metalgo/utils/units"
	"github.com/MetalBlockchain/metalgo/vms/platformvm/validator"
	"github.com/MetalBlockchain/metalgo/vms/secp256k1fx"
	"github.com/MetalBlockchain/metalgo/wallet/subnet/primary"
)

func main() {
	key := genesis.EWOQKey
	uri := primary.LocalAPIURI
	kc := secp256k1fx.NewKeychain(key)
	subnetIDStr := "29uVeLPJB1eQJkzRemU8g8wZDw5uJRqpab5U2mX9euieVwiEbL"
	startTime := time.Now().Add(time.Minute)
	duration := 2 * 7 * 24 * time.Hour // 2 weeks
	weight := units.Schmeckle

	subnetID, err := ids.FromString(subnetIDStr)
	if err != nil {
		log.Fatalf("failed to parse subnet ID: %s\n", err)
	}

	ctx := context.Background()
	infoClient := info.NewClient(uri)

	nodeInfoStartTime := time.Now()
	nodeID, _, err := infoClient.GetNodeID(ctx)
	if err != nil {
		log.Fatalf("failed to fetch node IDs: %s\n", err)
	}
	log.Printf("fetched node ID %s in %s\n", nodeID, time.Since(nodeInfoStartTime))

	// NewWalletFromURI fetches the available UTXOs owned by [kc] on the network
	// that [uri] is hosting.
	walletSyncStartTime := time.Now()
	wallet, err := primary.NewWalletWithTxs(ctx, uri, kc, subnetID)
	if err != nil {
		log.Fatalf("failed to initialize wallet: %s\n", err)
	}
	log.Printf("synced wallet in %s\n", time.Since(walletSyncStartTime))

	// Get the P-chain wallet
	pWallet := wallet.P()

	addValidatorStartTime := time.Now()
	addValidatorTxID, err := pWallet.IssueAddSubnetValidatorTx(&validator.SubnetValidator{
		Validator: validator.Validator{
			NodeID: nodeID,
			Start:  uint64(startTime.Unix()),
			End:    uint64(startTime.Add(duration).Unix()),
			Wght:   weight,
		},
		Subnet: subnetID,
	})
	if err != nil {
		log.Fatalf("failed to issue add subnet validator transaction: %s\n", err)
	}
	log.Printf("added new subnet validator %s to %s with %s in %s\n", nodeID, subnetID, addValidatorTxID, time.Since(addValidatorStartTime))
}
