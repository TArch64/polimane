import type { IUserSubscription } from './IUserSubscription';

export interface IUser {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  isEmailVerified: boolean;
  subscription: IUserSubscription;
}
