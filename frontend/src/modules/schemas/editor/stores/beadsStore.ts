import { defineStore } from 'pinia';
import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaRow } from '@/models';
import { setObjectParent } from '../models';

export function useBeadsStore(rowRef: MaybeRefOrGetter<ISchemaRow>) {
  const row = toValue(rowRef);

  return defineStore(`schemas/editor/rows/${row.id}/beads`, () => {
    const beads = Collection.fromParent(row, {
      onAdded: (parent, object) => setObjectParent(parent, object),
    });

    return { beads };
  })();
}
