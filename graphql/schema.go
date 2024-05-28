package graphql

import (
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/graphql-go/graphql"
	"github.com/priyanshu-banshisoft/tmdb_graphql/tmdb"
)

var Schema graphql.Schema

func init() {
	var loginMutation = &graphql.Field{
		Type: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Response",
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.String,
					},
                    "username" : &graphql.Field{ 
                        Type: graphql.String,
                    },
                    "email" : &graphql.Field{ 
                        Type: graphql.String,
                    },
                    "lastName" : &graphql.Field{ 
                        Type: graphql.String,
                    },
                    "firstName" : &graphql.Field{ 
                        Type: graphql.String,
                    },
                    "gender" : &graphql.Field{ 
                        Type: graphql.String,
                    },
                    "image" : &graphql.Field{
                        Type: graphql.String,
                    },
                    "token" : &graphql.Field{
                        Type: graphql.String,
                    },
				},
			},
		),
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			username := params.Args["username"].(string)
			password := params.Args["password"].(string)
			client := resty.New()
			resp, err := client.R().
				SetBody(map[string]string{
					"username": username,
					"password": password,
				}).
				SetResult(&tmdb.User{}).
				Post("https://dummyjson.com/auth/login")

			if err != nil {
				return nil, err
			}

			if resp.IsError() {
				return nil, errors.New("invalid username or password")
			}

			authResponse := resp.Result().(*tmdb.User)
			return authResponse, nil
		},
	}
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
				"trendingTv": &graphql.Field{
					Type: graphql.NewList(tvType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						apiKey := p.Context.Value("apiKey").(string)
						client := tmdb.NewClient(apiKey)
						return client.FetchTrendingTv()
					},
				},
				"genres": &graphql.Field{
					Type: graphql.NewList(genres),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
						"id": &graphql.ArgumentConfig{
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

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"login": loginMutation,
		},
	})

	var err error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: rootMutation,
		},
	)
	if err != nil {
		panic(err)
	}
}
