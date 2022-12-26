# federated-eats

Welcome to federated-eats! This is a small demo that's an example of a fullstack supergraph built with a couple of simple Go subgraphs, [gqlgen](https://gqlgen.com/), [Apollo Federation 2](https://www.apollographql.com/docs/federation/quickstart/setup/), [Apollo GraphOS](https://www.apollographql.com/docs/graphos/), [Railway](https://railway.app/) deployments, [@defer](https://www.apollographql.com/docs/router/executing-operations/defer-support/), and [Next.js](https://nextjs.org/).

### Demo
|  |  |
| :-- | :-- |
| Locations Subgraph | https://subgraph-locations.up.railway.app/sandbox |
| Eateries Subgraph | https://subgraph-eateries.up.railway.app/sandbox |
| Federated Eats Supergraph | https://main--federated-eats.apollographos.net/graphql , https://studio.apollographql.com/public/federated-eats/explorer?variant=main |
| Complete Demo App | https://federated-eats.vercel.app/ |

## Local Development
```
$ make help

          go-mod-tidy-eateries  go mod tidy for subgraph_eateries
         go-mod-tidy-locations  go mod tidy for subgraph_locations
      gqlgen-generate-eateries  Generate models/code based on schema for subgraph_eateries
     gqlgen-generate-locations  Generate models/code based on schema for subgraph_locations
                          help  Display this help screen
              print-needed-env  Print env variables you must add
            rover-dev-eateries  Start a local Apollo Router that automatically composes the eateries schema
           rover-dev-locations  Start a local Apollo Router that automatically composes the locations schema
                  start-client  Start local client dev server
       start-eateries-subgraph  Start the eateries subgraph server
      start-locations-subgraph  Start the locations subgraph server
```