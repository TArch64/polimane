import type { RemovableRef } from '@vueuse/core';
import { useAccessToken } from '@/composables';
import type { HttpClient } from './HttpClient';
import type { HttpMiddleware, IHttpBeforeRequestInterceptor } from './HttpMiddlewareExecutor';

export class HttpApiPing implements IHttpBeforeRequestInterceptor {
  static use(http: HttpClient): HttpMiddleware {
    return new HttpApiPing(http, useAccessToken());
  }

  private timeoutId: TimeoutId | null = null;

  constructor(
    private readonly http: HttpClient,
    private readonly token: RemovableRef<string | undefined>,
  ) {
    this.pingApi = this.pingApi.bind(this);
    requestIdleCallback(this.pingApi);
  }

  interceptBeforeRequest(): void {
    if (this.timeoutId) clearTimeout(this.timeoutId);
    this.timeoutId = setTimeout(this.pingApi, 60_000);
  }

  private async pingApi(): Promise<void> {
    if (!this.token.value) return;

    try {
      await this.http.get('/ping');
    } catch (error) {
      console.error(error);
    }
  }
}
