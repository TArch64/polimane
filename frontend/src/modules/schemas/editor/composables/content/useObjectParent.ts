import type { Canvas, Group } from 'fabric';
import { inject, type InjectionKey, provide } from 'vue';

export type ObjectParent = Canvas | Group;

const TOKEN = Symbol('ObjectParent') as InjectionKey<ObjectParent>;

export const useObjectParent = () => inject(TOKEN)!;
export const provideObjectParent = (object: ObjectParent) => provide(TOKEN, object);
