<template>
  <div class="sidebar-structure-empty">
    <Button
      variant="secondary"
      size="md"
      class="sidebar-structure-empty__button"
      @click="$emit('click')"
    >
      <PlusIcon class="sidebar-structure-empty__button-icon" />
      {{ buttonText }}
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { PlusIcon } from '@/components/icon';

defineProps<{
  buttonText: string;
}>();

defineEmits<{
  click: [];
}>();
</script>

<style scoped>
@layer page {
  .sidebar-structure-empty {
    padding: 12px 20px;
    z-index: 2;
    position: relative;
    background-color: var(--color-background-2);
  }

  .sidebar-structure-empty__button {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    padding-right: 20px;
  }

  .sidebar-structure-empty__button-icon {
    flex-shrink: 0;
    margin-right: 8px;
  }
}
</style>
