package services

import (
	"clients_ms/internal/api"
)

// GrpcClientsServer - struct for gRPC server, link to generated instance
type GrpcClientsServer struct {
	api.ClientsServicesServer
}
