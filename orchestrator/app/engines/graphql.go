package engine

import (
	"net/http"

	graph "github.com/bharath0292/quantdrey/graph/generated"
	resolver "github.com/bharath0292/quantdrey/graph/resolvers"
	strategyservice "github.com/bharath0292/quantdrey/internal/domains/strategy/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

type GrapQLEngine struct {
	server     *handler.Server
	playground http.HandlerFunc
}

func NewGraphQLEngine(strategyService strategyservice.IStrategyService) *GrapQLEngine {
	res := resolver.NewResolver(strategyService)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &res,
	}))

	/*
		- CORS preflight (OPTIONS) requests — important for browser compatibility
		- GET requests — useful for query caching and playground tools
		- POST requests — standard method for sending GraphQL queries/mutations
	*/
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	/*
		- Enable schema introspection queries
		- This allows tools like GraphQL Playground, Apollo Studio, or frontend dev tools to explore the schema
	*/
	srv.Use(extension.Introspection{})

	/*
		- Set up in-memory cache for parsed query documents using an LRU (Least Recently Used) cache
		- Clients can send a hash instead of the full query to reduce payload size and improve performance
	*/
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	playground := playground.Handler("GraphQL Playground", "/server/query")

	return &GrapQLEngine{
		server:     srv,
		playground: playground,
	}
}

func (g *GrapQLEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.server.ServeHTTP(w, r)
}

func (g *GrapQLEngine) PlaygroundHandler(w http.ResponseWriter, r *http.Request) {
	g.playground(w, r)
}
