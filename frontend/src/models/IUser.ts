import { SubscriptionPlan, SubscriptionStatus } from '@/enums';

export interface IUserSubscription {
  plan: SubscriptionPlan;
  status: SubscriptionStatus;
}

export interface IUser {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  isEmailVerified: boolean;
  subscription: IUserSubscription;
}
