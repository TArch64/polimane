import { onUnmounted } from 'vue';
import { wait } from '@/helpers';
import { authChannel, AuthChannelComplete } from '@/modules/auth/channel';
import { useAuthStore } from '../stores';

export interface IAuthPopupOptions {
  onSuccess: () => void;
}

export interface IAuthPopup {
  open: () => void;
}

export function useAuthPopup(options: IAuthPopupOptions): IAuthPopup {
  const store = useAuthStore();
  const abortController = new AbortController();
  let authWindow: Window | null = null;

  async function open(): Promise<void> {
    if (authWindow) {
      authWindow.focus();
      return;
    }

    authWindow = window.open(store.loginUrl, '_blank', 'popup=yes,width=600,height=600');

    if (!authWindow) {
      return;
    }

    cleanup();

    authChannel.addEventListener('message', (event) => {
      if (event.data === AuthChannelComplete) {
        options.onSuccess();
      }
    }, { signal: abortController.signal });
  }

  function isCompleteWindow(): boolean {
    try {
      return !!authWindow?.location.pathname;
    } catch {
      // ignore cross-origin errors
      return false;
    }
  }

  async function cleanup(): Promise<void> {
    while (true) {
      await wait(100);

      if (!authWindow || authWindow.closed) {
        authWindow = null;
        return;
      }
    }
  }

  onUnmounted(() => {
    abortController.abort();
    authChannel.close();
  });

  return { open };
}
