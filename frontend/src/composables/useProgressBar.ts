import { computed, type MaybeRefOrGetter, onUnmounted, toValue, watch } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import { done, start } from 'nprogress';
import { type AsyncAction, isAsyncAction } from '@/composables/useAsyncAction';

export function useProgressBar(isActiveRef: MaybeRefOrGetter<boolean> | AsyncAction) {
  const isActive = computed(() => {
    return isAsyncAction(isActiveRef) ? isActiveRef.isActive : toValue(isActiveRef);
  });

  const toggleBar = useDebounceFn((isActive: boolean) => {
    isActive ? start() : done();
  }, 10);

  watch(isActive, toggleBar);
  onUnmounted(() => toggleBar(false));
}
