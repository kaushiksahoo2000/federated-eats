import type { NextPage } from 'next'
import { useEffect, useState } from 'react'
import { useQuery, gql } from '@apollo/client'
import Card from '../components/card'
import Link from 'next/link'
import Table from '../components/table'

const QUERY = gql`
  query ($locationId: ID!) {
    location(id: $locationId) {
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
          URL
          reviewCount
          distance
        }
      }
    }
  }
`

const Home: NextPage = () => {
  const [latitude, setLatitude] = useState<number>()
  const [longitude, setLongitude] = useState<number>()

  useEffect(() => {
    navigator.geolocation.getCurrentPosition((position) => {
      setLatitude(position?.coords?.latitude)
      setLongitude(position?.coords?.longitude)
    })
  }, [])

  const validLocation = latitude && longitude
  const { data, error } = useQuery(QUERY, { skip: !validLocation, variables: { locationId: `${latitude},${longitude}` } })
  let loadingText = '...loading'
  if (validLocation && error) {
    console.log('error: ', `${latitude},${longitude}`, { error })
    loadingText =
      "Sorry sometimes we can't rely on positionstack: https://github.com/apilayer/positionstack/issues/7, please refresh the page again 🥲"
  }

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
          <div className="space-y-10">
            <div>
              <h3 className="text-lg font-medium sm:text-lg">Some current location information via the supergraph 📍:</h3>
              <p className="mt-4 text-sm text-gray-500">continent - {data?.location?.continent ? data?.location?.continent : loadingText} </p>
              <p className="mt-4 text-sm text-gray-500">county - {data?.location?.county ? data?.location?.county : loadingText} </p>
              <p className="mt-4 text-sm text-gray-500">label - {data?.location?.label ? data?.location?.label : loadingText} </p>
              <p className="mt-4 text-sm text-gray-500">latitude - {data?.location?.latitude ? data?.location?.latitude : loadingText} </p>
              <p className="mt-4 text-sm text-gray-500">longitude - {data?.location?.longitude ? data?.location?.longitude : loadingText} </p>
            </div>
            <h3 className="text-lg font-medium sm:text-lg">Some food in the area via the supergraph and @defer 🍟:</h3>
            <div className="mt-8 grid grid-cols-1 gap-8 md:mt-16 md:grid-cols-2 md:gap-12 lg:grid-cols-3">
              {data?.location?.eateriesForLocation
                ? data?.location?.eateriesForLocation.map(
                    (eatery: { id: string; name?: string; rating?: number; URL: string; reviewCount?: number; distance: number }) => (
                      <Card
                        key={eatery.id}
                        id={eatery.id}
                        name={eatery?.name}
                        rating={eatery?.rating}
                        URL={eatery.URL}
                        reviewCount={eatery?.reviewCount}
                        distance={eatery.distance}
                      />
                    ),
                  )
                : loadingText}
            </div>
          </div>
        </div>
      </section>
    </>
  )
}

export default Home
