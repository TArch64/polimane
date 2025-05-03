import { computed, type MaybeRefOrGetter, reactive, type Ref, toValue } from 'vue';
import { createEventHook, type EventHookOn, whenever } from '@vueuse/core';
import type { ISchemaObject } from '@/models';
import {
  type ActiveObjectStore,
  ActiveObjectTrigger,
  useFocusObjectStore,
  useHoverObjectStore,
} from '../stores';

export interface IActiveObjectState {
  isActive: boolean;
  isExactActive: boolean;
  onExactActive: EventHookOn<ActiveObjectTrigger>;
  activate: (trigger: ActiveObjectTrigger) => void;
  toggle: (trigger: ActiveObjectTrigger) => void;
  deactivate: () => void;
}

export interface IActiveObject {
  hover: IActiveObjectState;
  focus: IActiveObjectState;
}

function useActiveObjectState(store: ActiveObjectStore, object: Ref<ISchemaObject>): IActiveObjectState {
  const isActive = computed(() => store.isActiveObject(object.value));
  const isExactActive = computed(() => store.isExactActiveObject(object.value));

  const activate = (trigger: ActiveObjectTrigger) => store.activateObject(object.value, trigger);
  const deactivate = () => store.deactivatePath();

  function toggle(trigger: ActiveObjectTrigger) {
    isExactActive.value ? store.popPath() : activate(trigger);
  }

  const exactActiveHook = createEventHook<ActiveObjectTrigger>();
  whenever(isExactActive, () => exactActiveHook.trigger(store.activePathTrigger!));

  return reactive({
    isActive,
    isExactActive,
    onExactActive: exactActiveHook.on,
    activate,
    deactivate,
    toggle,
  });
}

export function useActiveObject(object: MaybeRefOrGetter<ISchemaObject>): IActiveObject {
  const hoverStore = useHoverObjectStore();
  const focusStore = useFocusObjectStore();
  const objectRef = computed(() => toValue(object));

  return {
    hover: useActiveObjectState(hoverStore, objectRef),
    focus: useActiveObjectState(focusStore, objectRef),
  };
}
