import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import postcssNesting from 'postcss-nesting';
import icons from 'unplugin-icons/vite';

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

    postcss: {
      plugins: [
        postcssNesting(),
      ],
    },
  },

  build: {
    cssMinify: 'lightningcss',
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
    }),
  ],
});
