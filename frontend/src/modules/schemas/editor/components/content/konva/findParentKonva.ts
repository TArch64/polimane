import type { ComponentInternalInstance } from 'vue';

export function findParentKonva(instance: ComponentInternalInstance) {
  function re(instance: ComponentInternalInstance | null): ComponentInternalInstance | null {
    if (instance?.__konvaNode) {
      return instance;
    }
    if (instance?.parent) {
      return re(instance.parent);
    }
    console.error('vue-konva error: Can not find parent node');
    return null;
  }

  return re(instance.parent);
}
