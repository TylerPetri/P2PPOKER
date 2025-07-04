package main

import (
	p2p "ggpoker/p2p"
	"log"
	"time"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:     "GGPOKER V0.1-alpha",
		ListenAddr:  ":3000",
		GameVariant: p2p.TexasHoldem,
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	time.Sleep(1 * time.Second)

	remoteCfg := p2p.ServerConfig{
		Version:     "GGPOKER V0.1-alpha",
		ListenAddr:  ":4000",
		GameVariant: p2p.TexasHoldem,
	}
	remoteServer := p2p.NewServer(remoteCfg)
	go remoteServer.Start()
	if err := remoteServer.Connect(":3000"); err != nil {
		log.Fatal(err)
	}

	otherCfg := p2p.ServerConfig{
		Version:     "GGPOKER V0.1-alpha",
		ListenAddr:  ":3001",
		GameVariant: p2p.TexasHoldem,
	}
	otherServer := p2p.NewServer(otherCfg)
	go otherServer.Start()
	if err := otherServer.Connect(":4000"); err != nil {
		log.Fatal(err)
	}

	select {} // for blocking, to keep connection live for now

	// fmt.Println(deck.New())
}
