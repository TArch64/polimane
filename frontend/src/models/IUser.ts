import { SubscriptionPlan, SubscriptionStatus } from '@/enums';

export interface ISubscriptionCounters {
  schemasCreated: number;
}

export interface ISubscriptionLimits {
  schemasCreated?: number;
  schemaBeads?: number;
  sharedAccess?: number;
}

export interface IUserSubscription {
  plan: SubscriptionPlan;
  status: SubscriptionStatus;
  counters: ISubscriptionCounters;
  limits: ISubscriptionLimits;
}

export interface IUser {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  isEmailVerified: boolean;
  subscription: IUserSubscription;
}
