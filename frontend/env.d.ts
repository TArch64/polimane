/// <reference types="vite/client" />
/// <reference types="unplugin-icons/types/vue" />

import Konva from 'konva';
import { ShallowRef } from 'vue';

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

  interface Window {
    __KONVA_STAGE_REF__: ShallowRef<Konva.Stage | null>;
  }
}

declare module 'konva/lib/Node' {
  export interface NodeEventMap extends GlobalEventHandlersEventMap {
  }
}

export {};
