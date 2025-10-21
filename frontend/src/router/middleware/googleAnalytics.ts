import type { NavigationHookAfter, RouteLocationNormalized } from 'vue-router';

export const googleAnalytics: NavigationHookAfter = (to) => {
  const clientId = import.meta.env.FRONTEND_PUBLIC_GOOGLE_ANALYTICS_ID;
  if (!clientId || !gtag) return;

  gtag('config', clientId, {
    page_path: to.path,
    page_title: to.name || (to as RouteLocationNormalized).path,
  });
};
