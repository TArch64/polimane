<template>
  <Teleport
    :to="openedConfirmTopEl"
    :key="openedConfirm.id"
    v-if="openedConfirm && openedConfirmTopEl"
  >
    <Confirm :model="openedConfirm" />

    <VirtualTarget
      :model="openedConfirm"
      v-if="openedConfirm.virtualTarget"
    />
  </Teleport>
</template>

<script setup lang="ts">
import {
  computed,
  type FunctionalComponent,
  h,
  nextTick,
  type Ref,
  ref,
  toValue,
  watch,
} from 'vue';
import { useRouteTransition } from '@/composables';
import { ConfirmPlugin } from './ConfirmPlugin';
import type { Confirm as ConfirmModel } from './Confirm';
import Confirm from './Confirm.vue';

const plugin = ConfirmPlugin.inject();
const routeTransition = useRouteTransition();

const openedConfirm: Ref<ConfirmModel | null> = ref(null);
const openedConfirmTopEl = computed(() => toValue(openedConfirm.value?.topEl));

watch(() => plugin.openedConfirm?.id, async () => {
  routeTransition.start(async () => {
    openedConfirm.value = plugin.openedConfirm ?? null;
    await nextTick();
  });
});

const VirtualTarget: FunctionalComponent<{ model: ConfirmModel }> = (props) => h('div', {
  inert: true,

  style: {
    position: 'fixed',
    anchorName: props.model.anchorVar,
    top: `${props.model.virtualTarget!.top}px`,
    left: `${props.model.virtualTarget!.left}px`,
    width: `${props.model.virtualTarget!.width}px`,
  },
});
</script>
