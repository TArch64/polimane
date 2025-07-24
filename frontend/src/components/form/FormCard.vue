<template>
  <Card :title :binding :footerTransition class="form-card">
    <slot />

    <template #footer v-if="hasChanges">
      <div class="form-card__spacer" />

      <Button @click="$emit('reset')">
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
import { makeBinding } from '@/components/binding';
import { Card, type ICardFooterTransition } from '../card';
import { Button } from '../button';
import Form from './Form.vue';

defineProps<{
  title?: string;
  hasChanges: boolean;
}>();

const emit = defineEmits<{
  submit: [];
  reset: [];
}>();

defineSlots<{
  default: Slot;
}>();

const binding = makeBinding(Form, () => ({
  onSubmit: () => emit('submit'),
}));

const footerTransition: Partial<ICardFooterTransition> = {
  // duration: 2000,
  shift: -8,
};
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
