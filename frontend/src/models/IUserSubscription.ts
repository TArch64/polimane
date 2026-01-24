import { SubscriptionPlanId, SubscriptionStatus } from '@/enums';
import type { SubscriptionLimits, UserCounters } from './ISubscriptionPlan';

export interface IUserSubscription {
  planId: SubscriptionPlanId;
  status: SubscriptionStatus;
  counters: UserCounters;
  limits: SubscriptionLimits;
}
