import './style/main.css';

import './polyfills';
import { createApp, h, shallowRef } from 'vue';
import { createPinia } from 'pinia';
import VueKonva from 'vue-konva';
import { configure as configureProgress } from 'nprogress';
import { sentryPlugin } from '@/plugins';
import { ContextMenuPlugin } from './components/contextMenu';
import { ConfirmPlugin } from './components/confirm';
import { ModalPlugin } from './components/modal';
import { httpClientPlugin } from './composables';
import App from './App.vue';
import { router } from './router';

configureProgress({
  showSpinner: false,
});

window.__KONVA_STAGE_REF__ = shallowRef(null);

const app = createApp({
  render: () => h(App),
});

app.use(createPinia());
app.use(router);

app.use(httpClientPlugin, {
  baseUrl: import.meta.env.FRONTEND_PUBLIC_API_URL,
});

app.use(ModalPlugin);
app.use(ConfirmPlugin);
app.use(ContextMenuPlugin);

app.use(VueKonva, { prefix: 'Konva' });

app.use(sentryPlugin, {
  dsn: import.meta.env.FRONTEND_PUBLIC_SENTRY_DSN || '',
});

if (import.meta.env.DEV) {
  const m = await import('./modules/konvaInspector');
  app.use(m.konvaInspectorPlugin);
}

app.mount('#app');
