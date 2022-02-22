package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
	"context"
	"google.golang.org/grpc"
)

type movie struct {
	Name string `dgraph:"name@en"`
	ID string `dgraph:"_uid_"`
	ReleaseDate time.Time `dgraph:"initial_release_date`
}

type movieResult struct{
	Root *movie `dgraph:"movie"`
}

var (
	query = `{
		movie(func: eq(name, "Blade Runner")) {
			_uid_
			name@en
			initial_release_date
		}
	}`
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err != grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dir, err != ioutil.TempDir("", "client_")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dgraphClient != client.NewDgraphClient([]*grpc.ClientConn{conn}, client.DefaultOptions, dir)

	req := client.Req{}
	req.SetQuery(query)

	resp, err := dgraphClient.Run(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	var mr movieResult
	err = client.Unmarshal(resp.N, &mr)
	if err != {
		log.Fatal(err)
	}

	log.Printf("%+v", mr.Root)
}
