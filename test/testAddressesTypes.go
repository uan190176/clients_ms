package main

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
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

	start := time.Now()

	//var wg sync.WaitGroup
	//
	//inc := 50
	//wg.Add(inc)
	//cnt := 0
	//for i := 0; i < inc; i++ {
	//	go func(i int) {
	//		res, err := testUpdateAddressType(conn) //
	//		//err := ping(context.Background())
	//		//if err != nil {
	//		//	fmt.Println(err)
	//		//} else {
	//		//	fmt.Println("ping")
	//		//}
	//		cnt++
	//		fmt.Println(i, res, err, cnt)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	//for i := 0; i < 50; i++ {
	//	res, err := testUpdateAddressType(conn) //
	//	fmt.Println(i, res, err)
	//}

	//fmt.Println(cnt)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

	res, err := testGetAddressesTypes(conn) //pass
	//res, err := testCreateAddressType(conn) //pass
	//res, err := testUpdateAddressType(conn) //pass
	//res, err := testUpdateDeletedFlagAddressesTypes(conn) //pass

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
		AuthToken:       cliToken1,
		AuthorId:        7,
		AddrTypeName:    "до востребования",
		AddrTypeComment: "-",
		AddrTypeCode:    "demand",
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
		AddrTypeID:               10,
		AddrTypeNeedsNormalizing: api.ClientsMS_Bool_IS_TRUE,
		AddrTypeNeedsCleaning:    api.ClientsMS_Bool_IS_TRUE,
	})
	return res, err
}

func testUpdateDeletedFlagAddressesTypes(conn *grpc.ClientConn) (*api.ResponseAddressesTypes, error) {
	c := api.NewClientsServicesClient(conn)

	res, err := c.UpdateAddressTypeDeletionFlags(context.Background(), &api.RequestAddressesTypesDeletionFlags{
		AuthToken:         cliToken1,
		AuthorId:          7,
		AddrTypeID:        []uint64{9, 10},
		AddrTypeIsDeleted: api.ClientsMS_Bool_IS_FALSE,
	})
	return res, err
}
