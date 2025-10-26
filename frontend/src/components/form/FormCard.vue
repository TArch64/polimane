<template>
  <Card :title :binding :footerTransition class="form-card">
    <slot />

    <template #footer v-if="hasChanges || submitPersistent">
      <div class="form-card__spacer" />

      <Button
        class="vertical-slice-transition__item"
        :disabled="loading"
        @click="$emit('reset')"
        v-if="cancelable"
      >
        Відмінити
      </Button>

      <Button
        :loading
        type="submit"
        variant="primary"
        class="vertical-slice-transition__item"
      >
        {{ submitText }}
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

withDefaults(defineProps<{
  title?: string;
  hasChanges: boolean;
  submitText?: string;
  submitPersistent?: boolean;
  cancelable?: boolean;
  loading?: boolean;
}>(), {
  title: '',
  submitText: 'Зберегти',
  submitPersistent: false,
  cancelable: true,
});

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
  duration: 250,
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
