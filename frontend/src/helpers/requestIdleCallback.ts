import type { SafeAny } from '@/types';

export const requestIdleCallback = window.requestIdleCallback || ((callback: SafeAny) => setTimeout(callback)) as typeof window.requestIdleCallback;
export const cancelIdleCallback = window.cancelIdleCallback || window.clearTimeout as typeof window.cancelIdleCallback;
