import type { MaybePromise } from '@/types';
import { Callback, type UserCounters } from '@/models';
import { UserLimits } from '@/enums';
import type { IHttpResponseSuccessInterceptor } from './HttpMiddlewareExecutor';

export class UpdateUserCountersMiddleware implements IHttpResponseSuccessInterceptor {
  static use(): IHttpResponseSuccessInterceptor {
    return new UpdateUserCountersMiddleware();
  }

  readonly onUpdate = new Callback<[counters: UserCounters]>();

  interceptResponseSuccess(response: Response): MaybePromise<void> {
    const input = this.getCountersInput(response);
    if (!input) return;

    const counters = this.validateCounters(input);
    if (!counters) return;

    this.onUpdate.dispatch(counters);
  }

  private getCountersInput(response: Response): object | null {
    const encoded = response.headers.get('X-User-Counters');
    if (!encoded) return null;

    try {
      return JSON.parse(atob(encoded));
    } catch {
      return null;
    }
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
