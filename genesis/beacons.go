// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/MetalBlockchain/avalanchego/utils/constants"
	"github.com/MetalBlockchain/avalanchego/utils/sampler"
)

// getIPs returns the beacon IPs for each network
func getIPs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"54.94.43.49:9651",
			"52.79.47.77:9651",
			"18.229.206.191:9651",
			"3.34.221.73:9651",
			"13.244.155.170:9651",
			"13.244.47.224:9651",
			"122.248.200.212:9651",
			"52.30.9.211:9651",
			"122.248.199.127:9651",
			"18.202.190.40:9651",
			"15.206.182.45:9651",
			"15.207.11.193:9651",
			"44.226.118.72:9651",
			"54.185.87.50:9651",
			"18.158.15.12:9651",
			"3.21.38.33:9651",
			"54.93.182.129:9651",
			"3.128.138.36:9651",
			"3.104.107.241:9651",
			"3.106.25.139:9651",
			"18.162.129.129:9651",
			"18.162.161.230:9651",
			"52.47.181.114:9651",
			"15.188.9.42:9651",
		}
	case constants.TahoeID:
		return []string{
			"18.158.206.92:9651",
			"18.157.229.117:9651",
			"3.69.225.81:9651",
			"3.73.114.4:9651",
			"3.74.175.192:9651",
			"34.196.246.54:9651",
			"54.174.105.14:9651",
			"34.206.150.45:9651",
			"34.194.108.44:9651",
			"44.195.208.173:9651",
			"34.212.255.6:9651",
			"54.71.4.64:9651",
			"44.226.132.246:9651",
			"52.43.222.212:9651",
			"54.189.216.170:9651",
			"18.178.22.72:9651",
			"35.75.87.63:9651",
			"54.65.199.213:9651",
			"18.177.175.140:9651",
			"3.113.255.59:9651",
		}
	default:
		return nil
	}
}

// getNodeIDs returns the beacon node IDs for each network
func getNodeIDs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"NodeID-A6onFGyJjA37EZ7kYHANMR1PFRT8NmXrF",
			"NodeID-6SwnPJLH8cWfrJ162JjZekbmzaFpjPcf",
			"NodeID-GSgaA47umS1px2ohVjodW9621Ks63xDxD",
			"NodeID-BQEo5Fy1FRKLbX51ejqDd14cuSXJKArH2",
			"NodeID-Drv1Qh7iJvW3zGBBeRnYfCzk56VCRM2GQ",
			"NodeID-DAtCoXfLT6Y83dgJ7FmQg8eR53hz37J79",
			"NodeID-FGRoKnyYKFWYFMb6Xbocf4hKuyCBENgWM",
			"NodeID-Dw7tuwxpAmcpvVGp9JzaHAR3REPoJ8f2R",
			"NodeID-4kCLS16Wy73nt1Zm54jFZsL7Msrv3UCeJ",
			"NodeID-9T7NXBFpp8LWCyc58YdKNoowDipdVKAWz",
			"NodeID-6ghBh6yof5ouMCya2n9fHzhpWouiZFVVj",
			"NodeID-HiFv1DpKXkAAfJ1NHWVqQoojjznibZXHP",
			"NodeID-Fv3t2shrpkmvLnvNzcv1rqRKbDAYFnUor",
			"NodeID-AaxT2P4uuPAHb7vAD8mNvjQ3jgyaV7tu9",
			"NodeID-kZNuQMHhydefgnwjYX1fhHMpRNAs9my1",
			"NodeID-A7GwTSd47AcDVqpTVj7YtxtjHREM33EJw",
			"NodeID-Hr78Fy8uDYiRYocRYHXp4eLCYeb8x5UuM",
			"NodeID-9CkG9MBNavnw7EVSRsuFr7ws9gascDQy3",
			"NodeID-A8jypu63CWp76STwKdqP6e9hjL675kdiG",
			"NodeID-HsBEx3L71EHWSXaE6gvk2VsNntFEZsxqc",
			"NodeID-Nr584bLpGgbCUbZFSBaBz3Xum5wpca9Ym",
			"NodeID-QKGoUvqcgormCoMj6yPw9isY7DX9H4mdd",
			"NodeID-HCw7S2TVbFPDWNBo1GnFWqJ47f9rDJtt1",
			"NodeID-FYv1Lb29SqMpywYXH7yNkcFAzRF2jvm3K",
		}
	case constants.TahoeID:
		return []string{
			"NodeID-EiRsmffiu3t1MFnhjrrFJEJGqfdtnkEvc",
			"NodeID-6QRpj3qBVPzNwdD2YewfB66Aa9Z9pv243",
			"NodeID-EW9JuwPmgSdt5SXkePz1JnfW3x9zzsqkd",
			"NodeID-AGXv6QpCY4m8Qmug7A9pDYrSSxqdbJAZS",
			"NodeID-3BPMUK91rSbiG1mdBmpissDqTmVFAfc8J",
			"NodeID-B73oENBBKtquFNWkjPfjN9F2pHZp5r16g",
			"NodeID-GJhvteVE4Rn7unxmzZrTwpUmuHvT4ioJU",
			"NodeID-CfsdUtCrAB6grr8kBSLiHSdrKe993Gwvg",
			"NodeID-Ew9uKh7xnyEiZHU1xwi5MMVjHsxkmPScN",
			"NodeID-3uC5oYpq5hqn8Ky46qT13pEYK6avjZY6T",
			"NodeID-GmRnQWwWNjeZ2g7hbaHabHMgpcUABZzXV",
			"NodeID-HSxLZv4zNkRF9vH2xYQW9zkLwtyQdjnne",
			"NodeID-H6xhpvuVgKEXjTYbfAH629SnoZFLtsTMW",
			"NodeID-2eypiRnhHyoZe94PBxbwu1TBDNAP9Te7e",
			"NodeID-AHxkNdDLHMY7x3xFoGXetZgQVfQK9mHpz",
			"NodeID-BJYAmxn5SMHaz3J61fN97r7sAiNGp1o7H",
			"NodeID-971kRywJGDu651QBGDJDPSDyth5NTb2Yy",
			"NodeID-7XmnU74M3APxBV6WLZS51NH9r7KhNoCTr",
			"NodeID-GGCBNGeqoLbHp6jhc4wtn5o92Eszj9HZt",
			"NodeID-M1ivUWVi9Ad2HxJcNcXTwoSfimLBhz8SK",
		}
	default:
		return nil
	}
}

// SampleBeacons returns the some beacons this node should connect to
func SampleBeacons(networkID uint32, count int) ([]string, []string) {
	ips := getIPs(networkID)
	ids := getNodeIDs(networkID)

	if numIPs := len(ips); numIPs < count {
		count = numIPs
	}

	sampledIPs := make([]string, 0, count)
	sampledIDs := make([]string, 0, count)

	s := sampler.NewUniform()
	_ = s.Initialize(uint64(len(ips)))
	indices, _ := s.Sample(count)
	for _, index := range indices {
		sampledIPs = append(sampledIPs, ips[int(index)])
		sampledIDs = append(sampledIDs, ids[int(index)])
	}

	return sampledIPs, sampledIDs
}
