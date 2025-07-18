import type { RemovableRef } from '@vueuse/core';
import { type LocationQueryRaw, type RouteMap, type Router, useRouter } from 'vue-router';
import type { MaybePromise } from '@/types';
import type {
  HttpMiddleware,
  IHttpBeforeRequestInterceptor,
  IHttpResponseErrorInterceptor,
  IHttpResponseSuccessInterceptor,
} from './HttpMiddlewareExecutor';
import { useAccessToken, useRefreshAccessToken } from './useAccessToken';
import { HttpError } from './HttpError';
import { HttpErrorReason } from './HttpErrorReason';

export class HttpAuthorization implements IHttpBeforeRequestInterceptor,
  IHttpResponseErrorInterceptor,
  IHttpResponseSuccessInterceptor {
  static use(): HttpMiddleware {
    return new HttpAuthorization(
      useRouter(),
      useAccessToken(),
      useRefreshAccessToken(),
    );
  }

  constructor(
    private readonly router: Router,
    private readonly accessToken: RemovableRef<string | undefined>,
    private readonly refreshToken: RemovableRef<string | undefined>,
  ) {
  }

  interceptBeforeRequest(request: Request): void {
    if (this.accessToken.value) {
      request.headers.set('Authorization', this.accessToken.value);
    }

    if (this.refreshToken.value) {
      request.headers.set('X-Refresh-Token', this.refreshToken.value);
    }
  }

  interceptResponseSuccess(response: Response): MaybePromise<void> {
    const refreshToken = response.headers.get('X-New-Refresh-Token');
    if (refreshToken) this.refreshToken.value = refreshToken;

    const accessToken = response.headers.get('X-New-Access-Token');
    if (accessToken) this.accessToken.value = accessToken;
  }

  async interceptResponseError(error: HttpError): Promise<void> {
    switch (error.reason) {
      case HttpErrorReason.UNAUTHORIZED:
        this.accessToken.value = undefined;
        this.refreshToken.value = undefined;
        return this.redirect('auth');
      case HttpErrorReason.NOT_FOUND:
        return this.redirect('home');
    }
  }

  private async redirect<N extends keyof RouteMap>(name: N, query?: LocationQueryRaw): Promise<void> {
    await this.router.push({ name, query });
  }
}
