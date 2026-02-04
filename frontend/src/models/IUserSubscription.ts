import { SubscriptionStatus } from '@/enums';
import type { ISubscriptionPlan, UserCounters } from './ISubscriptionPlan';

export type UserActivePlan = Pick<ISubscriptionPlan, 'id' | 'tier' | 'limits'>;

export interface IUserSubscription {
  status: SubscriptionStatus;
  counters: UserCounters;
  plan: UserActivePlan;
}
