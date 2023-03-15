package main

import (
	"clients_ms/internal/api"
	"clients_ms/pkg/config"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const cliToken3 string = "43718bcf1382b1baa178f648295f396380ca593af4c10b872b8450dda6eae76d"

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
	//res, err := testGetNotes(conn) //pass
	//res, err := testCreateNote(conn) //pass
	//res, err := testUpdateNote(conn) //pass
	res, err := testUpdateDelFlagNotes(conn) //pass

	fmt.Println(res, err)
}

func testGetNotes(conn *grpc.ClientConn) (*api.ResponseNotes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.GetNotes(context.Background(), &api.RequestNote{
		AuthToken: cliToken3,
		AuthorId:  7,
		ClientId:  25,
	})
	return res, err
}

func testCreateNote(conn *grpc.ClientConn) (*api.ResponseNotes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.CreateNote(context.Background(), &api.RequestNote{
		AuthToken: cliToken3,
		ClientId:  25,
		NoteText:  "test notes",
		AuthorId:  7,
	})
	return res, err
}

func testUpdateNote(conn *grpc.ClientConn) (*api.ResponseNotes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateNote(context.Background(), &api.RequestNote{
		AuthToken: cliToken3,
		Id:        2,
		AuthorId:  7,
		ClientId:  25,
		NoteText:  "test notes updated",
	})
	return res, err
}

func testUpdateDelFlagNotes(conn *grpc.ClientConn) (*api.ResponseNotes, error) {
	c := api.NewClientsServicesClient(conn)

	a := make([]uint64, 0)
	a = append(a, 20)
	res, err := c.UpdateNotesDeletionFlags(context.Background(), &api.RequestNotesDeletionFlags{
		AuthToken: cliToken3,
		Ids:       []uint64{2},
		AuthorId:  7,
		IsDeleted: api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}
