// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/MetalBlockchain/metalgo/ids"
)

// SubnetTracker describes the interface for checking if a node is tracking a
// subnet
type SubnetTracker interface {
	// TracksSubnet returns true if [nodeID] tracks [subnetID]
	TracksSubnet(nodeID ids.NodeID, subnetID ids.ID) bool
}
