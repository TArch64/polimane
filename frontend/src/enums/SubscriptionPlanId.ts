import { getMappedValue } from '@/helpers';

export enum SubscriptionPlanId {
  BETA = 'beta',
  BASIC = 'basic',
  PRO = 'pro',
}

export function getSubscriptionPlanName(planId: SubscriptionPlanId): string {
  return getMappedValue(planId, {
    [SubscriptionPlanId.BETA]: 'Бета',
    [SubscriptionPlanId.BASIC]: 'Базова',
    [SubscriptionPlanId.PRO]: 'Продвинута',
  });
}

export function isMaxSubscriptionPlan(planId: SubscriptionPlanId): boolean {
  return [SubscriptionPlanId.PRO, SubscriptionPlanId.BETA].includes(planId);
}
