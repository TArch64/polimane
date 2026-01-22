import type { Constructor, MaybePromise } from '@/types';
import type { HttpClient } from './HttpClient';

export type MiddlewareConstructor = Constructor<HttpMiddleware> & {
  use: (client: HttpClient) => HttpMiddleware;
};

export interface IInterceptorContext {
  meta: Record<string, unknown>;
}

export interface IHttpBeforeRequestInterceptor {
  interceptBeforeRequest(request: Request, ctx: IInterceptorContext): MaybePromise<void>;
}

export interface IHttpResponseErrorInterceptor {
  interceptResponseError(error: unknown, ctx: IInterceptorContext): MaybePromise<void>;
}

export interface IHttpResponseSuccessInterceptor {
  interceptResponseSuccess(response: Response, ctx: IInterceptorContext): MaybePromise<void>;
}

export type HttpMiddleware
  = IHttpBeforeRequestInterceptor
    | IHttpResponseErrorInterceptor
    | IHttpResponseSuccessInterceptor;

export class HttpMiddlewareExecutor {
  #middlewares = new Map<MiddlewareConstructor, HttpMiddleware>();
  client!: HttpClient;

  get list(): HttpMiddleware[] {
    return Array.from(this.#middlewares.values());
  }

  add<M extends MiddlewareConstructor>(Class: M): void {
    this.#middlewares.set(Class, Class.use(this.client));
  }

  get<M extends MiddlewareConstructor>(Class: M): InstanceType<M> | null {
    return (this.#middlewares.get(Class) as InstanceType<M>) || null;
  }

  async callBeforeRequestInterceptor(request: Request, ctx: IInterceptorContext): Promise<void> {
    const middlewares = this.list.filter((m): m is IHttpBeforeRequestInterceptor => 'interceptBeforeRequest' in m);

    for (const middleware of middlewares) {
      await middleware.interceptBeforeRequest(request, ctx);
    }
  }

  async callResponseErrorInterceptor(error: unknown, ctx: IInterceptorContext): Promise<void> {
    const middlewares = this.list.filter((m): m is IHttpResponseErrorInterceptor => 'interceptResponseError' in m);

    for (const middleware of middlewares) {
      await middleware.interceptResponseError(error, ctx);
    }
  }

  async callResponseSuccessInterceptor(response: Response, ctx: IInterceptorContext): Promise<void> {
    const middlewares = this.list.filter((m): m is IHttpResponseSuccessInterceptor => 'interceptResponseSuccess' in m);

    for (const middleware of middlewares) {
      await middleware.interceptResponseSuccess(response, ctx);
    }
  }
}
