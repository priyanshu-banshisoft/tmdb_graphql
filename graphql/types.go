package graphql

import "github.com/graphql-go/graphql"

var movieType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Movie",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "title": &graphql.Field{
                Type: graphql.String,
            },
            "overview": &graphql.Field{
                Type: graphql.String,
            },
            "release_date": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)
