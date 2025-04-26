import { defineConfig, ModuleNode } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import postcssNesting from 'postcss-nesting';
import icons from 'unplugin-icons/vite';

const RELOAD_ON_PATHS = [
  'src/modules/schemas/editor/components/content',
  'src/modules/schemas/editor/composables/content',
];

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

    {
      name: 'full-reload',

      handleHotUpdate({ file, server, modules, timestamp }) {
        const isReloadable = RELOAD_ON_PATHS.some((path) => {
          return file.startsWith(`${import.meta.dirname}/${path}`);
        });

        if (isReloadable) {
          const invalidatedModules = new Set<ModuleNode>();

          for (const mod of modules) {
            server.moduleGraph.invalidateModule(mod, invalidatedModules, timestamp, true);
          }

          server.ws.send({ type: 'full-reload' });
          return [];
        }
      },
    },
  ],
});
