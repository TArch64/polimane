import { AccessLevel } from '@/enums';

export interface ISchemaUser {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  access: AccessLevel;
}
