import { onMounted, onUnmounted, type Ref } from 'vue';

export function onBackdropClick(el: Ref<HTMLDialogElement>, onClick: (event: MouseEvent) => void) {
  const abortController = new AbortController();

  onMounted(() => requestAnimationFrame(() => {
    window.addEventListener('mousedown', (event: MouseEvent): void => {
      const rect = el.value.getBoundingClientRect();

      if (!(
        rect.top <= event.clientY
        && event.clientY <= rect.top + rect.height
        && rect.left <= event.clientX
        && event.clientX <= rect.left + rect.width
      )) {
        onClick(event);
      }
    }, { signal: abortController.signal });
  }));

  onUnmounted(() => abortController.abort());
}
