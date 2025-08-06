import { type MaybeRefOrGetter, toValue } from 'vue';
import {
  Collection,
  type ISchemaBead,
  type ISchemaPattern,
  type ISchemaSquareRow,
  type SchemaRow,
} from '@/models';
import { newArray, newId } from '@/helpers';
import { DynamicStore } from '@/stores';
import { PatternType } from '@/enums';

interface INewRowOptions {
  rows: number;
  type: PatternType;
  toIndex: number;
}

export interface INewSquareRowOptions extends INewRowOptions {
  size: number;
}

export interface INewDiamondRowOptions extends INewRowOptions {
  size: number;
  sideSize: number;
}

const rowsDynamicStore = new DynamicStore({
  buildPath(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
    const { id } = toValue(patternRef);
    return `schemas/editor/patterns/${id}/rows` as const;
  },

  setup(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
    const pattern = toValue(patternRef);
    const rows = Collection.fromParent(pattern);

    const createBead = (): ISchemaBead => ({
      id: newId(),
      color: '',
    });

    function addRows(newRows: SchemaRow[], toIndex: number): void {
      rows.insert(newRows, { toIndex });
    }

    function addSquareRow(options: INewSquareRowOptions) {
      const newRows = newArray<ISchemaSquareRow>(options.rows, () => ({
        id: newId(),
        content: newArray(options.size, createBead),
        square: { size: options.size },
      }));

      addRows(newRows, options.toIndex);
    }

    function addDiamondRow(options: INewDiamondRowOptions) {
    }

    function deleteRow(row: SchemaRow): void {
      rows.delete(row);
    }

    function moveRow(row: SchemaRow, shift: number): void {
      const index = rows.indexOf(row);
      const newIndex = index + shift;

      if (newIndex < 0 || newIndex === rows.size) {
        return;
      }

      rows.move(row, newIndex);
    }

    function resizeRow(row: SchemaRow, newSize: number): void {
      const newBeadsCount = newSize - row.content.length;
      const newBeads = newArray(newBeadsCount, createBead);
      rows.update(row, { content: [...row.content, ...newBeads] });
    }

    return { rows, addSquareRow, addDiamondRow, deleteRow, moveRow, resizeRow };
  },
});

export const {
  useStore: useRowsStore,
  disposeStores: disposeRowsStores,
} = rowsDynamicStore.build();
