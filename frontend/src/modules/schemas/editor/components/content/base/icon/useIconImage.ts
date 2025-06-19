import { computed, type Ref, watch } from 'vue';
import { useObjectUrl } from '@vueuse/core';

export interface IIconImageOptions {
  source: Ref<string>;
}

export function useIconImage(options: IIconImageOptions): HTMLImageElement {
  const iconBlob = computed(() => {
    if (!options.source.value) return null;
    return new Blob([options.source.value], { type: 'image/svg+xml;charset=utf-8' });
  });

  const iconUrl = useObjectUrl(iconBlob);
  const iconImageEl = document.createElement('img');

  watch(iconUrl, (url) => {
    if (url) iconImageEl.src = url;
  });

  return iconImageEl;
}
