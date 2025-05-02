import { defineStore } from 'pinia';
import { Collection, type ISchemaPattern } from '@/models';
import { newId } from '@/helpers';
import { setObjectParent } from '../models';

export interface INewSquareRowOptions {
  size: number;
}

export const useRowsStore = defineStore('schemas/editor/rows', () => {
  function addSquareRow(pattern: ISchemaPattern, options: INewSquareRowOptions): void {
    const collection = Collection.fromParent(pattern, {
      onAdded: (parent, object) => setObjectParent(parent, object),
    });

    collection.append({
      id: newId(),

      content: new Array(options.size).fill(0).map(() => ({
        id: newId(),
      })),
    });
  }

  return { addSquareRow };
});
