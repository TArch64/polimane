import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaRow } from '@/models';
import { StoreFactory } from '@/stores';
import { setObjectParent } from '../models';

const beadsStoreFactory = new StoreFactory({
  buildPath(rowRef: MaybeRefOrGetter<ISchemaRow>) {
    const { id } = toValue(rowRef);
    return `schemas/editor/rows/${id}/beads`;
  },

  setup(rowRef: MaybeRefOrGetter<ISchemaRow>) {
    const row = toValue(rowRef);

    const beads = Collection.fromParent(row, {
      onAdded: (parent, object) => setObjectParent(parent, object),
    });

    return { beads };
  },
});

export const {
  useStore: useBeadsStore,
  disposeStores: disposeBeadsStores,
} = beadsStoreFactory.build();
