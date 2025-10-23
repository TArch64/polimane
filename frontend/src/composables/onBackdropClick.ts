import { onMounted, onUnmounted } from 'vue';

export function onBackdropClick(onClick: (event: MouseEvent) => void) {
  const abortController = new AbortController();

  onMounted(() => requestAnimationFrame(() => {
    window.addEventListener('mousedown', (event: MouseEvent): void => {
      const elements = document.elementsFromPoint(event.clientX, event.clientY);

      for (const element of elements) {
        if (!(element instanceof HTMLDialogElement)) continue;

        const rect = element.getBoundingClientRect();

        if (
          rect.top <= event.clientY
          && event.clientY <= rect.top + rect.height
          && rect.left <= event.clientX
          && event.clientX <= rect.left + rect.width
        ) return;
      }

      onClick(event);
    }, { signal: abortController.signal });
  }));

  onUnmounted(() => abortController.abort());
}
