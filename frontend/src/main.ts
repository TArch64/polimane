import './style/main.css';

import { createApp, h, ref, shallowRef } from 'vue';
import { createPinia } from 'pinia';
import VueKonva from 'vue-konva';
import { TOKEN_SCROLLER, TOKEN_TOP_EL } from './InjectionToken';
import { ConfirmPlugin } from './components/confirm';
import { ModalPlugin } from './components/modal';
import { httpClientPlugin } from './composables';
import App from './App.vue';
import { router } from './router';

window.__KONVA_STAGE_REF__ = shallowRef(null);

const app = createApp({
  render: () => h(App),

  provide: {
    [TOKEN_SCROLLER]: ref<HTMLElement>(document.scrollingElement as HTMLElement),
    [TOKEN_TOP_EL]: ref<HTMLElement>(document.body),
  },
});

app.use(createPinia());
app.use(router);

app.use(httpClientPlugin, {
  baseUrl: import.meta.env.FRONTEND_PUBLIC_API_URL,
});

app.use(ModalPlugin);
app.use(ConfirmPlugin);

app.use(VueKonva, { prefix: 'Konva' });

if (import.meta.env.DEV) {
  const m = await import('./modules/konvaInspector');
  app.use(m.konvaInspectorPlugin);
}

document.startViewTransition(async () => {
  app.mount('#app');
  return router.isReady();
});
