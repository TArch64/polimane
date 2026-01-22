import { type LocationQueryRaw, type RouteMap, type Router, useRouter } from 'vue-router';
import { markRaw, type Ref } from 'vue';
import { useAuthorized } from '../useAuthorized';
import type { IHttpResponseErrorInterceptor, IInterceptorContext } from './HttpMiddlewareExecutor';
import { HttpError } from './HttpError';
import { HttpErrorReason } from './HttpErrorReason';

export class HttpAuthorizationMiddleware implements IHttpResponseErrorInterceptor {
  static use(): IHttpResponseErrorInterceptor {
    return markRaw(new HttpAuthorizationMiddleware(useRouter(), useAuthorized()));
  }

  constructor(
    private readonly router: Router,
    private readonly authorized: Ref<boolean>,
  ) {
  }

  async interceptResponseError(error: unknown, ctx: IInterceptorContext): Promise<void> {
    if (HttpError.isError(error)) {
      switch (error.reason) {
        case HttpErrorReason.UNAUTHORIZED:
          this.authorized.value = false;
          if (ctx.meta.unauthorizedRedirect === false) return;
          return this.#redirect('auth');
        case HttpErrorReason.NOT_FOUND:
          return this.#redirect('home');
      }
    }
  }

  async #redirect<N extends keyof RouteMap>(name: N, query?: LocationQueryRaw): Promise<void> {
    await this.router.push({ name, query });
  }
}
