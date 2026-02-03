import { AccessLevel } from '@/enums';

export interface ISchemaUser {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  access: AccessLevel;
  isEvenAccess: boolean;
  isAllAccess: boolean;
}

export interface ISchemaUserInvitation {
  email: string;
  access: AccessLevel;
  isEvenAccess: boolean;
  isAllAccess: boolean;
}
