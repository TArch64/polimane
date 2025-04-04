import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import vueDevTools from 'vite-plugin-vue-devtools';
import postcssNesting from 'postcss-nesting';
import icons from 'unplugin-icons/vite';

export default defineConfig({
  clearScreen: false,

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
    vueJsx(),
    vueDevTools(),
    icons({
      compiler: 'vue3',
    }),
  ],
});
