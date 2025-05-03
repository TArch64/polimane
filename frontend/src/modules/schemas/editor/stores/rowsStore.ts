import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaPattern, type ISchemaRow } from '@/models';
import { newId } from '@/helpers';
import { StoreFactory } from '@/stores';
import { setObjectParent } from '../models';

export interface INewSquareRowOptions {
  size: number;
}

const rowsStoreFactory = new StoreFactory({
  buildPath(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
    const { id } = toValue(patternRef);
    return `schemas/editor/patterns/${id}/rows` as const;
  },

  setup(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
    const pattern = toValue(patternRef);

    const rows = Collection.fromParent(pattern, {
      onAdded: (parent, object) => setObjectParent(parent, object),
    });

    const addSquareRow = (options: INewSquareRowOptions) => rows.append({
      id: newId(),

      content: new Array(options.size).fill(0).map(() => ({
        id: newId(),
      })),
    });

    function deleteRow(row: ISchemaRow): void {
      rows.delete(row);
    }

    return { rows, addSquareRow, deleteRow };
  },
});

export const {
  useStore: useRowsStore,
  disposeStores: disposeRowsStores,
} = rowsStoreFactory.build();
