import type { MaybePromise } from '@/types';
import type { UserCounters } from '@/models';
import { UserLimits } from '@/enums';
import type { useSessionStore } from '@/stores';
import { decryptXor, parseIdBytes } from '@/helpers';
import type { IHttpResponseSuccessInterceptor } from './HttpMiddlewareExecutor';

export class UpdateUserCountersMiddleware implements IHttpResponseSuccessInterceptor {
  static use(): IHttpResponseSuccessInterceptor {
    return new UpdateUserCountersMiddleware();
  }

  private sessionStore!: ReturnType<typeof useSessionStore>;

  setSessionStore(store: ReturnType<typeof useSessionStore>): void {
    this.sessionStore = store;
  }

  interceptResponseSuccess(response: Response): MaybePromise<void> {
    if (!this.sessionStore.isLoggedIn) return;

    const input = this.getCountersInput(response);
    if (!input) return;

    const counters = this.validateCounters(input);
    if (!counters) return;

    this.sessionStore.updateUser({
      subscription: {
        ...this.sessionStore.user.subscription,
        counters,
      },
    });
  }

  private getCountersInput(response: Response): object | null {
    const encoded = response.headers.get('X-User-Counters');
    if (!encoded) return null;

    try {
      return JSON.parse(this.decodeCountersInput(encoded));
    } catch {
      return null;
    }
  }

  private decodeCountersInput(encoded: string): string {
    const key = parseIdBytes(this.sessionStore.user.id);
    const encrypted = Uint8Array.from(atob(encoded), (c) => c.charCodeAt(0));
    return new TextDecoder().decode(decryptXor(key, encrypted));
  }

  private validateCounters(input: object): UserCounters | null {
    const output: UserCounters = {} as UserCounters;

    for (const key of UserLimits) {
      if (key in input && typeof input[key] === 'number') {
        output[key] = input[key];
      } else {
        location.reload();
        return null;
      }
    }

    return output;
  }
}
