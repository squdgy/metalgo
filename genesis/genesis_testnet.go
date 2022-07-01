// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
)

var (
	testnetGenesisConfigJSON = `{
		"networkID": 1191773,
		"allocations": [
			{
				"ethAddr": "0x15884b0f70b7c0db084e5aa738605f7a681a5d6e",
				"avaxAddr": "X-testnet1y5669vthzv2f9hy3uqryz8cf70kdqevssn6lwe",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 1000000000000000
					}
				]
			},
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-testnet1c0d772twvlhnhp637ddktxm08w3nmd2cyvgawu",
				"initialAmount": 166416666000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-testnet1h30nf9jj2g0dm5nh928my04kze2wlw4tnka6v5",
				"initialAmount": 166416666000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-testnet133qdfexfyhckkky2g59n4anekg2vwp287dvws9",
				"initialAmount": 166416666000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-testnet1zj0476f9gdwtevwlqv696fz4jfl04c7k89p9zp",
				"initialAmount": 166416666000000000,
				"unlockSchedule": []
			}
		],
		"startTime": 1656658190,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 54000,
		"initialStakedFunds": [
			"X-testnet1y5669vthzv2f9hy3uqryz8cf70kdqevssn6lwe"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-FURawqrxVH1HuEGXc2oke1gKZw2nJCheV",
				"rewardAddress": "X-testnet1y5669vthzv2f9hy3uqryz8cf70kdqevssn6lwe",
				"delegationFee": 1000000
			},
			{
				"nodeID": "NodeID-CrJaaZKxCaavaaq84phT5TxRzs2hjW1Kz",
				"rewardAddress": "X-testnet1y5669vthzv2f9hy3uqryz8cf70kdqevssn6lwe",
				"delegationFee": 500000
			},
			{
				"nodeID": "NodeID-7Uyxha3stBRRktZT2uWb4oEyaepttmj2t",
				"rewardAddress": "X-testnet1y5669vthzv2f9hy3uqryz8cf70kdqevssn6lwe",
				"delegationFee": 250000
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":43119,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0,\"apricotPhase3BlockTimestamp\":0,\"apricotPhase4BlockTimestamp\":0,\"apricotPhase5BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"15884B0F70b7C0db084e5aa738605f7a681A5D6E\":{\"balance\":\"0xDE0B6B3A7640000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "hi dad"
	}`

	// TestnetParams are the params used for the testnet
	TestnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAvax,
			CreateAssetTxFee:      10 * units.MilliAvax,
			CreateSubnetTxFee:     100 * units.MilliAvax,
			CreateBlockchainTxFee: 100 * units.MilliAvax,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 1 * units.Avax,
			MaxValidatorStake: 3 * units.MegaAvax,
			MinDelegatorStake: 1 * units.Avax,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  24 * time.Hour,
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
