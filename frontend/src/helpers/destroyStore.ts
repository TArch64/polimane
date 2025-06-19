import { getActivePinia, type Pinia, type Store, type StoreDefinition } from 'pinia';

export function destroyStore(store: Store, pinia: Pinia = getActivePinia()!): void {
  store.$dispose();
  delete pinia.state.value[store.$id];
}

export function lazyDestroyStore(definition: StoreDefinition, pinia: Pinia = getActivePinia()!): void {
  if (pinia.state.value[definition.$id]) {
    destroyStore(definition(), pinia);
  }
}
