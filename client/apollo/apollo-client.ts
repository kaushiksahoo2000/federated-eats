import { ApolloClient, InMemoryCache } from '@apollo/client'

const client = new ApolloClient({
  uri: 'https://main--federated-eats.apollographos.net/graphql',
  cache: new InMemoryCache(),
})

export default client
