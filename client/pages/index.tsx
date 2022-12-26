import type { NextPage } from 'next'
import { useQuery, gql } from '@apollo/client'
import Card from '../components/card'
import ClientOnly from '../apollo/client-only'
import Link from 'next/link'
import Table from '../components/table'

const QUERY = gql`
  query ($locationId: ID!) {
    ipLocation(id: $locationId) {
      continent
      locality
      county
      label
      latitude
      longitude
      ... @defer {
        eateriesForLocation {
          name
          rating
          id
        }
      }
    }
  }
`
const SPACEX = gql`
  query GetLaunches {
    launchesPast(limit: 10) {
      id
      mission_name
      launch_date_local
      launch_site {
        site_name_long
      }
      links {
        article_link
        video_link
        mission_patch
      }
      rocket {
        rocket_name
      }
    }
  }
`

const Home: NextPage = () => {
  const dummydata = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
  // const { data, loading, error } = useQuery(QUERY)
  const { data, loading, error } = useQuery(SPACEX)
  console.log({ data, loading, error })
  return (
    <>
      <section className="text-gray-700">
        <div className="px-4 py-16 sm:px-6 lg:px-8 space-y-10">
          <div className="max-w-xl space-y-10">
            <h2 className="text-3xl font-medium sm:text-4xl">Federated Eats</h2>
            <p className="mt-4">
              A simple demo utilizing the positionstack and yelp APIs via GraphQL. Built to show Apollo Federation 2, Apollo GraphOS, @defer, and
              more. Check out the{' '}
              <Link className="text-blue-500" target="_blank" href={'https://github.com/kaushiksahoo2000/federated-eats'}>
                repo
              </Link>{' '}
              for more of a breakdown 🚀.
            </p>
          </div>
          <Table />
          <ClientOnly>
            <div className="space-y-10">
              <div>
                <h3 className="text-lg font-medium sm:text-lg">Some current location information via the supergraph 📍:</h3>
                <p className="mt-4 text-sm text-gray-500">continent - </p>
                <p className="mt-4 text-sm text-gray-500">county - </p>
                <p className="mt-4 text-sm text-gray-500">label - </p>
                <p className="mt-4 text-sm text-gray-500">latitude - </p>
                <p className="mt-4 text-sm text-gray-500">longitude - </p>
              </div>
              <h3 className="text-lg font-medium sm:text-lg">Some food in the area via the supergraph and @defer 🍟:</h3>
              <div className="mt-8 grid grid-cols-1 gap-8 md:mt-16 md:grid-cols-2 md:gap-12 lg:grid-cols-3">
                {dummydata.map((x) => (
                  <Card />
                ))}
              </div>
            </div>
          </ClientOnly>
        </div>
      </section>
    </>
  )
}

export default Home
