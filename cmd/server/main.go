// Command server is the Whack authoritative game server.
package main

import (
	"fmt"
	"log"

	"github.com/opd-ai/whack/pkg/config"
	"github.com/opd-ai/whack/pkg/engine"
	"github.com/opd-ai/whack/pkg/network"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	world := engine.NewWorld()
	_ = world

	srv := network.NewServer(cfg.Server.Address)
	fmt.Printf("server: starting on %s (max %d players, %d tick/s)\n",
		cfg.Server.Address, cfg.Server.MaxPlayers, cfg.Server.TickRate)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
	defer srv.Stop()

	fmt.Println("server: listening")
	select {}
}
