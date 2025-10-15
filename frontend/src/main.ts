import './style/main.css';

import './polyfills';
import { createApp, h } from 'vue';
import { createPinia } from 'pinia';
import { configure as configureProgress } from 'nprogress';
import { SentryPlugin } from '@/plugins';
import { ContextMenuPlugin } from './components/contextMenu';
import { ConfirmPlugin } from './components/confirm';
import { ModalPlugin } from './components/modal';
import { httpClientPlugin } from './composables';
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

app.mount('#app');
