import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import icons from 'unplugin-icons/vite';
import { sentryVitePlugin } from '@sentry/vite-plugin';

const { SENTRY_AUTH_TOKEN, FRONTEND_PUBLIC_SENTRY_RELEASE } = process.env;

export default defineConfig({
  clearScreen: false,
  envPrefix: 'FRONTEND_PUBLIC_',

  resolve: {
    alias: {
      '@': Bun.fileURLToPath(new URL('./src', import.meta.url)),
    },
  },

  css: {
    transformer: 'lightningcss',
  },

  build: {
    cssMinify: 'lightningcss',
    sourcemap: true,
  },

  server: {
    host: '0.0.0.0',
    allowedHosts: true,
  },

  plugins: [
    vue(),
    vueDevTools(),

    icons({
      compiler: 'vue3',

      customCollections: {
        custom: {
          logo: () => Bun.file('./src/assets/logo.svg').text(),
        },
      },
    }),

    !!SENTRY_AUTH_TOKEN && sentryVitePlugin({
      org: 'myself-zmf',
      project: 'polimane-frontend',
      authToken: SENTRY_AUTH_TOKEN,
      telemetry: false,
      release: { name: FRONTEND_PUBLIC_SENTRY_RELEASE },
    }),
  ],
});
