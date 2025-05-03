import { ref } from 'vue';
import {
  type _ExtractActionsFromSetupStore,
  type _ExtractGettersFromSetupStore,
  type _ExtractStateFromSetupStore,
  defineStore,
  type Store,
} from 'pinia';
import type { ISchemaObject } from '@/models';
import { getObjectPath } from '../models';

export const enum ActiveObjectTrigger {
  SIDEBAR = 'sidebar',
  CANVAS = 'canvas',
}

function activeObjectStore() {
  const activePath = ref<string[]>([]);
  const activePathTrigger = ref<ActiveObjectTrigger | null>(null);

  function activatePath(path: string[], trigger: ActiveObjectTrigger | null) {
    activePath.value = path;
    activePathTrigger.value = trigger;
  }

  function activateObject(object: ISchemaObject, trigger: ActiveObjectTrigger) {
    activatePath(getObjectPath(object), trigger);
  }

  const deactivatePath = () => activatePath([], null);
  const popPath = () => activePath.value.pop();

  const isActiveObject = (object: ISchemaObject) => activePath.value.some((id) => id === object.id);
  const isExactActiveObject = (object: ISchemaObject) => activePath.value.at(-1) === object.id;

  return {
    activePath,
    activePathTrigger,
    activateObject,
    deactivatePath,
    popPath,
    isActiveObject,
    isExactActiveObject,
  };
}

type StoreReturn = ReturnType<typeof activeObjectStore>;

export type ActiveObjectStore = Store<
  string,
  _ExtractStateFromSetupStore<StoreReturn>,
  _ExtractGettersFromSetupStore<StoreReturn>,
  _ExtractActionsFromSetupStore<StoreReturn>
>;

export const useFocusObjectStore = defineStore('schemas/editor/focusObject', activeObjectStore);
export const useHoverObjectStore = defineStore('schemas/editor/hoverObject', activeObjectStore);
