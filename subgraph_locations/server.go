package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"subgraph_locations/graph"
	"subgraph_locations/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	positionstackAccessKey := os.Getenv("POSITIONSTACK_ACCESS_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	defaultUrl := fmt.Sprintf("http://localhost:%s", port)
	url := os.Getenv("RAILWAY_STATIC_URL")
	if url == "" {
		url = defaultUrl
	} else {
		url = fmt.Sprintf("https://%s", url)
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{defaultUrl, "https://studio.apollographql.com"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// new Resty Client
	client := resty.New()
	// sandbox html
	var sandbox = []byte(fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
		<div id="sandbox" style="height:100vh; width:100vw;"></div>
		<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
		<script>
		new window.EmbeddedSandbox({
		target: "#sandbox",
		// Pass through your server href if you are embedding on an endpoint.
		// Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
		initialEndpoint: "%s/graphql",
		});
		// advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
		</script>
		</body>
		
		</html>`, url))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{RestyClient: client, PositionstackAccessKey: positionstackAccessKey}}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandbox)
	}))
	router.Handle("/graphql", srv)

	log.Printf("connect to %s/playground for GraphQL playground", url)
	log.Printf("connect to %s/sandbox for Apollo Studio Sandbox", url)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
