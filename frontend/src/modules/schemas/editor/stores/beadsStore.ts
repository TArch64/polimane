import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaBead, type ISchemaRow } from '@/models';
import { DynamicStore } from '@/stores';
import { usePaletteStore } from './paletteStore';

const beadsDynamicStore = new DynamicStore({
  buildPath(rowRef: MaybeRefOrGetter<ISchemaRow>) {
    const { id } = toValue(rowRef);
    return `schemas/editor/rows/${id}/beads`;
  },

  setup(rowRef: MaybeRefOrGetter<ISchemaRow>) {
    const row = toValue(rowRef);
    const paletteStore = usePaletteStore();
    const beads = Collection.fromParent(row);

    function paint(bead: ISchemaBead): void {
      bead.color = paletteStore.activeColor;
    }

    return { beads, paint };
  },
});

export const {
  useStore: useBeadsStore,
  disposeStores: disposeBeadsStores,
} = beadsDynamicStore.build();
