import { type SchemaLimit, SubscriptionLimit, SubscriptionPlanId, type UserLimit } from '@/enums';

export type UserCounters = Record<UserLimit, number>;
export type SchemaCounters = Record<SchemaLimit, number>;
export type SubscriptionLimits = Partial<Record<SubscriptionLimit, number>>;

export interface ISubscriptionPlan {
  id: SubscriptionPlanId;
  monthlyPrice: number;
  yearlyPrice: number;
  limits: SubscriptionLimits;
}
