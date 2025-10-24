<template>
  <Teleport
    :to="openedConfirm.getTopEl()"
    :key="openedConfirm.id"
    v-if="openedConfirm"
  >
    <Confirm :model="openedConfirm" />

    <VirtualTarget
      :model="openedConfirm"
      v-if="openedConfirm.virtualTarget"
    />
  </Teleport>
</template>

<script setup lang="ts">
import { type FunctionalComponent, h, nextTick, type Ref, ref, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import { ConfirmPlugin } from './ConfirmPlugin';
import type { Confirm as ConfirmModel } from './Confirm';
import Confirm from './Confirm.vue';

const plugin = ConfirmPlugin.inject();
const routeTransition = useRouteTransition();
const openedConfirm: Ref<ConfirmModel | null> = ref(null);

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

VirtualTarget.displayName = 'VirtualTarget';
</script>
