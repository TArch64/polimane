export type UrlParams = Record<string, string | number>;
export type UrlPathItem = string | number;
export type UrlPath = UrlPathItem[] | UrlPathItem;

export function buildUrl(base: string, path_: UrlPath, params: UrlParams = {}): URL {
  const urlStr = [base, '', path_]
    .flat()
    .join('/')
    .replaceAll(/\/+/g, '/');

  const url = new URL(urlStr);

  if (Object.keys(params).length) {
    url.search = new URLSearchParams(params as Record<string, string>).toString();
  }

  return url;
}
