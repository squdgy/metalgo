// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/utils/perms"
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
)

var (
	//go:embed genesis_test.json
	customGenesisConfigJSON  []byte
	invalidGenesisConfigJSON = []byte(`{
		"networkID": 9999}}}}
	}`)
)

func TestValidateConfig(t *testing.T) {
	tests := map[string]struct {
		networkID uint32
		config    *Config
		err       string
	}{
		"mainnet": {
			networkID: 1191772,
			config:    &MainnetConfig,
		},
		"fuji": {
			networkID: 1191773,
			config:    &TestnetConfig,
		},
		"local": {
			networkID: 12345,
			config:    &LocalConfig,
		},
		"mainnet (networkID mismatch)": {
			networkID: 2,
			config:    &MainnetConfig,
			err:       "networkID 2 specified but genesis config contains networkID 1",
		},
		"invalid start time": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.StartTime = 999999999999999
				return &thisConfig
			}(),
			err: "start time cannot be in the future",
		},
		"no initial supply": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.Allocations = []Allocation{}
				return &thisConfig
			}(),
			err: "initial supply must be > 0",
		},
		"no initial stakers": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakers = []Staker{}
				return &thisConfig
			}(),
			err: "initial stakers must be > 0",
		},
		"invalid initial stake duration": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakeDuration = 0
				return &thisConfig
			}(),
			err: "initial stake duration must be > 0",
		},
		"invalid stake offset": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakeDurationOffset = 100000000
				return &thisConfig
			}(),
			err: "initial stake duration is 31536000 but need at least 400000000 with offset of 100000000",
		},
		"empty initial staked funds": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakedFunds = []ids.ShortID(nil)
				return &thisConfig
			}(),
			err: "initial staked funds cannot be empty",
		},
		"duplicate initial staked funds": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.InitialStakedFunds = append(thisConfig.InitialStakedFunds, thisConfig.InitialStakedFunds[0])
				return &thisConfig
			}(),
			err: "duplicated in initial staked funds",
		},
		"initial staked funds not in allocations": {
			networkID: 1191773,
			config: func() *Config {
				thisConfig := TestnetConfig
				thisConfig.InitialStakedFunds = append(thisConfig.InitialStakedFunds, LocalConfig.InitialStakedFunds[0])
				return &thisConfig
			}(),
			err: "does not have an allocation to stake",
		},
		"empty C-Chain genesis": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.CChainGenesis = ""
				return &thisConfig
			}(),
			err: "C-Chain genesis cannot be empty",
		},
		"empty message": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.Message = ""
				return &thisConfig
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			err := validateConfig(test.networkID, test.config)
			if len(test.err) > 0 {
				assert.Error(err)
				assert.Contains(err.Error(), test.err)
				return
			}
			assert.NoError(err)
		})
	}
}

func TestGenesisFromFile(t *testing.T) {
	tests := map[string]struct {
		networkID       uint32
		customConfig    []byte
		missingFilepath string
		err             string
		expected        string
	}{
		"mainnet": {
			networkID:    constants.MainnetID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network mainnet (1191772)",
		},
		"testnet": {
			networkID:    constants.TestnetID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network testnet (1191773)",
		},
		"testnet (with custom specified)": {
			networkID:    constants.TestnetID,
			customConfig: localGenesisConfigJSON, // won't load
			err:          "cannot override genesis config for standard network testnet (1191773)",
		},
		"local": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"local (with custom specified)": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"custom": {
			networkID:    9999,
			customConfig: customGenesisConfigJSON,
			expected:     "18dc0a7ab9b257b86ee65870649a0dabbb76c03d0ca9f5c44a43ce0d0a7f20c6",
		},
		"custom (networkID mismatch)": {
			networkID:    9999,
			customConfig: localGenesisConfigJSON,
			err:          "networkID 9999 specified but genesis config contains networkID 12345",
		},
		"custom (invalid format)": {
			networkID:    9999,
			customConfig: invalidGenesisConfigJSON,
			err:          "unable to load provided genesis config",
		},
		"custom (missing filepath)": {
			networkID:       9999,
			missingFilepath: "missing.json",
			err:             "unable to load provided genesis config",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// test loading of genesis from file

			assert := assert.New(t)
			var customFile string
			if len(test.customConfig) > 0 {
				customFile = filepath.Join(t.TempDir(), "config.json")
				assert.NoError(perms.WriteFile(customFile, test.customConfig, perms.ReadWrite))
			}

			if len(test.missingFilepath) > 0 {
				customFile = test.missingFilepath
			}

			genesisBytes, _, err := FromFile(test.networkID, customFile)
			if len(test.err) > 0 {
				assert.Error(err)
				assert.Contains(err.Error(), test.err)
				return
			}
			assert.NoError(err)

			genesisHash := fmt.Sprintf("%x", hashing.ComputeHash256(genesisBytes))
			assert.Equal(test.expected, genesisHash, "genesis hash mismatch")

			_, err = genesis.Parse(genesisBytes)
			assert.NoError(err)
		})
	}
}

func TestGenesisFromFlag(t *testing.T) {
	tests := map[string]struct {
		networkID    uint32
		customConfig []byte
		err          string
		expected     string
	}{
		"mainnet": {
			networkID: constants.MainnetID,
			err:       "cannot override genesis config for standard network mainnet (1191772)",
		},
		"testnet": {
			networkID: constants.TestnetID,
			err:       "cannot override genesis config for standard network testnet (1191773)",
		},
		"local": {
			networkID: constants.LocalID,
			err:       "cannot override genesis config for standard network local (12345)",
		},
		"local (with custom specified)": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"custom": {
			networkID:    9999,
			customConfig: customGenesisConfigJSON,
			expected:     "18dc0a7ab9b257b86ee65870649a0dabbb76c03d0ca9f5c44a43ce0d0a7f20c6",
		},
		"custom (networkID mismatch)": {
			networkID:    9999,
			customConfig: localGenesisConfigJSON,
			err:          "networkID 9999 specified but genesis config contains networkID 12345",
		},
		"custom (invalid format)": {
			networkID:    9999,
			customConfig: invalidGenesisConfigJSON,
			err:          "unable to load genesis content from flag",
		},
		"custom (missing content)": {
			networkID: 9999,
			err:       "unable to load genesis content from flag",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// test loading of genesis content from flag/env-var

			assert := assert.New(t)
			var genBytes []byte
			if len(test.customConfig) == 0 {
				// try loading a default config
				var err error
				switch test.networkID {
				case constants.MainnetID:
					genBytes, err = json.Marshal(&MainnetConfig)
					assert.NoError(err)
				case constants.TestnetID:
					genBytes, err = json.Marshal(&TestnetConfig)
					assert.NoError(err)
				case constants.LocalID:
					genBytes, err = json.Marshal(&LocalConfig)
					assert.NoError(err)
				default:
					genBytes = make([]byte, 0)
				}
			} else {
				genBytes = test.customConfig
			}
			content := base64.StdEncoding.EncodeToString(genBytes)

			genesisBytes, _, err := FromFlag(test.networkID, content)
			if len(test.err) > 0 {
				assert.Error(err)
				assert.Contains(err.Error(), test.err)
				return
			}
			assert.NoError(err)

			genesisHash := fmt.Sprintf("%x", hashing.ComputeHash256(genesisBytes))
			assert.Equal(test.expected, genesisHash, "genesis hash mismatch")

			_, err = genesis.Parse(genesisBytes)
			assert.NoError(err)
		})
	}
}

func TestVMGenesis(t *testing.T) {
	type vmTest struct {
		vmID       ids.ID
		expectedID string
	}
	tests := []struct {
		networkID uint32
		vmTest    []vmTest
	}{
		{
			networkID: constants.MainnetID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "61paDRc4BPvm6KFMMcWr3ZZvfb2FcFnFNMQHrZVNexWBzURA2",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2fJ5SRp51iZfrbzzGsWSBzphY9onEuhae7T5iXmiW32FytdsMU",
				},
			},
		},
		{
			networkID: constants.TestnetID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "2aBeLFAUpXPN8d7Uj52jUD8uqR2to6YvNSmEotpT35WHFC4Ufy",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2o1u8DaxAUKyJC2j6kQE3N3st92wTBM9VsCG7puJaYQWif3Xpf",
				},
			},
		},
		{
			networkID: constants.LocalID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "2EwQEpAoKi3b8GgAhKGprkGiJYFduE9bbbEANMWioSWXsb6ZQz",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU",
				},
			},
		},
	}

	for _, test := range tests {
		for _, vmTest := range test.vmTest {
			name := fmt.Sprintf("%s-%s",
				constants.NetworkIDToNetworkName[test.networkID],
				vmTest.vmID,
			)
			t.Run(name, func(t *testing.T) {
				assert := assert.New(t)

				config := GetConfig(test.networkID)
				genesisBytes, _, err := FromConfig(config)
				assert.NoError(err)

				genesisTx, err := VMGenesis(genesisBytes, vmTest.vmID)
				assert.NoError(err)

				assert.Equal(
					vmTest.expectedID,
					genesisTx.ID().String(),
					"%s genesisID with networkID %d mismatch",
					vmTest.vmID,
					test.networkID,
				)
			})
		}
	}
}

func TestAVAXAssetID(t *testing.T) {
	tests := []struct {
		networkID  uint32
		expectedID string
	}{
		{
			networkID:  constants.MainnetID,
			expectedID: "2wDksesCMAUDEuqeKX2g9Zkbm1dk48NBKiye4E37be1F7xqbMv",
		},
		{
			networkID:  constants.TestnetID,
			expectedID: "22isChTFe6CZnCCj5Kbi6S5efvooPAtUtkzZScGXsvoE5jczDZ",
		},
		{
			networkID:  constants.LocalID,
			expectedID: "ui6eSQwLELPt3zZsreVkdSFhimBdAyXgADd7CGfQ58NeDNzCS",
		},
	}

	for _, test := range tests {
		t.Run(constants.NetworkIDToNetworkName[test.networkID], func(t *testing.T) {
			assert := assert.New(t)

			config := GetConfig(test.networkID)
			_, avaxAssetID, err := FromConfig(config)
			assert.NoError(err)

			assert.Equal(
				test.expectedID,
				avaxAssetID.String(),
				"AVAX assetID with networkID %d mismatch",
				test.networkID,
			)
		})
	}
}
