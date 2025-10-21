import { type LocationQueryRaw, type RouteMap, type Router, useRouter } from 'vue-router';
import type { Ref } from 'vue';
import { useAuthorized } from '../useAuthorized';
import type { HttpMiddleware, IHttpResponseErrorInterceptor } from './HttpMiddlewareExecutor';
import { HttpError } from './HttpError';
import { HttpErrorReason } from './HttpErrorReason';

export class HttpAuthorization implements IHttpResponseErrorInterceptor {
  static use(): HttpMiddleware {
    return new HttpAuthorization(useRouter(), useAuthorized());
  }

  constructor(
    private readonly router: Router,
    private readonly authorized: Ref<boolean>,
  ) {
  }

  async interceptResponseError(error: HttpError): Promise<void> {
    switch (error.reason) {
      case HttpErrorReason.UNAUTHORIZED:
        this.authorized.value = false;
        if (error.meta.unauthorizedRedirect === false) return;
        return this.redirect('auth');
      case HttpErrorReason.NOT_FOUND:
        return this.redirect('home');
    }
  }

  private async redirect<N extends keyof RouteMap>(name: N, query?: LocationQueryRaw): Promise<void> {
    await this.router.push({ name, query });
  }
}
