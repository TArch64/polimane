import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import postcssNesting from 'postcss-nesting';

export default defineConfig({
  clearScreen: false,

  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
  ],

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },

  build: {
    cssMinify: 'lightningcss'
  },

  css: {
    transformer: 'lightningcss',

    postcss: {
      plugins: [
        postcssNesting()
      ]
    }
  }
})
