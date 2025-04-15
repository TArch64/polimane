import './style/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { httpClientPlugin } from './composables';
import { ModalPlugin } from './components/modal';
import App from './App.vue';
import { router } from './router';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(httpClientPlugin, {
  baseUrl: import.meta.env.FRONTEND_PUBLIC_API_URL,
});

app.use(ModalPlugin);

document.startViewTransition(async () => {
  app.mount('#app');
  return router.isReady();
});
