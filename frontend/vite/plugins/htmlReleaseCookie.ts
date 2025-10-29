import type { Plugin } from 'vite';

export const htmlReleaseCookie = (key: string): Plugin => ({
  name: 'htmlReleaseCookie',

  transformIndexHtml: () => [
    {
      tag: 'script',
      injectTo: 'head-prepend',
      children: `document.cookie="webapp-release=${key};path=/;max-age=31536000";`,
    },
  ],
});
