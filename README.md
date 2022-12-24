# federated-eats

Welcome to federated-eats! This is a small demo that's an example of a fullstack supergraph built with a couple of simple Go subgraphs, [gqlgen](https://gqlgen.com/), [Apollo Federation 2](https://www.apollographql.com/docs/federation/quickstart/setup/), [Apollo GraphOS](https://www.apollographql.com/docs/graphos/), [Railway](https://railway.app/) deployments, [@defer](https://www.apollographql.com/docs/router/executing-operations/defer-support/), and [Next.js](https://nextjs.org/).

### Demo
|  |  |
| :-- | :-- |
| Locations Subgraph | https://subgraph-locations.up.railway.app/sandbox |
| Eateries Subgraph | https://subgraph-eateries.up.railway.app/sandbox |
| Federated Eats Supergraph | https://subgraph-eateries.up.railway.app/sandbox |
| Demo | https://subgraph-eateries.up.railway.app/sandbox |

## Local Development
```
$ make help

          go-mod-tidy-eateries  go mod tidy for subgraph_eateries
         go-mod-tidy-locations  go mod tidy for subgraph_locations
      gqlgen-generate-eateries  generate models/code based on schema for subgraph_eateries
     gqlgen-generate-locations  generate models/code based on schema for subgraph_locations
                          help  Display this help screen
            rover-dev-eateries  start a local Apollo Router that automatically composes the eateries schema
           rover-dev-locations  start a local Apollo Router that automatically composes the locations schema
       start-eateries-subgraph  Start the eateries subgraph server
      start-locations-subgraph  Start the locations subgraph server
```