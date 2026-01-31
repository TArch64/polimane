/// <reference types="vite/client" />
/// <reference types="unplugin-icons/types/vue" />

import type { MaybePromise } from '@/types';

interface ImportMetaEnv {
  FRONTEND_PUBLIC_API_URL: string;
  FRONTEND_PUBLIC_SENTRY_DSN: string;
  FRONTEND_PUBLIC_CDN_HOST: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare global {
  // This is a workaround for the fact that `setTimeout` returns a different type in Node.js and browser environments.
  type TimeoutId = ReturnType<typeof setTimeout>;
}

declare module 'vue' {
  interface ComponentCustomOptions {
    beforeModalOpen?(props?: object): MaybePromise<void>;
  }
}

export {};
