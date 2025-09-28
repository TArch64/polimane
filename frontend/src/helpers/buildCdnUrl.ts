const cdnBaseUrl = `https://${import.meta.env.FRONTEND_PUBLIC_CDN_HOST}`;

export function buildCdnUrl(path: string | null): string | null {
  return path ? `${cdnBaseUrl}/${path}` : null;
}
