import type {
  _ExtractActionsFromSetupStore,
  _ExtractGettersFromSetupStore,
  _ExtractStateFromSetupStore,
  StoreDefinition,
} from 'pinia';
import type { SafeAny } from '@/types';

export type StoreFn = (...args: SafeAny) => SafeAny;

export type InferStoreDefinition<ID extends string, S extends StoreFn> = StoreDefinition<
  ID,
  _ExtractStateFromSetupStore<ReturnType<S>>,
  _ExtractGettersFromSetupStore<ReturnType<S>>,
  _ExtractActionsFromSetupStore<ReturnType<S>>
>;

export type InferStore<ID extends string, S extends StoreFn> = ReturnType<InferStoreDefinition<ID, S>>;
