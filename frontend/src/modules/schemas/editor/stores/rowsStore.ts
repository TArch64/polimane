import { defineStore } from 'pinia';
import { type MaybeRefOrGetter, toValue } from 'vue';
import { Collection, type ISchemaPattern, type ISchemaRow } from '@/models';
import { newId } from '@/helpers';
import { setObjectParent } from '../models';

export interface INewSquareRowOptions {
  size: number;
}

export function useRowsStore(patternRef: MaybeRefOrGetter<ISchemaPattern>) {
  const pattern = toValue(patternRef);

  return defineStore(`schemas/editor/rows/${pattern.id}`, () => {
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
  })();
}
