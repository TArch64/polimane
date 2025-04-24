import { defineStore } from 'pinia';
import type { ISchemaPattern } from '@/models';
import { newId } from '@/helpers';

export interface INewSquareRowOptions {
  size: number;
}

export const useRowsStore = defineStore('schemas/editor/rows', () => {
  function addSquareRow(pattern: ISchemaPattern, options: INewSquareRowOptions): void {
    pattern.content.push({
      id: newId(),

      content: new Array(options.size).fill(0).map(() => ({
        id: newId(),
      })),
    });
  }

  return { addSquareRow };
});
