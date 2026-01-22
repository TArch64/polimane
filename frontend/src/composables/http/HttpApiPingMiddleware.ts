import { markRaw, type Ref } from 'vue';
import { requestIdleCallback } from '@/helpers';
import type { MaybePromise } from '@/types';
import { useAuthorized } from '../useAuthorized';
import type { HttpClient } from './HttpClient';
import type {
  IHttpBeforeRequestInterceptor,
  IHttpResponseErrorInterceptor,
  IHttpResponseSuccessInterceptor,
  IInterceptorContext,
} from './HttpMiddlewareExecutor';

export class HttpApiPingMiddleware implements IHttpBeforeRequestInterceptor,
  IHttpResponseSuccessInterceptor,
  IHttpResponseErrorInterceptor {
  static use(http: HttpClient): IHttpBeforeRequestInterceptor {
    return markRaw(new HttpApiPingMiddleware(http, useAuthorized()));
  }

  #timeoutId: TimeoutId | null = null;
  #isBackendDown = false;

  constructor(
    private readonly http: HttpClient,
    private readonly authorized: Ref<boolean>,
  ) {
    requestIdleCallback(this.#pingApi.bind(this));
  }

  interceptBeforeRequest(_: Request, ctx: IInterceptorContext): void {
    if (ctx.meta.isPing) return;
    if (this.#timeoutId) clearTimeout(this.#timeoutId);
  }

  interceptResponseSuccess(): MaybePromise<void> {
    const wasBackendDown = this.#isBackendDown;
    this.#isBackendDown = false;

    this.#schedulePing();

    if (wasBackendDown) {
      window.location.reload();
    }
  }

  interceptResponseError(_: Response, ctx: IInterceptorContext): MaybePromise<void> {
    if (ctx.meta.isPing) {
      this.#isBackendDown = true;
    }

    this.#schedulePing();
  }

  #schedulePing(): void {
    if (this.#timeoutId) clearTimeout(this.#timeoutId);
    const timeout = this.#isBackendDown ? 500 : 60_000;
    this.#timeoutId = setTimeout(this.#pingApi.bind(this), timeout);
  }

  async #pingApi(): Promise<void> {
    if (!this.authorized.value) return;

    try {
      await this.http.get('/ping', {}, {
        meta: {
          unauthorizedRedirect: false,
          isPing: true,
        },
      });
    } catch (error) {
      console.error(error);
    }
  }
}
