import type { Constructor, MaybePromise } from '@/types';
import type { HttpError } from './HttpError';
import type { HttpClient } from './HttpClient';

export type MiddlewareConstructor = Constructor<HttpMiddleware> & {
  use: (client: HttpClient) => HttpMiddleware;
};

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
  private readonly middlewares = new Map<MiddlewareConstructor, HttpMiddleware>();
  client!: HttpClient;

  get list(): HttpMiddleware[] {
    return Array.from(this.middlewares.values());
  }

  add<M extends MiddlewareConstructor>(Class: M): void {
    this.middlewares.set(Class, Class.use(this.client));
  }

  get<M extends MiddlewareConstructor>(Class: M): InstanceType<M> | null {
    return (this.middlewares.get(Class) as InstanceType<M>) || null;
  }

  async callBeforeRequestInterceptor(request: Request): Promise<void> {
    const middlewares = this.list.filter((m): m is IHttpBeforeRequestInterceptor => 'interceptBeforeRequest' in m);

    for (const middleware of middlewares) {
      await middleware.interceptBeforeRequest(request);
    }
  }

  async callResponseErrorInterceptor(error: HttpError): Promise<void> {
    const middlewares = this.list.filter((m): m is IHttpResponseErrorInterceptor => 'interceptResponseError' in m);

    for (const middleware of middlewares) {
      await middleware.interceptResponseError(error);
    }
  }

  async callResponseSuccessInterceptor(response: Response): Promise<void> {
    const middlewares = this.list.filter((m): m is IHttpResponseSuccessInterceptor => 'interceptResponseSuccess' in m);

    for (const middleware of middlewares) {
      await middleware.interceptResponseSuccess(response);
    }
  }
}
