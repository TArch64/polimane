<template>
  <ColorItem
    :color="model"
    :active="isActive"
    @click.stop="onClick"
    @dblclick.stop="onDblClick"
  >
    <input
      ref="pickerRef"
      type="color"
      class="color-palette__spot-picker"
      v-model="model"
    >
  </ColorItem>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import ColorItem from './ColorItem.vue';

const emit = defineEmits<{
  choose: [];
}>();

const isActive = defineModel<boolean>('active', {
  required: true,
});

const model = defineModel<string>({
  required: true,

  set: (value) => {
    isActive.value = true;
    return value;
  },
});

const pickerRef = ref<HTMLInputElement>(null!);

function onClick(): void {
  if (!model.value) {
    pickerRef.value.click();
    return;
  }

  isActive.value = true;
  emit('choose');
}

function onDblClick(): void {
  if (model.value) pickerRef.value.click();
}
</script>

<style scoped>
@layer page {
  .color-palette__spot-picker {
    visibility: hidden;
    width: 100%;
    height: 100%;
  }
}
</style>
