import { addBreadcrumb } from '@sentry/vue';
import { buildUrl, type UrlParams, type UrlPath } from '@/helpers';
import {
  HttpLegacyTransport,
  HttpModernTransport,
  HttpTransport,
  type IHttpTransport,
} from './transports';
import { HttpError } from './HttpError';
import type { HttpMiddlewareExecutor, MiddlewareConstructor } from './HttpMiddlewareExecutor';

export type HttpBody = object | string;

export interface IHttpClientOptions {
  baseUrl: string;
  middlewareExecutor: HttpMiddlewareExecutor;
}

export interface IHttpRequestConfig {
  meta?: Record<string, unknown>;
  responseType?: 'json' | 'text';
  transport?: HttpTransport;
}

interface IRequestConfig<
  P extends UrlParams = UrlParams,
  B extends HttpBody = HttpBody,
> extends IHttpRequestConfig {
  method: 'GET' | 'POST' | 'PATCH' | 'DELETE';
  path: UrlPath;
  params?: P;
  body?: B;
}

export class HttpClient {
  private readonly baseUrl;
  private readonly middlewareExecutor;
  private readonly transports: Partial<Record<HttpTransport, IHttpTransport>> = {};

  constructor(options: IHttpClientOptions) {
    this.baseUrl = options.baseUrl;

    this.middlewareExecutor = options.middlewareExecutor;
    this.middlewareExecutor.client = this;
  }

  get<R extends HttpBody, P extends UrlParams = UrlParams>(
    path: UrlPath,
    params: P = {} as P,
    config: IHttpRequestConfig = {},
  ): Promise<R> {
    return this.request({
      method: 'GET',
      path,
      params,
      ...config,
    });
  }

  delete<R extends HttpBody, B extends HttpBody = HttpBody>(
    path: UrlPath,
    body: B = {} as B,
    config: IHttpRequestConfig = {},
  ): Promise<R> {
    return this.request({
      method: 'DELETE',
      path,
      body,
      ...config,
    });
  }

  post<R extends HttpBody, B extends HttpBody = HttpBody>(
    path: UrlPath,
    body: B,
    config: IHttpRequestConfig = {},
  ): Promise<R> {
    return this.request({
      method: 'POST',
      path,
      body,
      ...config,
    });
  }

  patch<R extends HttpBody, B extends HttpBody = HttpBody>(
    path: UrlPath,
    body: B,
    config: IHttpRequestConfig = {},
  ): Promise<R> {
    return this.request({
      method: 'PATCH',
      path,
      body,
      ...config,
    });
  }

  middleware<M extends MiddlewareConstructor>(Class: M): void {
    this.middlewareExecutor.add(Class);
  }

  getMiddleware<M extends MiddlewareConstructor>(Class: M): InstanceType<M> | null {
    return this.middlewareExecutor.get(Class);
  }

  private async request<
    R extends HttpBody,
    P extends UrlParams,
    B extends HttpBody,
  >(config: IRequestConfig<P, B>): Promise<R> {
    const body = config.body ? JSON.stringify(config.body) : undefined;
    const responseType = config.responseType ?? 'json';

    const request = new Request(this.buildUrl(config), {
      method: config.method,
      body,

      headers: {
        'Content-Type': responseType === 'text'
          ? 'text/plain'
          : 'application/json',
      },

      credentials: 'include',
    });

    await this.middlewareExecutor.callBeforeRequestInterceptor(request);
    try {
      const transport = this.getTransport(config);
      const response = await transport.send(request);

      if (!response.ok) {
        await this.handleError(response, config);
      }

      await this.middlewareExecutor.callResponseSuccessInterceptor(response);
      return response[responseType]();
    } catch (error) {
      addBreadcrumb({
        type: 'http-error',
        level: 'error',
        message: 'HTTP request failed',
        data: {
          exception: error,
          requestBody: body,
        },
      });

      throw error;
    }
  }

  private buildUrl(config: IRequestConfig): URL {
    return buildUrl(this.baseUrl, config.path, config.params);
  }

  private async handleError(response: Response, config: IRequestConfig): Promise<never> {
    const error = await HttpError.fromResponse(response);
    error.meta = config.meta ?? {};
    await this.middlewareExecutor.callResponseErrorInterceptor(error);
    throw error;
  }

  private getTransport(config: IRequestConfig): IHttpTransport {
    const transport = config.transport ?? HttpTransport.MODERN;

    if (!(transport in this.transports)) {
      this.transports[transport] = this.createTransport(transport);
    }

    return this.transports[transport]!;
  }

  private createTransport(transport: HttpTransport): IHttpTransport {
    switch (transport) {
      case HttpTransport.MODERN:
        return new HttpModernTransport();
      case HttpTransport.LEGACY:
        return new HttpLegacyTransport();
    }
  }
}
