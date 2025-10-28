import type { FunctionDirective } from 'vue';

export const vVisible: FunctionDirective<HTMLElement, boolean> = (el, binding) => {
  el.style.visibility = binding.value ? 'visible' : 'hidden';
};
