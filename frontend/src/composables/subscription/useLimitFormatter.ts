import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { useNumberFormatter } from '@/composables';
import type { ISubscriptionPlan } from '@/models';
import { SubscriptionLimit } from '@/enums';
import { INFINITY_SYMBOL } from '@/config';

export interface ILimitFormatterOptions {
  plan: MaybeRefOrGetter<ISubscriptionPlan>;
  limit: MaybeRefOrGetter<SubscriptionLimit>;
}

export function useLimitFormatter(options: ILimitFormatterOptions) {
  const plan = computed(() => toValue(options.plan));
  const limit = computed(() => toValue(options.limit));

  const number = useNumberFormatter(() => plan.value.limits[limit.value]);

  return computed(() => number.value || INFINITY_SYMBOL);
}
