import { SubscriptionLimit, SubscriptionPlan, SubscriptionStatus, type UserLimit } from '@/enums';

export type UserCounters = Record<UserLimit, number>;
export type SubscriptionLimits = Partial<Record<SubscriptionLimit, number>>;

export interface IUserSubscription {
  plan: SubscriptionPlan;
  status: SubscriptionStatus;
  counters: UserCounters;
  limits: SubscriptionLimits;
}
