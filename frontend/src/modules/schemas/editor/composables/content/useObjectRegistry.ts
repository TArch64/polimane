import type { InjectionKey } from 'vue';
import type { FabricObject } from 'fabric';
import { injectLocal, provideLocal } from '@vueuse/core';
import { EditorObjectType } from '../../enums';
import { injectCanvas } from '../useCanvas';
import { type EditorObjectTypeMap, isDestroyableObject } from './objects';

const TOKEN = Symbol('ObjectRegistry') as InjectionKey<Map<string, FabricObject>>;

export interface IObjectTypeRegistry<T extends EditorObjectType> {
  get: (id: string) => EditorObjectTypeMap[T];
  add: (id: string, object: EditorObjectTypeMap[T]) => void;
  remove: (id: string) => void;
}

export type ObjectRegistry = {
  [T in EditorObjectType]: IObjectTypeRegistry<T>;
};

export function initObjectRegistry(): void {
  provideLocal(TOKEN, new Map());
}

export function useObjectRegistry<T extends EditorObjectType>(type: T): IObjectTypeRegistry<T> {
  const registry = injectLocal(TOKEN)!;
  const canvas = injectCanvas();

  function get<T extends EditorObjectType>(type: T, id: string): EditorObjectTypeMap[T] {
    return registry.get(`${type}:${id}`) as EditorObjectTypeMap[T];
  }

  function add<T extends EditorObjectType>(type: T, id: string, object: EditorObjectTypeMap[T]): void {
    registry.set(`${type}:${id}`, object);
    canvas.add(object);
  }

  function remove(type: EditorObjectType, id: string): void {
    const object = get(type, id);
    registry.delete(`${type}:${id}`);
    canvas.remove(object);

    if (isDestroyableObject(object)) {
      object.destroy();
    }
  }

  return {
    get: (id) => get(type, id),
    add: (id, object) => add(type, id, object),
    remove: (id) => remove(type, id),
  };
}
