import type { HttpMiddleware } from './HttpMiddlewareExecutor';
import { useAuthToken } from './useAuthToken';

export function useAuthorizationMiddleware(): HttpMiddleware {
  const token = useAuthToken();

  return {
    interceptBeforeRequest(request: Request): void {
      if (token.value) request.headers.set('Authorization', token.value);
    },
  };
}
