// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/metalgo/snow/networking/router"
	"github.com/MetalBlockchain/metalgo/version"
)

var _ router.ExternalHandler = (*testHandler)(nil)

type testHandler struct {
	router.InboundHandler
	ConnectedF    func(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID)
	DisconnectedF func(nodeID ids.NodeID)
}

func (h *testHandler) Connected(id ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	if h.ConnectedF != nil {
		h.ConnectedF(id, nodeVersion, subnetID)
	}
}

func (h *testHandler) Disconnected(id ids.NodeID) {
	if h.DisconnectedF != nil {
		h.DisconnectedF(id)
	}
}
