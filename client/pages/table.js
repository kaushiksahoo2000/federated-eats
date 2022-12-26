import React from 'react'
import Link from 'next/link'

const Table = () => {
  return (
    <div className="max-w-xl">
      <table className="min-w-full divide-y divide-gray-200 text-sm">
        <thead className="bg-gray-100">
          <tr>
            <th className="whitespace-nowrap px-4 py-2 text-left font-medium text-gray-900">Graph</th>
            <th className="whitespace-nowrap px-4 py-2 text-left font-medium text-gray-900">Link</th>
          </tr>
        </thead>

        <tbody className="divide-y divide-gray-200">
          <tr>
            <td className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Locations Subgraph</td>
            <td className="whitespace-nowrap px-4 py-2 text-gray-700">
              <Link className="text-blue-500" target="_blank" href={'https://subgraph-locations.up.railway.app/sandbox'}>
                https://subgraph-locations.up.railway.app/sandbox
              </Link>
            </td>
          </tr>

          <tr>
            <td className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Eateries Subgraph</td>
            <td className="whitespace-nowrap px-4 py-2 text-gray-700">
              <Link className="text-blue-500" target="_blank" href={'https://subgraph-eateries.up.railway.app/sandbox'}>
                https://subgraph-eateries.up.railway.app/sandbox
              </Link>
            </td>
          </tr>

          <tr>
            <td className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">Federated Eats Supergraph</td>
            <td className="whitespace-nowrap px-4 py-2 text-gray-700">
              <Link className="text-blue-500" target="_blank" href={'https://subgraph-eateries.up.railway.app/sandbox'}>
                https://subgraph-eateries.up.railway.app/sandbox
              </Link>
              {', '}
              <Link className="text-blue-500" target="_blank" href={'https://subgraph-eateries.up.railway.app/sandbox'}>
                https://subgraph-eateries.up.railway.app/sandbox
              </Link>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  )
}

export default Table
