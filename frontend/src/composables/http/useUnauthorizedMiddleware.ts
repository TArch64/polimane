import type { RouteMap } from 'vue-router';
import { type LocationQueryRaw, useRouter } from 'vue-router';
import { useAuthToken } from '@/composables';
import type { HttpMiddleware } from './HttpMiddlewareExecutor';
import { HttpErrorReason } from './HttpErrorReason';

export function useUnauthorizedMiddleware(): HttpMiddleware {
  const router = useRouter();
  const authToken = useAuthToken();

  async function redirect<N extends keyof RouteMap>(name: N, query?: LocationQueryRaw): Promise<void> {
    await router.push({ name, query });
  }

  async function handleUnauthorized(): Promise<void> {
    authToken.value = undefined;

    return redirect('welcome', {
      'return-to': router.currentRoute.value.fullPath,
    });
  }

  return {
    async interceptResponseError(error): Promise<void> {
      switch (error.reason) {
        case HttpErrorReason.UNAUTHORIZED:
          return handleUnauthorized();
        case HttpErrorReason.NOT_FOUND:
          return redirect(authToken.value ? 'home' : 'welcome');
      }
    },
  };
}
