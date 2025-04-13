/// <reference types="vite/client" />
/// <reference types="unplugin-icons/types/vue" />

interface ImportMetaEnv {
  FRONTEND_PUBLIC_API_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare global {
  // This is a workaround for the fact that `setTimeout` returns a different type in Node.js and browser environments.
  type TimeoutId = ReturnType<typeof setTimeout>;
}

export {};
