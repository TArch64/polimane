import { type FunctionPlugin, inject, type InjectionKey } from 'vue';

const Provider = Symbol('HttpClient') as InjectionKey<HttpClient>;

export type HttpBody = object;
export type HttpParams = Record<string, string | number>;
export type PathItem = string;
export type Path = PathItem[] | PathItem;

interface IRequestConfig<P extends HttpParams, B extends HttpBody> {
  method: 'GET' | 'POST' | 'PATCH' | 'DELETE';
  path: Path;
  params?: P;
  body?: B;
}

export class HttpError extends Error {
  constructor(
    readonly response: unknown,
  ) {
    super('HTTP error');
  }
}

export class HttpClient {
  constructor(
    private readonly baseUrl: string,
  ) {
  }

  get<R extends HttpBody, P extends HttpParams = HttpParams>(path: Path, params: P = {} as P): Promise<R> {
    return this.request({
      method: 'GET',
      path,
      params,
    });
  }

  delete<R extends HttpBody, P extends HttpParams = HttpParams>(path: Path, params: P = {} as P): Promise<R> {
    return this.request({
      method: 'DELETE',
      path,
      params,
    });
  }

  post<R extends HttpBody, B extends HttpBody>(path: Path, body: B): Promise<R> {
    return this.request({
      method: 'POST',
      path,
      body,
    });
  }

  patch<R extends HttpBody, B extends HttpBody>(path: Path, body: B): Promise<R> {
    return this.request({
      method: 'PATCH',
      path,
      body,
    });
  }

  private async request<
    R extends HttpBody,
    P extends HttpParams,
    B extends HttpBody,
  >(config: IRequestConfig<P, B>): Promise<R> {
    const body = config.body ? JSON.stringify(config.body) : undefined;

    const response = await fetch(this.buildUrl(config), {
      method: config.method,
      body,

      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      return this.handleError(response);
    }

    return response.json();
  }

  private buildUrl<P extends HttpParams>(config: IRequestConfig<P, HttpBody>): URL {
    const path = [config.path].flat().join('/');
    const url = new URL(this.baseUrl + path);

    if (config.params && Object.keys(config.params)) {
      url.search = new URLSearchParams(config.params as Record<string, string>).toString();
    }

    return url;
  }

  private async handleError(response: Response): Promise<never> {
    const headers = response.headers ?? new Headers();
    const isJson = headers.get('content-type')?.includes('application/json');
    const body = isJson ? await response.json() : await response.text();
    throw new HttpError(body);
  }
}

export interface IPluginHttpClientOptions {
  baseUrl: string;
}

export const pluginHttpClient: FunctionPlugin<IPluginHttpClientOptions> = (app, options) => {
  app.provide(Provider, new HttpClient(options.baseUrl));
};

export const useHttpClient = () => inject(Provider)!;
