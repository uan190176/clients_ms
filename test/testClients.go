package main

import (
	"clients_ms/internal/api"
	"clients_ms/pkg/config"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const cliToke5 string = "43718bcf1382b1baa178f648295f396380ca593af4c10b872b8450dda6eae76d"

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
	//res, err := testGetClients(conn) //pass
	//res, err := testCreateClient(conn) //pass
	//res, err := testUpdateClient(conn) //pass
	//res, err := testUpdateDelFlagClients(conn) //pass

	//Misc
	res, err := testGetMiscData(conn) //

	fmt.Println(res, err)
}

func testGetMiscData(conn *grpc.ClientConn) (*api.ResponseMiscData, error) {
	client := api.NewClientsServicesClient(conn)
	ctx := context.Background()
	res, err := client.GetMiscDataToCreateAndUpdateClientAddress(ctx, &api.RequestMiscData{
		AuthToken: cliToke5,
		AuthorId:  7,
	})

	return res, err
}

func testGetClients(conn *grpc.ClientConn) (*api.ResponseClients, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.GetClients(context.Background(), &api.RequestClient{
		AuthToken: cliToke5,
		AuthorId:  7,
		//Id:        25,
		Address: "test",
	})
	return res, err
}

func testCreateClient(conn *grpc.ClientConn) (*api.ResponseClients, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.CreateClient(context.Background(), &api.RequestClient{
		AuthToken:  cliToke5,
		AuthorId:   7,
		FirstName:  "Александр",
		MiddleName: "Николаевич",
		LastName:   "Усольцев",
		Phone:      "79005553534",
		Email:      "nN9Xw@example.com",
	})
	return res, err
}

func testUpdateClient(conn *grpc.ClientConn) (*api.ResponseClients, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateClient(context.Background(), &api.RequestClient{
		AuthToken: cliToke5,
		Id:        30,
		AuthorId:  7,
		Comment:   "test clients updated",
	})
	return res, err
}

func testUpdateDelFlagClients(conn *grpc.ClientConn) (*api.ResponseClients, error) {
	c := api.NewClientsServicesClient(conn)

	a := make([]uint64, 0)
	a = append(a, 20)
	res, err := c.UpdateClientsDeletionFlags(context.Background(), &api.RequestClientsDeletionFlags{
		AuthToken: cliToke5,
		Ids:       []uint64{25, 30},
		AuthorId:  7,
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}
