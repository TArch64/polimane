import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaBead, type ISchemaPattern, type ISchemaRow } from '@/models';
import { newId } from '@/helpers';
import { DynamicStore } from '@/stores';
import { setObjectParent } from '../models';

export interface INewSquareRowOptions {
  rows: number;
  size: number;
  toIndex: number;
}

const rowsDynamicStore = new DynamicStore({
  buildPath(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
    const { id } = toValue(patternRef);
    return `schemas/editor/patterns/${id}/rows` as const;
  },

  setup(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
    const pattern = toValue(patternRef);

    const rows = Collection.fromParent(pattern, {
      onAdded: (parent, object) => setObjectParent(parent, object),
    });

    const createBead = (): ISchemaBead => ({
      id: newId(),
      color: '',
    });

    const createRow = (size: number): ISchemaRow => ({
      id: newId(),
      content: new Array(size).fill(0).map(createBead),
    });

    function addSquareRow(options: INewSquareRowOptions) {
      const newRows = new Array(options.rows)
        .fill(0)
        .map(() => createRow(options.size));

      rows.insert(newRows, {
        toIndex: options.toIndex,
      });
    }

    function deleteRow(row: ISchemaRow): void {
      rows.delete(row);
    }

    return { rows, addSquareRow, deleteRow };
  },
});

export const {
  useStore: useRowsStore,
  disposeStores: disposeRowsStores,
} = rowsDynamicStore.build();
