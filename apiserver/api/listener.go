package api

import (
	durationpb "github.com/golang/protobuf/ptypes/duration"
	"github.com/google/uuid"
)

// Listener API
type Listener struct {
	Name                             string                  `json:"name,omitempty"`
	Address                          *Address                `json:"address,omitempty"`
	FilterChains                     []FilterChain           `json:"filter_chains,omitempty"`
	DefaultFilterChain               FilterChain             `json:"default_filter_chain,omitempty"`
	PerConnectionBufferLimitBytes    uint32                  `json:"per_connection_buffer_limit_bytes,omitempty"`
	Metadata                         Metadata                `json:"metadata,omitempty"`
	DrainType                        DrainType               `json:"drain_type,omitempty"`
	ListenerFilters                  []ListenerFilter        `json:"listener_filters,omitempty"`
	ListenerFiltersTimeout           durationpb.Duration     `json:"listener_filters_timeout,omitempty"`
	ContinueOnListenerFiltersTimeout bool                    `json:"continue_on_listener_filters_timeout,omitempty"`
	Transparent                      bool                    `json:"transparent,omitempty"`
	Freebind                         bool                    `json:"freebind,omitempty"`
	SocketOption                     SocketOption            `json:"socket_option,omitempty"`
	TcpFastOpenQueueLength           uint32                  `json:"tcp_fast_open_queue_length,omitempty"`
	TrafficDirection                 TrafficDirection        `json:"traffic_direction,omitempty"`
	UdpListenerConfig                UdpListenerConfig       `json:"udp_listener_config,omitempty"`
	ApiListener                      ApiListener             `json:"api_listener,omitempty"`
	ConnectionBalanceConfig          ConnectionBalanceConfig `json:"connection_balance_config,omitempty"`
	ReusePort                        bool                    `json:"reuse_port,omitempty"`
	AccessLog                        AccessLog               `json:"access_log,omitempty"`
	TcpBacklogSize                   uint32                  `json:"tcp_backlog_size,omitempty"`
}

func (l *Listener) CheckValues() error {
	if l.Name == "" {
		l.Name = uuid.New().String()
	}
	panic("todo")
}

// Address listen addr
type Address struct {
	SocketAddress SocketAddress
	Pipe          Pipe
}

type SocketAddress struct{}

type Pipe struct{}

type FilterChain struct {
}

type Metadata struct {
}

type DrainType struct{}

type ListenerFilter struct{}

type SocketOption struct{}

type TrafficDirection struct{}

type UdpListenerConfig struct{}
type ApiListener struct{}
type ConnectionBalanceConfig struct{}
type AccessLog struct{}
