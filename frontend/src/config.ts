import type { SubscriptionLimits } from '@/models';
import { SubscriptionLimitType } from '@/enums';

export const LOCALE = 'uk-UA';
export const isMac = navigator.platform.toUpperCase().includes('MAC');
export const DEFAULT_SCHEMA_BACKGROUND = '#f8f8f8';

export const INFINITY_SYMBOL = '∞';

type LimitKey = keyof SubscriptionLimits;

export interface IPlanLimitConfig {
  type: SubscriptionLimitType;
  title: string;
}

export const PLAN_LIMIT_CONFIGS: Record<LimitKey, IPlanLimitConfig> = {
  schemasCreated: {
    type: SubscriptionLimitType.USER,
    title: 'Схеми',
  },

  schemaBeads: {
    type: SubscriptionLimitType.FEATURE,
    title: 'Бісер в Схемі',
  },

  sharedAccess: {
    type: SubscriptionLimitType.FEATURE,
    title: 'Користувачі Схеми',
  },
};
