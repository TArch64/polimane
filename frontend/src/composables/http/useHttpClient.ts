import { type FunctionPlugin, inject, type InjectionKey } from 'vue';
import { HttpClient } from './HttpClient';
import { HttpMiddlewareExecutor } from './HttpMiddlewareExecutor';
import { HttpAuthorization } from './HttpAuthorization';
import { HttpApiPing } from './HttpApiPing';

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
    client.middleware(HttpAuthorization.use());
    client.middleware(HttpApiPing.use(client));
  });
};

export const useHttpClient = () => inject(Provider)!;
