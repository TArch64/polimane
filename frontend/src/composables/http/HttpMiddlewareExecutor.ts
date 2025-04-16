import type { MaybePromise } from '@/types';
import type { HttpError } from './HttpError';

export interface IHttpMiddleware {
  interceptResponseError: (error: HttpError) => MaybePromise<void>;
}

export class HttpMiddlewareExecutor {
  private readonly middlewares: IHttpMiddleware[] = [];

  add(middleware: IHttpMiddleware): void {
    this.middlewares.push(middleware);
  }

  async callResponseErrorInterceptor(error: HttpError): Promise<void> {
    for (const middleware of this.middlewares) {
      await middleware.interceptResponseError(error);
    }
  }
}
