package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/briankscheong/go-graphql-gateway/graph"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/ast"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Get k8s cluster config
	log.Info().Msg("Getting in-cluster config to initialize clienset...")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Err(err).Msg("In-cluster config not found")
		panic(err)
	}

	log.Info().Msg("Successfully retrieved in-cluster config. Using config to create Kubernetes clienset...")

	// Create new k8s clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Successfully created Kubernetes clienset")

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		K8sClient: clientset,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info().Msgf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Err(http.ListenAndServe(":"+port, nil))
}
