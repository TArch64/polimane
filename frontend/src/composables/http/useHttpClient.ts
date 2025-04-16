import { type FunctionPlugin, inject, type InjectionKey } from 'vue';
import { HttpClient } from './HttpClient';
import { HttpMiddlewareExecutor } from './HttpMiddlewareExecutor';
import { useUnauthorizedMiddleware } from './useUnauthorizedMiddleware';

const Provider = Symbol('HttpClient') as InjectionKey<HttpClient>;

export interface IPluginHttpClientOptions {
  baseUrl: string;
}

export const httpClientPlugin: FunctionPlugin<IPluginHttpClientOptions> = (app, options) => {
  const client = new HttpClient({
    baseUrl: options.baseUrl,
    middlewareExecutor: new HttpMiddlewareExecutor(),
  });

  app.provide(Provider, client);

  app.runWithContext(() => {
    client.middleware(useUnauthorizedMiddleware());
  });
};

export const useHttpClient = () => inject(Provider)!;
