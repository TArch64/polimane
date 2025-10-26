import { type FunctionDirective, nextTick } from 'vue';
import { type IPadding, type PaddingInput, resolvePadding } from '@/helpers';
import { NodeRect } from '@/models';

type Modifiers = 'defer';

interface IPopoverShiftProps {
  padding?: PaddingInput;
}

function getOffsetX(rect: NodeRect, padding: IPadding): number {
  if (rect.left < padding.left) {
    return padding.left - Math.max(rect.left, 0);
  }

  const offset = window.innerWidth - rect.right - padding.right;
  return offset > 0 ? 0 : offset;
}

function getOffsetY(rect: NodeRect, padding: IPadding): number {
  if (rect.top < padding.top) {
    return padding.top - Math.max(rect.top, 0);
  }

  const offset = window.innerHeight - rect.bottom - padding.bottom;
  return offset > 0 ? 0 : offset;
}

export const vPopoverShift: FunctionDirective<HTMLElement, IPopoverShiftProps, Modifiers> = (el, binding) => {
  function render() {
    const padding = resolvePadding(binding.value?.padding ?? 8);
    const rect = new NodeRect(el.getBoundingClientRect());

    const offsetX = getOffsetX(rect, padding);
    const offsetY = getOffsetY(rect, padding);

    el.style.translate = `${offsetX}px ${offsetY}px`;
  }

  binding.modifiers.defer ? nextTick(render) : render();
};
