import { useRouter } from 'vue-router';
import { useAuthToken } from '@/composables';
import type { HttpMiddleware } from './HttpMiddlewareExecutor';
import { HttpErrorReason } from './HttpErrorReason';

export function useUnauthorizedMiddleware(): HttpMiddleware {
  const router = useRouter();
  const authToken = useAuthToken();

  return {
    async interceptResponseError(error): Promise<void> {
      if (error.reason === HttpErrorReason.UNAUTHORIZED) {
        authToken.value = undefined;

        await router.push({
          name: 'welcome',
          query: { 'return-to': router.currentRoute.value.fullPath },
        });
      }
    },
  };
}
