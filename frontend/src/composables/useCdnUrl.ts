import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import { buildUrl, type UrlParams, type UrlPath } from '@/helpers';

export interface ICdnUrlOptions {
  path: UrlPath;
  params?: UrlParams;
}

const cdnBaseUrl = `https://${import.meta.env.FRONTEND_PUBLIC_CDN_HOST}`;

export function useCdnUrl(options_: MaybeRefOrGetter<ICdnUrlOptions | null>): ComputedRef<URL | null> {
  return computed(() => {
    const options = toValue(options_);
    if (!options) return null;
    return buildUrl(cdnBaseUrl, options.path, options.params);
  });
}
