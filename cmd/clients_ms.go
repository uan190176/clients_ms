package main

import (
	"clients_ms/internal/api"
	"clients_ms/internal/services"
	stg "clients_ms/internal/storage/utils"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	defaultCfgPath = "clients.cfg.toml"
)

var (
	Version = "1.1.0"
	cfgPath string
)

// init - starting once before main
func init() {
	flag.StringVar(&cfgPath, "config-path", "clients.cfg.toml", "path to config file")
}

func main() {
	// Parse parameters on starting app
	// if empty, use default
	flag.Parse()
	if cfgPath == "" {
		cfgPath = defaultCfgPath
	}

	// Init config
	_cfg, err := cfg.NewConfig(cfgPath)
	if err != nil {
		os.Exit(1)
	}

	// Init logger
	_ = lgr.GetLogger(_cfg)
	lgr.LOG.Info("===", "custom logger is init")

	//Ping database
	lgr.LOG.Info("_ACTION_: ", "testing connection to database")
	if !stg.PingDatabase(context.Background()) {
		lgr.LOG.Fatal("_FATAL_: ", "not enough connect to database")
		os.Exit(1)
	}

	// Init gRPC server
	lgr.LOG.Info("_ACTION_: ", "starting main gRPC server")
	grpcServer := grpc.NewServer()

	lgr.LOG.Info("_ACTION_: ", "starting microservice to manage clients")
	clientsServer := &services.GrpcClientsServer{}

	lgr.LOG.Info("_ACTION_: ", "registration microservice in main server")
	api.RegisterClientsServicesServer(grpcServer, clientsServer)

	// Init listener
	lgr.LOG.Info("_ACTION_: ", "initialising listener")
	listener, err := net.Listen("tcp", cfg.CFG.Server.BindAddress)
	if err != nil {
		lgr.LOG.Fatal("error while running listener")
		lgr.LOG.Fatal(err)
	}

	lgr.LOG.Info(fmt.Sprintf("clients microservice (v.%s) is started, listening address %s", Version, cfg.CFG.Server.BindAddress))
	fmt.Printf("clients microservice (v.%s) is started, listening address %s", Version, cfg.CFG.Server.BindAddress)
	if err = grpcServer.Serve(listener); err != nil {
		lgr.LOG.Fatal("_FATAL_: ", "error while initialising listener")
		lgr.LOG.Fatal("_FATAL_: ", err)
		os.Exit(1)
	}

}
