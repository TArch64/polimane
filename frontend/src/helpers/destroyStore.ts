import { getActivePinia, type Pinia, type Store } from 'pinia';

export function destroyStore(store: Store, pinia: Pinia = getActivePinia()!): void {
  store.$dispose();
  delete pinia.state.value[store.$id];
}
