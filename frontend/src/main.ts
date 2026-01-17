import './style/main.css';
import { createApp, h } from 'vue';
import { createPinia } from 'pinia';
import { configure as configureProgress } from 'nprogress';
import { SentryPlugin } from '@/plugins';
import { useSessionStore } from '@/stores';
import { ContextMenuPlugin } from './components/contextMenu';
import { ConfirmPlugin } from './components/confirm';
import { ModalPlugin } from './components/modal';
import { httpClientPlugin, UpdateUserCountersMiddleware, useHttpClient } from './composables';
import App from './App.vue';
import { router } from './router';

configureProgress({
  showSpinner: false,
});

const app = createApp({
  render: () => h(App),
});

app.use(createPinia());
app.use(router);

app.use(httpClientPlugin, {
  baseUrl: import.meta.env.FRONTEND_PUBLIC_API_URL!,
});

app.use(ModalPlugin);
app.use(ConfirmPlugin);
app.use(ContextMenuPlugin);

app.use(SentryPlugin, {
  dsn: import.meta.env.FRONTEND_PUBLIC_SENTRY_DSN || '',
});

app.runWithContext(() => {
  const httpClient = useHttpClient();
  const middleware = httpClient.getMiddleware(UpdateUserCountersMiddleware);
  middleware?.setSessionStore(useSessionStore());
});

app.mount('#app');
