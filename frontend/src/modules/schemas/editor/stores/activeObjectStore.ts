import { ref } from 'vue';
import type { ISchemaObject } from '@/models';
import { StoreFactory } from '@/stores';
import { getObjectPath } from '../models';

export const enum ActiveObjectTrigger {
  SIDEBAR = 'sidebar',
  CANVAS = 'canvas',
}

const activeObjectStoreFactory = new StoreFactory({
  buildPath(type: 'hover' | 'focus') {
    return `schemas/editor/object/${type}` as const;
  },

  setup() {
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

    function deactivateObject(object: ISchemaObject) {
      const index = activePath.value.findIndex((id) => id === object.id);
      if (index !== -1) activePath.value = activePath.value.slice(0, index);
    }

    const isActiveObject = (object: ISchemaObject) => activePath.value.some((id) => id === object.id);
    const isExactActiveObject = (object: ISchemaObject) => activePath.value.at(-1) === object.id;

    return {
      activePath,
      activePathTrigger,
      activateObject,
      deactivatePath,
      deactivateObject,
      isActiveObject,
      isExactActiveObject,
    };
  },
});

const { useStore } = activeObjectStoreFactory.build();
export type ActiveObjectStore = typeof activeObjectStoreFactory['$storeType'];
export const useFocusObjectStore = () => useStore('focus');
export const useHoverObjectStore = () => useStore('hover');
