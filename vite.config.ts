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

  build: {
    cssMinify: 'lightningcss',
  },

  css: {
    transformer: 'lightningcss',

    postcss: {
      plugins: [
        postcssNesting(),
      ],
    },
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
