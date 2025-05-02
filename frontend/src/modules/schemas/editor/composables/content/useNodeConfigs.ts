import Konva from 'konva';
import { computed, type MaybeRefOrGetter, type Ref, ref, toValue, watch } from 'vue';

export type MaybeNodeConfig<C extends Konva.NodeConfig> = Partial<C> | null | undefined | boolean;

export function flattenNodeConfigs<C extends Konva.NodeConfig>(configs: MaybeNodeConfig<C>[]): Partial<C> {
  const result: Partial<C> = {};

  for (const config of configs) {
    if (config) {
      Object.assign(result, config);
    }
  }

  return result;
}

export function useNodeConfigs<C extends Konva.NodeConfig>(configRefs: MaybeRefOrGetter<MaybeNodeConfig<C>>[]): Ref<Partial<C>> {
  const refs = configRefs.map((configRef) => computed(() => toValue(configRef)));
  const flatten: Ref<Partial<C>> = ref({});
  watch(refs, (values) => flatten.value = flattenNodeConfigs(values), { immediate: true });
  return flatten;
}
