import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { reactiveComputed } from '@vueuse/core';
import { AccessLevel } from '@/enums';

export interface IAccessPermissions {
  read: boolean;
  write: boolean;
  admin: boolean;
}

export function useAccessPermissions(accessRef: MaybeRefOrGetter<AccessLevel>): IAccessPermissions {
  const access = computed(() => toValue(accessRef));

  return reactiveComputed(() => ({
    read: access.value >= AccessLevel.READ,
    write: access.value >= AccessLevel.WRITE,
    admin: access.value === AccessLevel.ADMIN,
  }));
}
