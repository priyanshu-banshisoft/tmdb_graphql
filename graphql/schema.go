package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/priyanshu-banshisoft/tmdb_graphql/tmdb"
)



var Schema graphql.Schema

func init() {
    var queryType = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Query",
            Fields: graphql.Fields{
                "popularMovies": &graphql.Field{
                    Type: graphql.NewList(movieType),
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        return client.FetchPopularMovies()
                    },
                },
                "trendingMovies": &graphql.Field{
                    Type: graphql.NewList(movieType),
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        return client.FetchTrendingMovies()
                    },
                },
                "trendingTv" : &graphql.Field{
                    Type: graphql.NewList(tvType),
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        return client.FetchTrendingTv()
                    },
                },
                "genres" : &graphql.Field{
                    Type: graphql.NewList(genres),
                    Resolve: func (p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        return client.FetchMovieGenres()
                    },
                },
                "movie": &graphql.Field{
                    Type: movieType,
                    Args: graphql.FieldConfigArgument{
                        "id": &graphql.ArgumentConfig{
                            Type: graphql.Int,
                        },
                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        id, ok := p.Args["id"].(int)
                        if ok {
                            return client.FetchMovieDetails(id)
                        }
                        return nil, nil
                    },
                },
                "credit": &graphql.Field{
                    Type: credit,
                    Args: graphql.FieldConfigArgument{
                        "id" : &graphql.ArgumentConfig{
                            Type: graphql.Int,
                        },

                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        id, ok := p.Args["id"].(int)
                        if ok {
                            return client.FetchMovieCast(id)
                        }
                        return nil, nil
                    },
                },
            },
        },
    )

    var err error
    Schema, err = graphql.NewSchema(
        graphql.SchemaConfig{
            Query: queryType,
        },
    )
    if err != nil {
        panic(err)
    }
}
