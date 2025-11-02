import type { MaybeArray } from '@/types';
import { getObjectEntries } from './object';

export type UrlParams = Record<string, MaybeArray<string | number>>;
export type UrlPathItem = string | number;
export type UrlPath = UrlPathItem[] | UrlPathItem;

function serializeParams(raw: UrlParams): URLSearchParams {
  const params = new URLSearchParams();

  for (const [key, value] of getObjectEntries(raw)) {
    if (Array.isArray(value)) {
      for (const v of value) {
        params.append(key, String(v));
      }
    } else {
      params.append(key, String(value));
    }
  }

  return params;
}

export function buildUrl(base: string, path_: UrlPath, params: UrlParams = {}): URL {
  const urlStr = [base, '', path_]
    .flat()
    .join('/')
    .replaceAll(/\/+/g, '/');

  const url = new URL(urlStr);

  if (Object.keys(params).length) {
    url.search = serializeParams(params).toString();
  }

  return url;
}
