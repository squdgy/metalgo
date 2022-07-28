// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/MetalBlockchain/avalanchego/snow"
	"github.com/MetalBlockchain/avalanchego/snow/uptime"
	"github.com/MetalBlockchain/avalanchego/utils"
	"github.com/MetalBlockchain/avalanchego/utils/timer/mockable"
	"github.com/MetalBlockchain/avalanchego/vms/platformvm/config"
	"github.com/MetalBlockchain/avalanchego/vms/platformvm/fx"
	"github.com/MetalBlockchain/avalanchego/vms/platformvm/reward"
	"github.com/MetalBlockchain/avalanchego/vms/platformvm/state"
	"github.com/MetalBlockchain/avalanchego/vms/platformvm/utxo"
)

type Backend struct {
	Config        *config.Config
	Ctx           *snow.Context
	Clk           *mockable.Clock
	Fx            fx.Fx
	FlowChecker   utxo.Verifier
	Uptimes       uptime.Manager
	Rewards       reward.Calculator
	Bootstrapped  *utils.AtomicBool
	StateVersions state.Versions
}
