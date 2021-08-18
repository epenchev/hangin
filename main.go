package main

import (
	"log"
	"time"
	"fmt"

	"github.com/coreos/etcd/embed"
	//"github.com/coreos/pkg/capnslog"
)

func main() {
	cfg := embed.NewConfig()
	cfg.Dir = "default.etcd"
	//capnslog.SetGlobalLogLevel(capnslog.CRITICAL)
	cfg.LogPkgLevels = "*=E"
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", cfg.LogOutput);
	defer e.Close()
	select {
	case <-e.Server.ReadyNotify():
		log.Printf("Server is ready!")
	case <-time.After(60 * time.Second):
		e.Server.Stop() // trigger a shutdown
		log.Printf("Server took too long to start!")
	}
	log.Fatal(<-e.Err())
}
