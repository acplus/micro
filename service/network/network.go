// Package network implements micro network node
package network

import (
	"github.com/micro/go-micro/v3/metrics"
	"github.com/micro/go-micro/v3/metrics/noop"
	"github.com/micro/go-micro/v3/network"
	"github.com/micro/go-micro/v3/network/mucp"
)

var (
	// DefaultMetricsReporter implementation
	DefaultMetricsReporter metrics.Reporter = noop.New()
	// DefaultNetwork implementation
	DefaultNetwork network.Network = mucp.NewNetwork()
)
