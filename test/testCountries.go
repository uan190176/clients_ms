package main

import (
	"clients_ms/internal/api"
	"clients_ms/pkg/config"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const cliToken string = "43718bcf1382b1baa178f648295f396380ca593af4c10b872b8450dda6eae76d"

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

	//Countries
	//res, err := testGetCountries(conn) //pass
	//res, err := testCreateCountry(conn) //pass
	//res, err := testUpdateCountry(conn) //pass
	res, err := testUpdateDelFlagCountry(conn) //pass

	fmt.Println(res, err)
}

func testGetCountries(conn *grpc.ClientConn) (*api.ResponseCountries, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.GetCountries(context.Background(), &api.RequestCountry{
		AuthToken: cliToken,
		AuthorId:  7,
		//CountryName: "рос",
		//CountryId: 2,
		//CountryIsDeleted: api.BooleanField_IS_TRUE,
	})
	return res, err
}

func testCreateCountry(conn *grpc.ClientConn) (*api.ResponseCountries, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.CreateCountry(context.Background(), &api.RequestCountry{
		AuthToken: cliToken,
		Name:      "РФ",
		//CountryComment: "need deleted",
		AuthorId: 7,
	})
	return res, err
}

func testUpdateCountry(conn *grpc.ClientConn) (*api.ResponseCountries, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateCountry(context.Background(), &api.RequestCountry{
		AuthToken: cliToken,
		Id:        31,
		//CountryName: "Россия",
		//CountryComment: "need deleted",
		AuthorId:  7,
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}

func testUpdateDelFlagCountry(conn *grpc.ClientConn) (*api.ResponseCountries, error) {
	c := api.NewClientsServicesClient(conn)

	a := make([]uint64, 0)
	a = append(a, 20)
	res, err := c.UpdateCountriesDeletionFlags(context.Background(), &api.RequestCountriesDeletionFlags{
		AuthToken: cliToken,
		Ids:       []uint64{20, 22, 24, 31},
		AuthorId:  7,
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}
