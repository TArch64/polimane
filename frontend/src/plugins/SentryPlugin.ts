import { type App, type FunctionPlugin, watch } from 'vue';
import { init, setUser, type User } from '@sentry/vue';
import { useSessionStore } from '@/stores';
import type { IUser } from '@/models';

export interface ISentryPluginOptions {
  dsn: string;
}

export class SentryPlugin {
  static install: FunctionPlugin<ISentryPluginOptions> = (app, options) => {
    if (options.dsn) {
      new SentryPlugin(app, options);
    }
  };

  constructor(
    private readonly app: App,
    private readonly options: ISentryPluginOptions,
  ) {
    this.init();

    this.app.runWithContext(() => {
      this.linkUser();
    });
  }

  private init() {
    init({
      app: this.app,
      dsn: this.options.dsn,
      environment: import.meta.env.MODE,
    });
  }

  private linkUser() {
    const sessionStore = useSessionStore();

    watch(() => sessionStore.user, (user) => {
      setUser(this.buildSentryUser(user));
    }, { immediate: true, deep: true });
  }

  private buildSentryUser(user: IUser | null): User | null {
    if (!user) {
      return null;
    }

    return {
      id: user.id,
      username: this.buildUsername(user),
    };
  }

  private buildUsername(user: IUser): string | undefined {
    if (!user.firstName && !user.lastName) {
      return undefined;
    }

    let username = user.firstName || '';

    if (user.lastName) {
      const prefix = username ? ' ' : '';
      const lastName = username ? user.lastName.charAt(0) + '.' : user.lastName;
      username += prefix + lastName;
    }

    return username;
  }
}
