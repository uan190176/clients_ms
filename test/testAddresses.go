package main

import (
	"clients_ms/internal/api"
	"clients_ms/pkg/config"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const cliToken4 string = "43718bcf1382b1baa178f648295f396380ca593af4c10b872b8450dda6eae76d"

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
	//res, err := testGetAddresses(conn) //pass
	res, err := testCreateAddress(conn) //pass
	//res, err := testUpdateAddress(conn) //pass
	//res, err := testUpdateDelFlagAddresses(conn) //pass

	fmt.Println(res, err)
}

func testGetAddresses(conn *grpc.ClientConn) (*api.ResponseAddresses, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.GetAddresses(context.Background(), &api.RequestAddress{
		AuthToken: cliToken4,
		AuthorId:  7,
		ClientId:  25,
	})
	return res, err
}

func testCreateAddress(conn *grpc.ClientConn) (*api.ResponseAddresses, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.CreateAddress(context.Background(), &api.RequestAddress{
		AuthToken:     cliToken4,
		ClientId:      25,
		AuthorId:      7,
		Address:       "Миасское ул Комсомола д 60 кв 1",
		AddressTypeId: 9,
		DeliveryId:    9,
		CountryId:     19,
	})
	return res, err
}

func testUpdateAddress(conn *grpc.ClientConn) (*api.ResponseAddresses, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateAddress(context.Background(), &api.RequestAddress{
		AuthToken: cliToken4,
		Id:        40,
		AuthorId:  7,
		Address:   "test address updated",
		PostCode:  "888888",
	})
	return res, err
}

func testUpdateDelFlagAddresses(conn *grpc.ClientConn) (*api.ResponseAddresses, error) {
	c := api.NewClientsServicesClient(conn)

	a := make([]uint64, 0)
	a = append(a, 20)
	res, err := c.UpdateAddressesDeletionFlags(context.Background(), &api.RequestAddressesDeletionFlags{
		AuthToken: cliToken4,
		Ids:       []uint64{43, 45},
		AuthorId:  7,
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}
