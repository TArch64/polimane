import type { MaybePromise } from '@/types';
import type { HttpError } from './HttpError';

export interface IHttpBeforeRequestInterceptor {
  interceptBeforeRequest(request: Request): MaybePromise<void>;
}

export interface IHttpResponseErrorInterceptor {
  interceptResponseError(error: HttpError): MaybePromise<void>;
}

export interface IHttpResponseSuccessInterceptor {
  interceptResponseSuccess(response: Response): MaybePromise<void>;
}

export type HttpMiddleware
  = IHttpBeforeRequestInterceptor
    | IHttpResponseErrorInterceptor
    | IHttpResponseSuccessInterceptor;

export class HttpMiddlewareExecutor {
  private readonly middlewares: HttpMiddleware[] = [];

  add(middleware: HttpMiddleware): void {
    this.middlewares.push(middleware);
  }

  async callBeforeRequestInterceptor(request: Request): Promise<void> {
    const middlewares = this.middlewares.filter((m): m is IHttpBeforeRequestInterceptor => 'interceptBeforeRequest' in m);

    for (const middleware of middlewares) {
      await middleware.interceptBeforeRequest(request);
    }
  }

  async callResponseErrorInterceptor(error: HttpError): Promise<void> {
    const middlewares = this.middlewares.filter((m): m is IHttpResponseErrorInterceptor => 'interceptResponseError' in m);

    for (const middleware of middlewares) {
      await middleware.interceptResponseError(error);
    }
  }

  async callResponseSuccessInterceptor(response: Response): Promise<void> {
    const middlewares = this.middlewares.filter((m): m is IHttpResponseSuccessInterceptor => 'interceptResponseSuccess' in m);

    for (const middleware of middlewares) {
      await middleware.interceptResponseSuccess(response);
    }
  }
}
