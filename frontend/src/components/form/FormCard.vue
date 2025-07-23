<template>
  <Card :title :as="Form" class="form-card">
    <slot />

    <template #footer>
      <div class="form-card__spacer" />

      <Button @click="$emit('reset')" v-if="hasChanges">
        Відмінити
      </Button>

      <Button type="submit" variant="primary">
        Зберегти
      </Button>
    </template>
  </Card>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Card } from '../card';
import { Button } from '../button';
import Form from './Form.vue';

defineProps<{
  title?: string;
  hasChanges?: boolean;
}>();

defineEmits<{
  submit: [];
  reset: [];
}>();

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .form-card {
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 12px;
  }

  .form-card__spacer {
    margin-left: auto;
  }
}
</style>
