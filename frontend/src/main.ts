import './style/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import { pluginHttpClient } from '@/composables';
import App from './App.vue';
import { router } from './router';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(pluginHttpClient, {
  baseUrl: import.meta.env.FRONTEND_PUBLIC_API_URL,
});

document.startViewTransition(async () => {
  app.mount('#app');
  return router.isReady();
});
