import { useRouter } from 'vue-router';
import type { IHttpMiddleware } from './HttpMiddlewareExecutor';
import { HttpErrorReason } from './HttpErrorReason';

export function useUnauthorizedMiddleware(): IHttpMiddleware {
  const router = useRouter();

  return {
    async interceptResponseError(error) {
      if (error.reason === HttpErrorReason.UNAUTHORIZED) {
        await router.push({
          name: 'welcome',
          query: { 'return-to': router.currentRoute.value.path },
        });
      }
    },
  };
}
