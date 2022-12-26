#########################
## ENV
#########################
.PHONY: print-needed-env
print-needed-env: ## Print env variables you must add 
	@echo 'check env.example files in subgraph_locations and subgraph_eateries'
	@echo 'POSITIONSTACK_ACCESS_KEY=xyz'
	@echo 'YELP_API_KEY=xyz'

#########################
## CLIENT
#########################
.PHONY: start-client
start-client: ## Start local client dev server
	cd client && npm run dev

#########################
## LOCATIONS SUBGRAPH
#########################
.PHONY: start-locations-subgraph
start-locations-subgraph: ## Start the locations subgraph server
	cd subgraph_locations && go run ./server.go

.PHONY: go-mod-tidy-locations
go-mod-tidy-locations: ## go mod tidy for subgraph_locations
	cd subgraph_locations && go mod tidy

.PHONY: gqlgen-generate-locations
gqlgen-generate-locations: ## generate models/code based on schema for subgraph_locations
	cd subgraph_locations && go run github.com/99designs/gqlgen generate

.PHONY: rover-dev-locations
rover-dev-locations: ## start a local Apollo Router that automatically composes the locations schema
	rover dev --name locations --schema ./subgraph_locations/graph/schema.graphqls --url http://localhost:8080/graphql


#########################
## EATERIES SUBGRAPH
#########################

.PHONY: start-eateries-subgraph
start-eateries-subgraph: ## Start the eateries subgraph server
	cd subgraph_eateries && go run ./server.go

.PHONY: go-mod-tidy-eateries
go-mod-tidy-eateries: ## go mod tidy for subgraph_eateries
	cd subgraph_eateries && go mod tidy

.PHONY: gqlgen-generate-eateries
gqlgen-generate-eateries: ## generate models/code based on schema for subgraph_eateries
	cd subgraph_eateries && go run github.com/99designs/gqlgen generate

.PHONY: rover-dev-eateries
rover-dev-eateries: ## start a local Apollo Router that automatically composes the eateries schema
	rover dev --name eateries --schema ./subgraph_eateries/graph/schema.graphqls --url http://localhost:8081/graphql

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%30s\033[0m  %s\n", $$1, $$2}'