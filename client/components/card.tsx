import React from 'react'

const Card = ({ name, rating, id }: { name?: string; rating?: number; id?: string }) => {
  console.log(name, rating, id)
  return (
    <a className="relative block rounded-xl border border-gray-100 p-8 shadow-xl" href="">
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
      </div>
    </a>
  )
}

export default Card
