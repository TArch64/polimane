import { computed, type MaybeRefOrGetter, reactive, type Ref, toValue } from 'vue';
import type { ISchemaObject } from '@/models';
import {
  type ActiveObjectStore,
  ActiveObjectTrigger,
  useFocusObjectStore,
  useHoverObjectStore,
} from '../stores';

export interface IActiveObjectState {
  isActive: boolean;
  fromSidebar: boolean;
  fromCanvas: boolean;
  activate: (trigger: ActiveObjectTrigger) => void;
  deactivate: () => void;
}

export interface IActiveObject {
  hover: IActiveObjectState;
  focus: IActiveObjectState;
}

function useActiveObjectState(store: ActiveObjectStore, object: Ref<ISchemaObject>): IActiveObjectState {
  const isActive = computed(() => store.isActiveObject(object.value));
  const fromSidebar = computed(() => isActive.value && store.activePathTrigger === ActiveObjectTrigger.SIDEBAR);
  const fromCanvas = computed(() => isActive.value && store.activePathTrigger === ActiveObjectTrigger.CANVAS);
  const activate = (trigger: ActiveObjectTrigger) => store.activateObject(object.value, trigger);
  const deactivate = () => store.deactivatePath();

  return reactive({
    isActive,
    fromSidebar,
    fromCanvas,
    activate,
    deactivate,
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
