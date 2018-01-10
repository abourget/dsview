package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

// Make sure you export:
//   DATASTORE_HOST or DATASTORE_EMULATOR_HOST,
// and
//   DATASTORE_DATASET or DATASTORE_PROJECT_ID,
//
var kind = flag.String("kind", "", "The Datastore Kind to query")
var objKey = flag.String("key", "", "Simple Key to query")

func main() {
	flag.Parse()

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Fatalln("Couldn't setup Datastore client:", err)
	}
	k := datastore.NameKey(*kind, *objKey, nil)

	fmt.Println("Querying database...")
	var data interface{}
	if err := client.Get(context.Background(), k, &data); err != nil {
		log.Fatalln("Failed querying the thing:", err)
	}

	cnt, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalln("Couldn't marshal the data thing:", err)
	}

	fmt.Println("Output:")
	fmt.Print(string(cnt))

}
