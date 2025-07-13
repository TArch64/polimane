export function buildServerUrl(path: string, query: Record<string, string>): URL {
  const url = new URL(import.meta.env.FRONTEND_PUBLIC_API_URL);
  url.pathname += path;
  url.search = new URLSearchParams(query).toString();
  return url;
}
