package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

func main() {
	// type DrawflowDiagram struct {
	// 	ID             string `json: "id"`
	// 	Name           string `json: "name"`
	// 	DrawflowOutput string `json: "drawflow_output"`
	// }

	// p := DrawflowDiagram{
	// 	ID:             "1",
	// 	Name:           "Pony",
	// 	DrawflowOutput: "Content",
	// }

	conn, err := dgo.DialCloud("https://blue-surf-560027.us-east-1.aws.cloud.dgraph.io/graphql", "MjRkNGEwMTYyNmRiYmExZTZjY2FlYjQ3YmNiZTc2ZjM=")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	// txn := dgraphClient.NewTxn()
	// ctx := context.Background()

	// assigned, err := txn.Mutate(ctx, mu)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// commit_err := txn.Commit(ctx)
	// if err != nil {
	// 	log.Fatal(commit_err)
	// }
	// fmt.Print(assigned)

	run()
}

func run() {
	type DrawflowDiagram struct {
		ID             string `json: "id"`
		Name           string `json: "name"`
		DrawflowOutput string `json: "drawflow_output"`
	}

	p := DrawflowDiagram{
		ID:             "1",
		Name:           "Pony",
		DrawflowOutput: "Content",
	}

	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson: pb, CommitNow: true,
	}

	conn, err := dgo.DialCloud("https://blue-surf-560027.us-east-1.aws.cloud.dgraph.io/graphql", "MjRkNGEwMTYyNmRiYmExZTZjY2FlYjQ3YmNiZTc2ZjM=")
	if err != nil {
		log.Fatal(err)
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	txn := dgraphClient.NewTxn()
	ctx := context.Background()

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
