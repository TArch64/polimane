import type { Ref } from 'vue';
import { requestIdleCallback } from '@/helpers';
import { useAuthorized } from '../useAuthorized';
import type { HttpClient } from './HttpClient';
import type { IHttpBeforeRequestInterceptor } from './HttpMiddlewareExecutor';

export class HttpApiPingMiddleware implements IHttpBeforeRequestInterceptor {
  static use(http: HttpClient): IHttpBeforeRequestInterceptor {
    return new HttpApiPingMiddleware(http, useAuthorized());
  }

  private timeoutId: TimeoutId | null = null;

  constructor(
    private readonly http: HttpClient,
    private readonly authorized: Ref<boolean>,
  ) {
    this.pingApi = this.pingApi.bind(this);
    requestIdleCallback(this.pingApi);
  }

  interceptBeforeRequest(): void {
    if (this.timeoutId) clearTimeout(this.timeoutId);
    this.timeoutId = setTimeout(this.pingApi, 60_000);
  }

  private async pingApi(): Promise<void> {
    if (!this.authorized.value) return;

    try {
      await this.http.get('/ping', {}, {
        meta: { unauthorizedRedirect: false },
      });
    } catch (error) {
      console.error(error);
    }
  }
}
