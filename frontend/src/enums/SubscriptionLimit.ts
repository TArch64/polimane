export enum SubscriptionLimit {
  SCHEMAS_CREATED = 'schemasCreated',
  SCHEMA_BEADS = 'schemaBeads',
  SHARED_ACCESS = 'sharedAccess',
}

export const UserLimits = [
  SubscriptionLimit.SCHEMAS_CREATED,
] as const;

export type UserLimit = typeof UserLimits[number];

export function isUserLimit(limit: string): limit is UserLimit {
  return UserLimits.includes(limit as UserLimit);
}

export const SchemaLimits = [
  SubscriptionLimit.SCHEMA_BEADS,
  SubscriptionLimit.SHARED_ACCESS,
] as const;

export type SchemaLimit = typeof SchemaLimits[number];

export function isSchemaLimit(limit: string): limit is SchemaLimit {
  return SchemaLimits.includes(limit as SchemaLimit);
}

export {
  isSchemaLimit as isFeatureLimit,
  SchemaLimits as FeatureLimits,
  type SchemaLimit as FeatureLimit,
};
