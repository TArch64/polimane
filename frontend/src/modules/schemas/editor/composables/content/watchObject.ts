import { computed, watch } from 'vue';
import type { ISchemaObject } from '@/models';

export type WatchableObject<O extends ISchemaObject> = Omit<O, 'content'>;
export type WatchObjectSource<O extends ISchemaObject> = () => O;
export type WatchObjectCallback<O extends ISchemaObject> = (object: WatchableObject<O>) => void;

export function watchObject<O extends ISchemaObject>(source: WatchObjectSource<O>, onChange: WatchObjectCallback<O>): void {
  const watchable = computed(() => {
    const object = source();

    if ('content' in object) {
      const { content: _, ...objectSelf } = object;
      return objectSelf;
    }

    return object;
  });

  watch(watchable, (object) => onChange(object));
}
