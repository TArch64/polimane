import Konva from 'konva';
import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';

export type MaybeNodeConfig<C extends Konva.NodeConfig> = Partial<C> | null | undefined | boolean;

export function flattenNodeConfigs<C extends Konva.NodeConfig>(configRefs: MaybeRefOrGetter<MaybeNodeConfig<C>>[]): Partial<C> {
  const result: Partial<C> = {};

  for (const configRef of configRefs) {
    const config = toValue(configRef);

    if (config) {
      Object.assign(result, config);
    }
  }

  return result;
}

export function useNodeConfigs<C extends Konva.NodeConfig>(configRefs: MaybeRefOrGetter<MaybeNodeConfig<C>>[]): ComputedRef<Partial<C>> {
  return computed(() => flattenNodeConfigs(configRefs));
}
