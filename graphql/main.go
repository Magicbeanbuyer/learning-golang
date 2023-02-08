package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

func main() {
	fmt.Println("Learning GraphQL...")

	resolver := func(param graphql.ResolveParams) (interface{}, error) {
		// dummy http request
		return "world", nil
	}

	helloFiled := &graphql.Field{
		Type:    graphql.String,
		Resolve: resolver,
	}

	fields := graphql.Fields{
		"hello": helloFiled,
	}

	rootQueryConf := graphql.ObjectConfig{
		Name:   "rootQuery",
		Fields: fields,
	}
	// main entry point for every graphql query coming into the graphql application
	rootQuery := graphql.NewObject(rootQueryConf)

	schemaConf := graphql.SchemaConfig{
		Query: rootQuery,
	}

	schema, err := graphql.NewSchema(schemaConf)

	if err != nil {
		log.Fatal(err)
	}

	query := `
		{
			hello
		}
	`
	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}

	resp := graphql.Do(params)
	if len(resp.Errors) > 0 {
		log.Fatal(resp.Errors)
	}

	responseJson, _ := json.Marshal(resp)

	fmt.Printf("%s \n", responseJson)
}
