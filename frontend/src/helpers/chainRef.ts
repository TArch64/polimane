import { isRef, type VNodeRef } from 'vue';

export function chainRef(...targetRefs: VNodeRef[]): VNodeRef {
  return (ref, refs) => {
    for (const targetRef of targetRefs) {
      if (typeof targetRef === 'function') {
        targetRef(ref, refs);
      } else if (isRef(targetRef)) {
        targetRef.value = ref;
      }
    }
  };
}
