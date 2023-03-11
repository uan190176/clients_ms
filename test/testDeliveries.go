package main

import (
	"clients_ms/internal/api"
	"clients_ms/pkg/config"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const cliToken2 string = "43718bcf1382b1baa178f648295f396380ca593af4c10b872b8450dda6eae76d"

func main() {

	// Init config
	cfg, err := cfg.NewConfig("clients.cfg.toml")
	if err != nil {
		panic(err)
	}

	// Init gRPC
	conn, err := grpc.Dial(cfg.Server.BindAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	//Deliveries
	res, err := testGetDeliveries(conn) //pass
	//res, err := testCreateDelivery(conn) //pass
	//res, err := testUpdateDelivery(conn) //pass
	//res, err := testUpdateDelFlagDeliveries(conn) //pass

	fmt.Println(res, err)
}

func testGetDeliveries(conn *grpc.ClientConn) (*api.ResponseDeliveries, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.GetDeliveries(context.Background(), &api.RequestDelivery{
		AuthToken: cliToken2,
		AuthorId:  7,
		//CountryName: "рос",
		//CountryId: 2,
		//CountryIsDeleted: api.BooleanField_IS_TRUE,
	})
	return res, err
}

func testCreateDelivery(conn *grpc.ClientConn) (*api.ResponseDeliveries, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.CreateDelivery(context.Background(), &api.RequestDelivery{
		AuthToken: cliToken2,
		Name:      "test delivery",
		Comment:   "need deleted",
		AuthorId:  7,
	})
	return res, err
}

func testUpdateDelivery(conn *grpc.ClientConn) (*api.ResponseDeliveries, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateDelivery(context.Background(), &api.RequestDelivery{
		AuthToken: cliToken2,
		Id:        9,
		//CountryName: "Россия",
		Comment:  "*",
		AuthorId: 7,
		//IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}

func testUpdateDelFlagDeliveries(conn *grpc.ClientConn) (*api.ResponseDeliveries, error) {
	c := api.NewClientsServicesClient(conn)

	a := make([]uint64, 0)
	a = append(a, 20)
	res, err := c.UpdateDeliveriesDeletionFlags(context.Background(), &api.RequestDeliveriesDeletionFlags{
		AuthToken: cliToken2,
		Ids:       []uint64{9},
		AuthorId:  7,
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}
