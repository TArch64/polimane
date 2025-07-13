import { HttpError } from './HttpError';
import type { HttpMiddleware, HttpMiddlewareExecutor } from './HttpMiddlewareExecutor';

export type HttpBody = object;
export type HttpParams = Record<string, string | number>;
export type PathItem = string | number;
export type Path = PathItem[] | PathItem;

export interface IHttpClientOptions {
  baseUrl: string;
  middlewareExecutor: HttpMiddlewareExecutor;
}

export interface IHttpRequestConfig {
  meta?: Record<string, unknown>;
}

interface IRequestConfig<
  P extends HttpParams = HttpParams,
  B extends HttpBody = HttpBody,
> extends IHttpRequestConfig {
  method: 'GET' | 'POST' | 'PATCH' | 'DELETE';
  path: Path;
  params?: P;
  body?: B;
}

export class HttpClient {
  private readonly baseUrl;
  private readonly middlewareExecutor;

  constructor(options: IHttpClientOptions) {
    this.baseUrl = options.baseUrl;
    this.middlewareExecutor = options.middlewareExecutor;
  }

  get<R extends HttpBody, P extends HttpParams = HttpParams>(
    path: Path,
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

  delete<R extends HttpBody, P extends HttpParams = HttpParams>(
    path: Path,
    params: P = {} as P,
    config: IHttpRequestConfig = {},
  ): Promise<R> {
    return this.request({
      method: 'DELETE',
      path,
      params,
      ...config,
    });
  }

  post<R extends HttpBody, B extends HttpBody>(
    path: Path,
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

  patch<R extends HttpBody, B extends HttpBody>(
    path: Path,
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

  private async request<
    R extends HttpBody,
    P extends HttpParams,
    B extends HttpBody,
  >(config: IRequestConfig<P, B>): Promise<R> {
    const body = config.body ? JSON.stringify(config.body) : undefined;

    const request = new Request(this.buildUrl(config), {
      method: config.method,
      body,
      headers: { 'Content-Type': 'application/json' },
    });

    await this.middlewareExecutor.callBeforeRequestInterceptor(request);
    const response = await fetch(request);

    if (!response.ok) {
      return this.handleError(response, config);
    }

    await this.middlewareExecutor.callResponseSuccessInterceptor(response);
    return response.json();
  }

  private buildUrl(config: IRequestConfig): URL {
    const path = [config.path].flat().join('/');
    const url = new URL(this.baseUrl + path);

    if (config.params && Object.keys(config.params).length) {
      url.search = new URLSearchParams(config.params as Record<string, string>).toString();
    }

    return url;
  }

  private async handleError(response: Response, config: IRequestConfig): Promise<never> {
    const error = await HttpError.fromResponse(response);
    error.meta = config.meta ?? {};
    await this.middlewareExecutor.callResponseErrorInterceptor(error);
    throw error;
  }

  middleware(middleware: HttpMiddleware): void {
    this.middlewareExecutor.add(middleware);
  }
}
