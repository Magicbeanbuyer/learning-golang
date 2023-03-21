package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

type Tutorial struct {
	Id       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name":      &graphql.Field{Type: graphql.String},
			"Tutorials": &graphql.Field{Type: graphql.NewList(graphql.Int)},
		},
	},
)

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{Type: graphql.String},
		},
	},
)

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"Id":       &graphql.Field{Type: graphql.Int},
			"Title":    &graphql.Field{Type: graphql.String},
			"Author":   &graphql.Field{Type: authorType},
			"Comments": &graphql.Field{Type: graphql.NewList(commentType)},
		},
	},
)

// tutorial field which will allow us to retrieve individual Tutorials based on an ID passed in to the query.
// list field which will allow us to retrieve the full array of Tutorials that we have defined in memory.

func populate() []Tutorial {
	jane := &Author{
		Name:      "Jane",
		Tutorials: []int{1},
	}

	tutorial := Tutorial{
		Id:     1,
		Title:  "hello",
		Author: *jane,
		Comments: []Comment{
			{Body: "great"},
		},
	}
	tutorials := []Tutorial{
		tutorial,
	}
	return tutorials
}

func main() {
	fmt.Println("Learning GraphQL...")

	tutorials := populate()

	tutorialResolver := func(param graphql.ResolveParams) (interface{}, error) {
		id, ok := param.Args["id"].(int) //cast to int type
		if ok {
			for _, tutorial := range tutorials {
				if int(tutorial.Id) == id {
					return tutorial, nil
				}
			}
		}
		return nil, nil
	}

	listResolver := func(param graphql.ResolveParams) (interface{}, error) {
		return tutorials, nil
	}

	tutorialField := &graphql.Field{
		Type: tutorialType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: tutorialResolver,
	}

	listField := &graphql.Field{
		Name:    "list",
		Type:    graphql.NewList(tutorialType),
		Resolve: listResolver,
	}

	fields := graphql.Fields{
		"tutorial": tutorialField,
		"list":     listField,
	}

	rootQueryConf := graphql.ObjectConfig{
		Name:   "rootQuery",
		Fields: fields,
	}
	// main entry point for every graphql query coming into the graphql application
	rootQuery := graphql.NewObject(rootQueryConf)

	// what fields can we select? What kinds of objects might they return?
	//What fields are available on those sub-objects?
	schemaConf := graphql.SchemaConfig{
		Query: rootQuery,
	}

	schema, err := graphql.NewSchema(schemaConf)

	if err != nil {
		log.Fatal(err)
	}

	query := `
		{
			tutorial(id:1) {
				Title
				Author {
					Name
					Tutorials
				}
			}
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

// https://www.youtube.com/watch?v=AlLBG6HrE7E&t=1139s
