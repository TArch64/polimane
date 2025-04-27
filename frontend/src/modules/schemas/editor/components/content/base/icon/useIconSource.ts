import { computed, reactive, type Ref, ref } from 'vue';
import { useDebounceFn, useMutationObserver } from '@vueuse/core';

export interface IIconSourceOptions {
  size: Ref<number | string>;
  color: Ref<string>;
}

export interface IIconSource {
  hostEl: HTMLElement;
  source: string;
}

export function useIconSource(options: IIconSourceOptions): IIconSource {
  const hostEl = document.createElement('div');
  const originalSource = ref('');

  const onMutation = useDebounceFn(() => {
    const svgEl = hostEl.querySelector('svg');

    if (!svgEl) {
      return;
    }

    svgEl.setAttribute('width', String(options.size.value));
    svgEl.setAttribute('height', String(options.size.value));
    originalSource.value = svgEl.outerHTML;

    if (originalSource.value && !originalSource.value.includes('xmlns="http://www.w3.org/2000/svg"')) {
      originalSource.value = originalSource.value.replace('<svg', '<svg xmlns="http://www.w3.org/2000/svg"');
    }
  }, 0);

  useMutationObserver(hostEl, onMutation, {
    subtree: true,
    childList: true,
  });

  const iconSource = computed(() => {
    return options.color.value
      ? originalSource.value.replaceAll('currentColor', options.color.value)
      : originalSource.value;
  });

  return reactive({
    hostEl,
    source: iconSource,
  });
}
