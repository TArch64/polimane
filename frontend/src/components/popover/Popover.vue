<template>
  <slot
    name="activator"
    :open="state.open"
    :isOpened="state.isOpened"
    :ref="setActivatorRef"
  />

  <Teleport to="body" v-if="state.isOpened || state.transition.isActive">
    <FadeTransition appear :state="state.transition">
      <Card
        as="dialog"
        ref="dialogRef"
        class="popover"
        popover="manual"
        :style="popoverStyles"
        v-popover-shift.defer
        v-if="state.isOpened"
      >
        <slot />

        <PopoverTip
          :reference-el="activatorRef"
          :floating-el="dialogRef"
          v-if="dialogRef && activatorRef"
        />
      </Card>
    </FadeTransition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, ref, type Slot, toRef, useId, type VNodeRef } from 'vue';
import { unrefElement } from '@vueuse/core';
import { useDomRef, usePopoverState } from '@/composables';
import { vPopoverShift } from '@/directives';
import type { NumericString } from '@/types';
import { FadeTransition } from '../transition';
import { Card } from '../card';
import PopoverTip from './PopoverTip.vue';
import { POPOVER_TIP_HEIGHT } from './config';

const props = withDefaults(defineProps<{
  positionArea: string;
  positionOffsetBlock?: number | NumericString;
}>(), {
  positionOffsetBlock: 0,
});

defineSlots<{
  activator: Slot<{
    ref: VNodeRef;
    isOpened: boolean;
    open: () => void;
  }>;
  default: Slot;
}>();

const dialogRef = useDomRef<HTMLDialogElement | null>();
const state = usePopoverState(dialogRef);

const activatorRef = ref<HTMLElement | null>(null);

const setActivatorRef: VNodeRef = (ref) => {
  activatorRef.value = ref ? unrefElement(ref as HTMLElement)! : null;

  if (activatorRef.value) {
    const style = activatorRef.value.style.getPropertyValue('anchor-name');
    const anchors = style ? style.split(/, ?/) : [];
    const set = new Set(anchors).add(anchorName);
    activatorRef.value.style.setProperty('anchor-name', [...set].join(', '));
  }
};

const anchorName = `--popover-${useId()}`;

const marginBlockEnd = computed(() => {
  const offset = Number(props.positionOffsetBlock);
  return `${offset + POPOVER_TIP_HEIGHT}px`;
});

const popoverStyles = computed(() => ({
  'position-anchor': anchorName,
  'position-area': props.positionArea,
  'margin-block-end': marginBlockEnd.value,
}));

defineExpose({
  open: state.open,
  close: state.close,
  isOpened: toRef(state, 'isOpened'),
});
</script>

<style scoped>
@layer components {
  .popover {
    outline: none;
    overflow: initial;
  }
}
</style>
