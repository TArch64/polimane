import { createGlobalState } from '@vueuse/core';
import { ref } from 'vue';

export const useAuthorized = createGlobalState(() => ref(false));
