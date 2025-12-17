import { onUnmounted } from 'vue';
import { wait } from '@/helpers';
import { authChannel, AuthChannelEvent } from '@/modules/auth/channel';
import { useAuthStore } from '../stores';

export interface IAuthPopupOptions {
  onSuccess: () => void;
  onDeletedUser: () => void;
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
      switch (event.data.type) {
        case AuthChannelEvent.COMPLETE:
          return options.onSuccess();
        case AuthChannelEvent.DELETED_USER:
          return options.onDeletedUser();
      }
    }, { signal: abortController.signal });
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

  onUnmounted(() => abortController.abort());

  return { open };
}
