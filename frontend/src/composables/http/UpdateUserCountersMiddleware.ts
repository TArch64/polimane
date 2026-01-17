import type { MaybePromise } from '@/types';
import type { UserCounters } from '@/models';
import { UserLimits } from '@/enums';
import type { useSessionStore } from '@/stores';
import { decryptXor, parseIdBytes } from '@/helpers';
import type { IHttpResponseSuccessInterceptor } from './HttpMiddlewareExecutor';

const USER_COUNTERS_HEADER_BYTES = [88, 45, 85, 115, 101, 114, 45, 67, 111, 117, 110, 116, 101, 114, 115];
const USER_COUNTERS_HEADER = String.fromCharCode(...USER_COUNTERS_HEADER_BYTES);

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

    const input = this.getInput(response);
    if (!input) return;

    const counters = this.validate(input);
    if (!counters) return;

    this.sessionStore.updateUser({
      subscription: {
        ...this.sessionStore.user.subscription,
        counters,
      },
    });
  }

  private getInput(response: Response): object | null {
    const encoded = response.headers.get(USER_COUNTERS_HEADER);
    if (!encoded) return null;

    try {
      return JSON.parse(this.decodeInput(encoded));
    } catch {
      return null;
    }
  }

  private decodeInput(encoded: string): string {
    const key = parseIdBytes(this.sessionStore.user.id);
    const encrypted = Uint8Array.from(atob(encoded), (c) => c.charCodeAt(0));
    return new TextDecoder().decode(decryptXor(key, encrypted));
  }

  private validate(input: object): UserCounters | null {
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
