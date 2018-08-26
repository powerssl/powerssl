package cmd

import (
	"encoding/json"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	api_v1 "powerssl.io/pkg/api/v1"
)

const (
	address = "localhost:3000"
)

func ListDomains() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api_v1.NewDomainServiceClient(conn)
	log.Println(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.ListDomains(ctx, &api_v1.ListDomainsRequest{})
	if err != nil {
		log.Fatalf("could not list: %v", err)
	}
	b, err := json.MarshalIndent(response.GetDomains(), "", "  ")
	if err != nil {
		log.Fatalf("could not marshal: %v", err)
	}
	log.Printf("Domains:\n\n%s", string(b))
}
