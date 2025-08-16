import Konva from 'konva';
import { type MaybeRefOrGetter, ref, type Ref } from 'vue';
import { useNodeListener } from './useNodeListener';

export function useNodeHover(
  nodeRef: MaybeRefOrGetter<Konva.Node>,
  onChange?: (isHovered: boolean) => void,
): Readonly<Ref<boolean>> {
  const isHovered = ref(false);

  function update(toHovered: boolean): void {
    isHovered.value = toHovered;
    onChange?.(toHovered);
  }

  useNodeListener(nodeRef, 'mouseover', () => update(true));
  useNodeListener(nodeRef, 'mouseout', () => update(false));

  return isHovered;
}
