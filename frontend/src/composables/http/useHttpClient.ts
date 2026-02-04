import { type FunctionPlugin, inject, type InjectionKey, markRaw } from 'vue';
import { HttpClient } from './HttpClient';
import { HttpMiddlewareExecutor } from './HttpMiddlewareExecutor';
import { HttpAuthorizationMiddleware } from './HttpAuthorizationMiddleware';
import { HttpApiPingMiddleware } from './HttpApiPingMiddleware';
import { UpdateCountersMiddleware } from './UpdateCountersMiddleware';

const Provider = Symbol('HttpClient') as InjectionKey<HttpClient>;

export interface IPluginHttpClientOptions {
  baseUrl: string;
}

export const httpClientPlugin: FunctionPlugin<IPluginHttpClientOptions> = (app, options) => {
  const client = markRaw(new HttpClient({
    baseUrl: options.baseUrl,
    middlewareExecutor: markRaw(new HttpMiddlewareExecutor()),
  }));

  app.provide(Provider, client);

  app.runWithContext(() => {
    client.middleware(HttpAuthorizationMiddleware);
    client.middleware(HttpApiPingMiddleware);
    client.middleware(UpdateCountersMiddleware);
  });
};

export const useHttpClient = () => inject(Provider)!;
