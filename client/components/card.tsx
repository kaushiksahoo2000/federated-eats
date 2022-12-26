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
            d="M21 15.546c-.523 0-1.046.151-1.5.454a2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.701 2.701 0 00-1.5-.454M9 6v2m3-2v2m3-2v2M9 3h.01M12 3h.01M15 3h.01M21 21v-7a2 2 0 00-2-2H5a2 2 0 00-2 2v7h18zm-3-9v-2a2 2 0 00-2-2H8a2 2 0 00-2 2v2h12z"
          ></path>
        </svg>
        <h3 className="mt-4 text-sm font-medium text-gray-900">{name}</h3>
        <p className="mt-2 hidden text-sm sm:block">Yelp Business ID: {id}</p>
        <p className="mt-2 hidden text-sm sm:block">
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
