import type { NextPage } from 'next'
import Card from './card'
import Link from 'next/link'
import Table from './table'

const Home: NextPage = () => {
  const dummydata = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
  return (
    <>
      <section className="text-gray-700">
        <div className="px-4 py-16 sm:px-6 lg:px-8 space-y-10">
          <div className="max-w-xl space-y-10">
            <h2 className="text-3xl font-normal sm:text-4xl">Federated Eats</h2>
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

          <div>
            <h3 className="text-lg font-normal sm:text-lg">Some current location information via the supergraph 📍:</h3>
            <p className="mt-4">x1 - </p>
            <p className="mt-4">x2 - </p>
            <p className="mt-4">x2 - </p>
          </div>
          <h3 className="text-lg font-normal sm:text-lg">Some food in the area via the supergraph 🍟:</h3>
          <div className="mt-8 grid grid-cols-1 gap-8 md:mt-16 md:grid-cols-2 md:gap-12 lg:grid-cols-3">
            {dummydata.map((x) => (
              <Card />
            ))}
          </div>
        </div>
      </section>
    </>
  )
}

export default Home
