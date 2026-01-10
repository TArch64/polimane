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

export const FeatureLimits = [
  SubscriptionLimit.SCHEMA_BEADS,
  SubscriptionLimit.SHARED_ACCESS,
] as const;

export type FeatureLimit = typeof FeatureLimits[number];

export function isFeatureLimit(limit: string): limit is FeatureLimit {
  return FeatureLimits.includes(limit as FeatureLimit);
}
