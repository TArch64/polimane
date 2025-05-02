import { computed, type MaybeRefOrGetter, ref, toValue } from 'vue';
import { defineStore } from 'pinia';
import { type ISchemaObject, isSchemaWithContent } from '@/models';

const OBJECT_PARENT = Symbol('[[OBJECT_PARENT]]');

export function setObjectParent(parent: ISchemaObject, object: ISchemaObject): void {
  Object.defineProperty(object, OBJECT_PARENT, {
    value: parent,
    enumerable: false,
    configurable: false,
    writable: false,
  });

  if (isSchemaWithContent(object)) {
    for (const child of object.content) {
      setObjectParent(object, child);
    }
  }
}

export function getObjectParent(object: ISchemaObject): ISchemaObject | undefined {
  return Object.getOwnPropertyDescriptor(object, OBJECT_PARENT)?.value;
}

export function getObjectPath(object: ISchemaObject): string[] {
  const path: string[] = [object.id];
  let parent = getObjectParent(object);

  while (parent) {
    path.unshift(parent.id);
    parent = getObjectParent(parent);
  }

  return path;
}

function activeObjectStore() {
  const activePath = ref<string[]>([]);
  const activatePath = (path: string[]) => activePath.value = path;
  const activateObject = (object: ISchemaObject) => activatePath(getObjectPath(object));
  const deactivatePath = () => activatePath([]);
  const isActiveObject = (object: ISchemaObject) => activePath.value.some((id) => id === object.id);
  const useActiveObject = (object: MaybeRefOrGetter<ISchemaObject>) => computed(() => isActiveObject(toValue(object)));

  return {
    activePath,
    activateObject,
    isActiveObject,
    useActiveObject,
    deactivatePath,
  };
}

export const useFocusObjectStore = defineStore('schemas/editor/focusObject', activeObjectStore);
export const useHoverObjectStore = defineStore('schemas/editor/hoverObject', activeObjectStore);
