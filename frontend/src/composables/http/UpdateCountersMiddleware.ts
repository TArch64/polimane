import { markRaw } from 'vue';
import type { MaybePromise } from '@/types';
import { Callback, type IUser, type SchemaCounters, type UserCounters } from '@/models';
import { SchemaLimits, UserLimits } from '@/enums';
import { decryptXor, parseIdBytes } from '@/helpers';
import type { IHttpResponseSuccessInterceptor } from './HttpMiddlewareExecutor';

type AnyLimitKeys = readonly string[];
type AnyCounter<K extends AnyLimitKeys> = Record<K[number], number>;

export class UpdateCountersMiddleware implements IHttpResponseSuccessInterceptor {
  static use(): IHttpResponseSuccessInterceptor {
    return markRaw(new UpdateCountersMiddleware());
  }

  user: IUser | null = null;
  onUserUpdate = Callback.create<[counters: UserCounters]>();
  onSchemaUpdate = Callback.create<[counters: SchemaCounters]>();

  interceptResponseSuccess(response: Response): MaybePromise<void> {
    if (!this.user) return;
    this.#updateUserCounter(response);
    this.#updateSchemaCounter(response);
  }

  #updateUserCounter(response: Response): void {
    const counters = this.#getCounter(response, 'X-UC', UserLimits);

    if (counters) {
      this.onUserUpdate.dispatch(counters);
    }
  }

  #updateSchemaCounter(response: Response): void {
    const counters = this.#getCounter(response, 'X-SC', SchemaLimits);

    if (counters) {
      this.onSchemaUpdate.dispatch(counters);
    }
  }

  #getCounter<K extends AnyLimitKeys>(response: Response, header: string, keys: K): AnyCounter<K> | null {
    const encoded = response.headers.get(header);
    if (!encoded) return null;

    try {
      const counters = JSON.parse(this.#decodeInput(encoded));
      if (!counters) return null;

      return this.#validate(counters, keys);
    } catch {
      return null;
    }
  }

  #decodeInput(encoded: string): string {
    const key = parseIdBytes(this.user!.id);
    const encrypted = Uint8Array.from(atob(encoded), (c) => c.charCodeAt(0));
    return new TextDecoder().decode(decryptXor(key, encrypted));
  }

  #validate<K extends AnyLimitKeys>(input: object, keys: K): AnyCounter<K> | null {
    const output = {} as Record<K[number], number>;

    for (const key of keys) {
      // @ts-expect-error hard to type correctly
      const value = input[key] as unknown;
      if (value && typeof value === 'number' && value > 0) {
        output[key as K[number]] = value;
      } else {
        location.reload();
        return null;
      }
    }

    return output;
  }
}
