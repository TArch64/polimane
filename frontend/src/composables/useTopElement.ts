import { inject, type InjectionKey, provide, type Ref } from 'vue';

const TOP_ELEMENT_PROVIDER = Symbol('topElementProvider') as InjectionKey<Ref<HTMLElement>>;

export function useTopElement(): Ref<HTMLElement> {
  return inject(TOP_ELEMENT_PROVIDER)!;
}

export function provideTopElement(element: Ref<HTMLElement>): void {
  provide(TOP_ELEMENT_PROVIDER, element);
}
