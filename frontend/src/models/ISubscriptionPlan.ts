import { type SchemaLimit, SubscriptionLimit, SubscriptionPlanId, type UserLimit } from '@/enums';

export type UserCounters = Record<UserLimit, number>;
export type SchemaCounters = Record<SchemaLimit, number>;
export type SubscriptionLimits = Partial<Record<SubscriptionLimit, number>>;

export interface ISubscriptionPlan {
  id: SubscriptionPlanId;
  tier: number;
  monthlyPrice: number;
  limits: SubscriptionLimits;
}
