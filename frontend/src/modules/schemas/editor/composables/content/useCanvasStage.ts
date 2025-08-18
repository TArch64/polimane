import Konva from 'konva';
import { type InjectionKey, onUnmounted, type ShallowRef, watch } from 'vue';
import { injectLocal, provideLocal } from '@vueuse/core';

const PROVIDER = Symbol('CanvasStage') as InjectionKey<ShallowRef<Konva.Stage>>;

export function provideCanvasStage(stageRef: ShallowRef<Konva.Stage>): void {
  provideLocal(PROVIDER, stageRef);

  if (import.meta.env.DEV) {
    watch(stageRef, (stage) => {
      // @ts-expect-error This is a global variable for debugging purposes
      window.$konvaStage = stage;
    }, { immediate: true });

    onUnmounted(() => {
      // @ts-expect-error This is a global variable for debugging purposes
      window.$konvaStage = undefined;
    });
  }
}

export function useCanvasStage() {
  return injectLocal(PROVIDER)!;
}
