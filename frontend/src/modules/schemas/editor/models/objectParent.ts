import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import {
  type ISchema,
  type ISchemaObject,
  type ISchemaWithContent,
  isSchemaWithContent,
} from '@/models';

const OBJECT_PARENT = Symbol('[[OBJECT_PARENT]]');

export function setSchemaRelations(schema: ISchema): void {
  for (const object of schema.content) {
    setObjectParent(schema, object);
  }
}

export function setObjectParent(parent: ISchemaObject, object: ISchemaObject): void {
  if (!Object.getOwnPropertyDescriptor(object, OBJECT_PARENT)) {
    Object.defineProperty(object, OBJECT_PARENT, {
      value: parent,
      enumerable: false,
      configurable: false,
      writable: false,
    });
  }

  if (isSchemaWithContent(object)) {
    for (const child of object.content) {
      setObjectParent(object, child);
    }
  }
}

export function getObjectParent<P extends ISchemaWithContent>(object: P['content'][number]): P | undefined {
  return Object.getOwnPropertyDescriptor(object, OBJECT_PARENT)?.value;
}

export function useObjectParent<P extends ISchemaWithContent>(objectRef: MaybeRefOrGetter<P['content'][number]>): ComputedRef<P> {
  return computed(() => getObjectParent(toValue(objectRef))!);
}

export function getObjectPath(object: ISchemaObject): string[] {
  const path: string[] = [object.id as string];
  let parent = getObjectParent(object);

  while (parent) {
    path.unshift(parent.id as string);
    parent = getObjectParent(parent);
  }

  return path;
}
