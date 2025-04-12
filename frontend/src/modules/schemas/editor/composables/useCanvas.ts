import { injectLocal, provideLocal } from '@vueuse/core';
import type { InjectionKey, Ref } from 'vue';
import type { Canvas } from 'fabric';

const TOKEN = Symbol('canvas') as InjectionKey<Ref<Canvas>>;

export const provideCanvas = (canvas: Ref<Canvas>) => provideLocal(TOKEN, canvas);
export const injectCanvas = () => injectLocal(TOKEN)!;
