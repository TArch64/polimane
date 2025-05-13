import { inject, type InjectionKey, provide, shallowRef, type ShallowRef } from 'vue';
import Konva from 'konva';

const TOKEN = Symbol('CANVAS_STAGE') as InjectionKey<ShallowRef<Konva.Stage | null>>;

export const useMaybeCanvasStage = () => inject(TOKEN)!;
export const useCanvasStage = () => inject(TOKEN) as ShallowRef<Konva.Stage>;
export const provideCanvasStage = () => provide(TOKEN, shallowRef(null));
