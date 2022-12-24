package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"subgraph_eateries/graph"
	"subgraph_eateries/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jmatth11/yfusion"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8081"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	yelpApiKey := os.Getenv("YELP_API_KEY")
	yelpClient := yfusion.NewYelpFusion(yelpApiKey)
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{YelpClient: yelpClient}}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandbox)
	}))
	router.Handle("/graphql", srv)

	log.Printf("connect to %s/playground for GraphQL playground", url)
	log.Printf("connect to %s/sandbox for Apollo Studio Sandbox", url)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
