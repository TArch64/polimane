import Konva from 'konva';
import type { InjectionKey, ShallowRef } from 'vue';
import { injectLocal, provideLocal } from '@vueuse/core';

const PROVIDER = Symbol('CanvasStage') as InjectionKey<ShallowRef<Konva.Stage>>;

export function provideCanvasStage(stage: ShallowRef<Konva.Stage>): void {
  provideLocal(PROVIDER, stage);
}

export function useCanvasStage() {
  return injectLocal(PROVIDER)!;
}
