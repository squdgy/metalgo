// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/metalgo/message"
	"github.com/MetalBlockchain/metalgo/snow/networking/handler"
	"github.com/MetalBlockchain/metalgo/snow/networking/timeout"
	"github.com/MetalBlockchain/metalgo/trace"
	"github.com/MetalBlockchain/metalgo/utils/logging"
	"github.com/MetalBlockchain/metalgo/version"
)

var _ Router = (*tracedRouter)(nil)

type tracedRouter struct {
	router Router
	tracer trace.Tracer
}

func Trace(router Router, tracer trace.Tracer) Router {
	return &tracedRouter{
		router: router,
		tracer: tracer,
	}
}

func (r *tracedRouter) Initialize(
	nodeID ids.NodeID,
	log logging.Logger,
	timeoutManager timeout.Manager,
	closeTimeout time.Duration,
	criticalChains ids.Set,
	whitelistedSubnets ids.Set,
	onFatal func(exitCode int),
	healthConfig HealthConfig,
	metricsNamespace string,
	metricsRegisterer prometheus.Registerer,
) error {
	return r.router.Initialize(
		nodeID,
		log,
		timeoutManager,
		closeTimeout,
		criticalChains,
		whitelistedSubnets,
		onFatal,
		healthConfig,
		metricsNamespace,
		metricsRegisterer,
	)
}

func (r *tracedRouter) RegisterRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	requestingChainID ids.ID,
	respondingChainID ids.ID,
	requestID uint32,
	op message.Op,
	failedMsg message.InboundMessage,
) {
	r.router.RegisterRequest(
		ctx,
		nodeID,
		requestingChainID,
		respondingChainID,
		requestID,
		op,
		failedMsg,
	)
}

func (r *tracedRouter) HandleInbound(ctx context.Context, msg message.InboundMessage) {
	ctx, span := r.tracer.Start(ctx, "tracedRouter.HandleInbound")
	defer span.End()

	r.router.HandleInbound(ctx, msg)
}

func (r *tracedRouter) Shutdown() {
	r.router.Shutdown()
}

func (r *tracedRouter) AddChain(chain handler.Handler) {
	r.router.AddChain(chain)
}

func (r *tracedRouter) Connected(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	r.router.Connected(nodeID, nodeVersion, subnetID)
}

func (r *tracedRouter) Disconnected(nodeID ids.NodeID) {
	r.router.Disconnected(nodeID)
}

func (r *tracedRouter) Benched(chainID ids.ID, nodeID ids.NodeID) {
	r.router.Benched(chainID, nodeID)
}

func (r *tracedRouter) Unbenched(chainID ids.ID, nodeID ids.NodeID) {
	r.router.Unbenched(chainID, nodeID)
}

func (r *tracedRouter) HealthCheck() (interface{}, error) {
	return r.router.HealthCheck()
}
