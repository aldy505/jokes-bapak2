import { $fetch } from 'ohmyfetch';
import env from '../lib/env';

interface TotalResponse {
  message: number;
}

/** @type {import('./$types').PageServerLoad} */
export async function load() {
  const response = await $fetch<TotalResponse>('total', {
    method: 'GET',
    baseURL: env.SERVER_API_ENDPOINT,
    parseResponse: JSON.parse,
  });

  return {
    total: response.message,
  };
}
