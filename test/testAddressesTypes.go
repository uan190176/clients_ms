package main

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const cliToken1 string = "43718bcf1382b1baa178f648295f396380ca593af4c10b872b8450dda6eae76d"

func main() {

	// Init config
	cfg, err := cfg.NewConfig("clients.cfg.toml")
	if err != nil {
		panic(err)
	}

	// Init gRPC
	conn, err := grpc.Dial(cfg.Server.BindAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)

	//res, err := testGetAddressesTypes(conn) //pass
	//res, err := testCreateAddressType(conn) //pass
	//res, err := testUpdateAddressType(conn) //pass
	res, err := testUpdateDeletedFlagAddressesTypes(conn) //pass

	fmt.Println(res, err)
}

func testGetAddressesTypes(conn *grpc.ClientConn) (*api.ResponseAddressesTypes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.GetAddressesTypes(context.Background(), &api.RequestAddressType{
		AuthToken: cliToken1,
		AuthorId:  7,
		//CountryName: "рос",
		//CountryId: 2,
		//CountryIsDeleted: api.BooleanField_IS_TRUE,
	})
	return res, err
}

func testCreateAddressType(conn *grpc.ClientConn) (*api.ResponseAddressesTypes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.CreateAddressType(context.Background(), &api.RequestAddressType{
		AuthToken: cliToken1,
		AuthorId:  7,
		Name:      "тестовый",
		Comment:   "-",
		Code:      "test",
	})
	return res, err
}

func testUpdateAddressType(conn *grpc.ClientConn) (*api.ResponseAddressesTypes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateAddressType(context.Background(), &api.RequestAddressType{
		AuthToken: cliToken1,
		AuthorId:  7,
		//AddrTypeName: "обычный",
		//AddrTypeCode: "default",
		Id:               11,
		NeedsNormalizing: api.ClientsMS_Bool_IS_TRUE,
		NeedsCleaning:    api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}

func testUpdateDeletedFlagAddressesTypes(conn *grpc.ClientConn) (*api.ResponseAddressesTypes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateAddressTypeDeletionFlags(context.Background(), &api.RequestAddressesTypesDeletionFlags{
		AuthToken: cliToken1,
		AuthorId:  7,
		Ids:       []uint64{11},
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}
