// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

const (
	// TestNetFaucetServerHost -- defualt testnet faucet server.
	TestNetFaucetServerHost string = "faucet.testnet.libra.org"
	// TestNetGRPCServerHost -- default testnet grpc server.
	TestNetGRPCServerHost string = "ac.testnet.libra.org:8000"
)

// Network --
type Network int

// network.
const (
	TestNet = iota
	MainNet
)
