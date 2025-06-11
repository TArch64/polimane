import { wait } from '@/helpers';
import type { HttpClient } from './HttpClient';

async function pingApi(http: HttpClient) {
  try {
    await http.get(['/ping']);
  } catch (error) {
    console.error(error);
  }
}

export async function useApiPing(http: HttpClient) {
  while (true) {
    await pingApi(http);
    await wait(60_000);
  }
}
