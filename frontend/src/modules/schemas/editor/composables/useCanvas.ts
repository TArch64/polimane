import { injectLocal, provideLocal, reactiveComputed } from '@vueuse/core';
import type { InjectionKey, Ref } from 'vue';
import type { Canvas } from 'fabric';

const TOKEN = Symbol('canvas') as InjectionKey<Ref<Canvas | null>>;

export const provideCanvas = (canvas: Ref<Canvas | null>) => provideLocal(TOKEN, canvas);
export const injectCanvasRef = () => injectLocal(TOKEN)!;

export function injectCanvas() {
  const canvas = injectCanvasRef();
  return reactiveComputed(() => canvas.value ?? {}) as Canvas;
}
