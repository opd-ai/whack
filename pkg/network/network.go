// Package network provides TCP server, client prediction, and delta compression for multiplayer.
package network

import "net"

// InputFrame represents a client input packet.
type InputFrame struct {
	Frame   uint32
	Buttons uint16
	StickX  int8
	StickY  int8
}

// Server is the authoritative game server.
type Server struct {
	Address  string
	listener net.Listener
}

// NewServer creates a new game server.
func NewServer(address string) *Server {
	return &Server{Address: address}
}

// Start begins listening for client connections.
func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	return nil
}

// Stop shuts down the server.
func (s *Server) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

// Client handles connection to a game server.
type Client struct {
	ServerAddress string
	conn          net.Conn
}

// NewClient creates a new game client.
func NewClient(serverAddress string) *Client {
	return &Client{ServerAddress: serverAddress}
}

// NetworkSyncSystem synchronizes game state over the network.
type NetworkSyncSystem struct {
	// Skeleton: will handle state synchronization.
}
