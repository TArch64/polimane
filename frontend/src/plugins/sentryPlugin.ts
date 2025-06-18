import type { FunctionPlugin } from 'vue';
import { init } from '@sentry/vue';

export interface ISentryPluginOptions {
  dsn: string;
}

export const sentryPlugin: FunctionPlugin<ISentryPluginOptions> = (app, options) => {
  options.dsn && init({
    app,
    dsn: options.dsn,
  });
};
