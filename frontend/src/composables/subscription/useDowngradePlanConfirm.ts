import type { Ref } from 'vue';
import { type IConfirm, type IConfirmResult, useConfirm } from '@/components/confirm';
import type { ISubscriptionPlan } from '@/models';
import { useSchemasCreatedCounter } from './userCounters';

export function useDowngradePlanConfirm(plan: Ref<ISubscriptionPlan>): IConfirm {
  const schemasCreatedCounter = useSchemasCreatedCounter();

  const unavailableSchemasConfirm = useConfirm({
    message: `Ви маєте більше схем ніж дозволяє ця підписка тому їх редагування буде недоступне до тих пір поки ви не зменшите їх кількість або не повернеся назад до вищої підписки`,
    acceptButton: 'Продовжити',
  });

  async function ask(): Promise<IConfirmResult> {
    if (schemasCreatedCounter.current > plan.value.limits.schemasCreated!) {
      return unavailableSchemasConfirm.ask();
    }

    return { isAccepted: true };
  }

  return {
    anchorStyle: unavailableSchemasConfirm.anchorStyle,
    ask,
  };
}
