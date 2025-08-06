import { computed, type ComputedRef, type MaybeRefOrGetter, ref, type Ref } from 'vue';
import Konva from 'konva';
import { useNodeListener } from './useNodeListener';

export function useNodeHovering(nodeRef: MaybeRefOrGetter<Konva.Node>, onChange?: (isHovering: boolean) => void): Ref<boolean> {
  const isHovering = ref(false);

  function setHovering(hovering: boolean) {
    isHovering.value = hovering;
    onChange?.(hovering);
  }

  useNodeListener(nodeRef, 'mouseover', () => setHovering(true));
  useNodeListener(nodeRef, 'mouseout', () => setHovering(false));

  return isHovering;
}

export function useNodeHoveringConfig<C extends Konva.NodeConfig>(nodeRef: MaybeRefOrGetter<Konva.Node>, config: Partial<C>): ComputedRef<Partial<C>> {
  const isHovering = useNodeHovering(nodeRef);
  return computed(() => isHovering.value ? config : {});
}
