import { readdir } from 'node:fs/promises';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import type { InlineCollection } from 'unplugin-icons';
import icons from 'unplugin-icons/vite';
import { sentryVitePlugin } from '@sentry/vite-plugin';
import { VitePluginRadar } from 'vite-plugin-radar';
import { htmlReleaseCookie } from './vite/plugins';

const {
  SENTRY_AUTH_TOKEN,
  SENTRY_COMMIT_SHA,
  FRONTEND_RELEASE,
  FRONTEND_GOOGLE_ANALYTICS_ID,
} = process.env;

async function createCustomIconsCollection(): Promise<InlineCollection> {
  const files = await readdir('./src/assets/icons', { recursive: true });
  const iconFiles = files.filter((file) => file.endsWith('.svg'));

  const icons = iconFiles.map((file) => [
    file.replace('.svg', '').replaceAll('/', '-'),
    () => Bun.file(`./src/assets/icons/${file}`).text(),
  ]);

  return Object.fromEntries(icons);
}

export default defineConfig({
  clearScreen: false,
  envPrefix: 'FRONTEND_PUBLIC_',

  resolve: {
    alias: {
      '@': Bun.fileURLToPath(new URL('./src', import.meta.url)),
      '@editor': Bun.fileURLToPath(new URL('./src/modules/schemas/editor', import.meta.url)),
    },
  },

  css: {
    transformer: 'lightningcss',
  },

  build: {
    outDir: 'dist/public',
    sourcemap: 'hidden',
    cssMinify: 'lightningcss',
  },

  server: {
    host: '0.0.0.0',
    allowedHosts: true,
    origin: 'https://app.polimane.localhost',
    proxy: {}, // force http1

    https: {
      key: '/certs/polimane.localhost+1-key.pem',
      cert: '/certs/polimane.localhost+1.pem',
    },
  },

  plugins: [
    vue(),
    vueDevTools(),

    icons({
      compiler: 'vue3',

      customCollections: {
        custom: await createCustomIconsCollection(),
      },
    }),

    !!SENTRY_AUTH_TOKEN && sentryVitePlugin({
      org: 'myself-zmf',
      project: 'polimane-frontend',
      authToken: SENTRY_AUTH_TOKEN,
      telemetry: false,

      release: {
        name: FRONTEND_RELEASE,

        deploy: {
          env: 'production',
        },

        setCommits: {
          repo: 'TArch64/polimane',
          commit: SENTRY_COMMIT_SHA!,
        },
      },
    }),

    !!FRONTEND_GOOGLE_ANALYTICS_ID && VitePluginRadar({
      analytics: { id: FRONTEND_GOOGLE_ANALYTICS_ID },
    }),

    !!FRONTEND_RELEASE && htmlReleaseCookie(FRONTEND_RELEASE),
  ],
});
