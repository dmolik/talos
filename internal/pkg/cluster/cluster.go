// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cluster provides functions to access, check and inspect Talos clusters.
package cluster

import (
	"context"
	"io"

	"k8s.io/client-go/kubernetes"

	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
	"github.com/talos-systems/talos/pkg/client"
)

// ClientProvider builds Talos client by endpoint.
//
// Client instance should be cached and closed when Close() is called.
type ClientProvider interface {
	// Client returns Talos client instance for default (if no endpoints are given) or
	// specific endpoint.
	Client(endpoints ...string) (*client.Client, error)
	// Close client connections.
	Close() error
}

// K8sProvider builds Kubernetes client to access Talos cluster.
type K8sProvider interface {
	K8sClient(ctx context.Context) (*kubernetes.Clientset, error)
}

// CrashDumper captures Talos cluster state to the specified writer for debugging.
type CrashDumper interface {
	CrashDump(ctx context.Context, out io.Writer)
}

// Info describes the Talos cluster.
type Info interface {
	// Nodes returns list of all node endpoints (IPs).
	Nodes() []string
	// NodesByType return list of node endpoints by type.
	NodesByType(runtime.MachineType) []string
}
