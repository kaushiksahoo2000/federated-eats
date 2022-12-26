import React from 'react'
import Link from 'next/link'

const Card = ({
  name,
  rating,
  id,
  URL,
  reviewCount,
  distance,
}: {
  name?: string
  rating?: number
  id: string
  URL: string
  reviewCount?: number
  distance: number
}) => {
  // console.log({ name, rating, id, URL, reviewCount, distance })
  return (
    <div className="relative block rounded-xl border border-gray-100 p-8 shadow-xl">
      <span className="absolute right-4 top-4 rounded-full bg-green-100 px-3 py-1.5 text-xs font-medium text-green-600">{rating}</span>

      <div className="mt-4 text-gray-500 sm:pr-8">
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
          ></path>
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
        </svg>
        <h3 className="mt-4 text-sm font-medium text-gray-900">{name}</h3>
        <p className="mt-2 text-sm sm:block">Yelp Business ID: {id}</p>
        <p className="mt-2 text-sm sm:block">
          <Link className="text-blue-600" target="_blank" href={URL}>
            Yelp Business Page
          </Link>
        </p>
        <div className="mt-4 flex flex-wrap gap-1">
          <span className="whitespace-nowrap rounded-full bg-blue-100 px-2.5 py-0.5 text-xs text-blue-600">{reviewCount} reviews</span>
          <span className="whitespace-nowrap rounded-full bg-blue-100 px-2.5 py-0.5 text-xs text-blue-600">{Math.round(distance)} meters away</span>
        </div>
      </div>
    </div>
  )
}

export default Card
