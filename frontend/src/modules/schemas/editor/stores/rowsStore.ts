import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaBead, type ISchemaPattern, type ISchemaRow } from '@/models';
import { newArray, newId } from '@/helpers';
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
      content: newArray(size, createBead),
    });

    function addSquareRow(options: INewSquareRowOptions) {
      const newRows = newArray(options.rows, () => createRow(options.size));

      rows.insert(newRows, {
        toIndex: options.toIndex,
      });
    }

    function deleteRow(row: ISchemaRow): void {
      rows.delete(row);
    }

    function moveRow(row: ISchemaRow, shift: number): void {
      const index = rows.indexOf(row);
      const newIndex = index + shift;

      if (newIndex < 0 || newIndex === rows.size) {
        return;
      }

      rows.move(row, newIndex);
    }

    function resizeRow(row: ISchemaRow, newSize: number): void {
      const newBeadsCount = newSize - row.content.length;
      const newBeads = newArray(newBeadsCount, createBead);
      rows.update(row, { content: [...row.content, ...newBeads] });
    }

    return { rows, addSquareRow, deleteRow, moveRow, resizeRow };
  },
});

export const {
  useStore: useRowsStore,
  disposeStores: disposeRowsStores,
} = rowsDynamicStore.build();
