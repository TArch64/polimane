import { buildServerUrl } from './buildServerUrl';

export async function openLoginUrl(returnTo: string): Promise<void> {
  const url = buildServerUrl('/auth/login', {
    'return-to': returnTo,
  });

  const response = await fetch(url, { redirect: 'manual' });
  window.location.assign(response.url);
}
