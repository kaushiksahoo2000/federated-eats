import React from 'react'

const Card = () => {
  return (
    <a class="relative block rounded-xl border border-gray-100 p-8 shadow-xl" href="">
      <span class="absolute right-4 top-4 rounded-full bg-green-100 px-3 py-1.5 text-xs font-medium text-green-600">4.3</span>

      <div class="mt-4 text-gray-500 sm:pr-8">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 15.546c-.523 0-1.046.151-1.5.454a2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.701 2.701 0 00-1.5-.454M9 6v2m3-2v2m3-2v2M9 3h.01M12 3h.01M15 3h.01M21 21v-7a2 2 0 00-2-2H5a2 2 0 00-2 2v7h18zm-3-9v-2a2 2 0 00-2-2H8a2 2 0 00-2 2v2h12z"
          ></path>
        </svg>

        <h3 class="mt-4 text-xl font-medium text-gray-900">Science of Chemistry</h3>

        <p class="mt-2 hidden text-sm sm:block">Yelp Business ID: x</p>
      </div>
    </a>
  )
}

export default Card
