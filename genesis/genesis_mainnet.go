// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1191772,
		"allocations": [
			{
				"ethAddr": "0x295fd034acc976e746edd3174800c22bc61d7c76",
				"avaxAddr": "X-metal1gz2y6c0822jq8glt2ezfyvp6zlgwsqfp5kmyhq",
				"initialAmount": 240000000000,
				"unlockSchedule": [
					{
						"amount": 360000000000
					}
				]
			}
		],
		"startTime": 1656606845,
		"initialStakeDuration": 7776000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"X-metal1gz2y6c0822jq8glt2ezfyvp6zlgwsqfp5kmyhq"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-EGBCfWkdCePkCbvAfSBDwWabP5qfa8k83",
				"rewardAddress": "X-metal1gz2y6c0822jq8glt2ezfyvp6zlgwsqfp5kmyhq",
				"delegationFee": 1000000
			},
			{
				"nodeID": "NodeID-2tXdxQWjXaRu4A6eQWfKiD4zmiJdhxkTL",
				"rewardAddress": "X-metal1gz2y6c0822jq8glt2ezfyvp6zlgwsqfp5kmyhq",
				"delegationFee": 500000
			},
			{
				"nodeID": "NodeID-KEb5v3V8FJfbyHtYm1b9CL3fJd4M8DJP7",
				"rewardAddress": "X-metal1gz2y6c0822jq8glt2ezfyvp6zlgwsqfp5kmyhq",
				"delegationFee": 500000
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":43112,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"295fD034aCC976e746EDD3174800c22bC61d7c76\":{\"balance\":\"0x295BE96E64066972000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "Hello world from the Metal team!"
	}`

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAvax,
			CreateAssetTxFee:      10 * units.MilliAvax,
			CreateSubnetTxFee:     1 * units.Avax,
			CreateBlockchainTxFee: 1 * units.Avax,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 2 * units.KiloAvax,
			MaxValidatorStake: 3 * units.MegaAvax,
			MinDelegatorStake: 25 * units.Avax,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  2 * 7 * 24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          720 * units.MegaAvax,
			},
		},
	}
)
