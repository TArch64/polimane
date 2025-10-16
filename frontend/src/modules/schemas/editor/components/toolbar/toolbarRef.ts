import { inject, type InjectionKey, provide, type Ref } from 'vue';

const TOOLBAR_REF = Symbol('toolbarRef') as InjectionKey<Ref<HTMLElement>>;

export function provideToolbarRef(toolbarRef: Ref<HTMLElement>) {
  provide(TOOLBAR_REF, toolbarRef);
}

export function useToolbarRef(): Ref<HTMLElement> {
  return inject(TOOLBAR_REF)!;
}
