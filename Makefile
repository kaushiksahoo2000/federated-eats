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

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%30s\033[0m  %s\n", $$1, $$2}'