// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/MetalBlockchain/metalgo/codec"
	"github.com/MetalBlockchain/metalgo/codec/linearcodec"
	"github.com/MetalBlockchain/metalgo/utils/wrappers"
	"github.com/MetalBlockchain/metalgo/vms/platformvm/stakeable"
	"github.com/MetalBlockchain/metalgo/vms/secp256k1fx"
)

// Version is the current default codec version
const Version = 0

// Codec does serialization and deserialization
var Codec codec.Manager

func init() {
	c := linearcodec.NewDefault()
	Codec = codec.NewDefaultManager()

	// Order in which type are registered affect the byte representation
	// generated by marshalling ops. To maintain codec type ordering,
	// we skip positions for the blocks.
	c.SkipRegistrations(5)

	errs := wrappers.Errs{}
	errs.Add(
		RegisterUnsignedTxsTypes(c),
		Codec.RegisterCodec(Version, c),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}

// RegisterUnsignedTxsTypes allows registering relevant type of unsigned package
// in the right sequence. Following repackaging of platformvm package, a few
// subpackage-level codecs were introduced, each handling serialization of specific types.
// RegisterUnsignedTxsTypes is made exportable so to guarantee that other codecs
// are coherent with components one.
func RegisterUnsignedTxsTypes(targetCodec codec.Registry) error {
	errs := wrappers.Errs{}
	errs.Add(
		// The Fx is registered here because this is the same place it is
		// registered in the AVM. This ensures that the typeIDs match up for
		// utxos in shared memory.
		targetCodec.RegisterType(&secp256k1fx.TransferInput{}),
		targetCodec.RegisterType(&secp256k1fx.MintOutput{}),
		targetCodec.RegisterType(&secp256k1fx.TransferOutput{}),
		targetCodec.RegisterType(&secp256k1fx.MintOperation{}),
		targetCodec.RegisterType(&secp256k1fx.Credential{}),
		targetCodec.RegisterType(&secp256k1fx.Input{}),
		targetCodec.RegisterType(&secp256k1fx.OutputOwners{}),

		targetCodec.RegisterType(&AddValidatorTx{}),
		targetCodec.RegisterType(&AddSubnetValidatorTx{}),
		targetCodec.RegisterType(&AddDelegatorTx{}),
		targetCodec.RegisterType(&CreateChainTx{}),
		targetCodec.RegisterType(&CreateSubnetTx{}),
		targetCodec.RegisterType(&ImportTx{}),
		targetCodec.RegisterType(&ExportTx{}),
		targetCodec.RegisterType(&AdvanceTimeTx{}),
		targetCodec.RegisterType(&RewardValidatorTx{}),

		targetCodec.RegisterType(&stakeable.LockIn{}),
		targetCodec.RegisterType(&stakeable.LockOut{}),

		// Blueberry additions:
		targetCodec.RegisterType(&RemoveSubnetValidatorTx{}),
		targetCodec.RegisterType(&TransformSubnetTx{}),
		targetCodec.RegisterType(&AddPermissionlessValidatorTx{}),
		targetCodec.RegisterType(&AddPermissionlessDelegatorTx{}),
	)
	return errs.Err
}
