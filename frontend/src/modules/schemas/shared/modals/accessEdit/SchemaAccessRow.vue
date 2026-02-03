<template>
  <li class="access-row">
    <p class="text-truncate access-row__name">
      <span
        class="access-row__name-text"
        :class="nameClasses"
        :title="nameHint"
      >
        {{ displayName }}
      </span>

      <span v-if="!editable" class="access-row__name-label">
        (Ви)
      </span>
    </p>

    <div class="access-row__actions" v-if="editable">
      <SchemaAccessField v-model="access" />

      <Button
        icon
        danger
        size="md"
        variant="secondary"
        :style="deleteConfirm.anchorStyle"
        @click="deleteUserIntent"
      >
        <TrashIcon />
      </Button>
    </div>
  </li>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { AccessLevel } from '@/enums';
import SchemaAccessField from './SchemaAccessField.vue';

const props = withDefaults(defineProps<{
  displayName: string;
  evenAccess: boolean;
  allAccess: boolean;
  editable?: boolean;
}>(), {
  editable: true,
});

const access = defineModel<AccessLevel>('access', { required: true });

const emit = defineEmits<{
  delete: [];
}>();

const hasIssue = computed(() => {
  return !props.evenAccess || !props.allAccess;
});

const nameClasses = computed(() => ({
  'access-row__name-text--issue': hasIssue.value,
}));

const nameHint = computed(() => {
  if (!props.evenAccess) {
    return 'Має не однаковий доступ до схем';
  }
  if (!props.allAccess) {
    return 'Має доступ не до всіх схем';
  }
  return '';
});

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете заборонити доступ до схему цьому користувачу?',
  acceptButton: 'Заборонити',
});

async function deleteUserIntent(): Promise<void> {
  const confirmation = await deleteConfirm.ask();
  if (confirmation.isAccepted) emit('delete');
}
</script>

<style scoped>
@layer page {
  .access-row {
    display: flex;
    align-items: center;
    min-height: 32px;
    gap: 8px;
  }

  .access-row__name {
    flex-basis: 0;
    flex-grow: 1;
    min-width: 0;
  }

  .access-row__name-text--issue {
    text-decoration: underline dashed;
  }

  .access-row__name-label {
    font-weight: 400;
    color: var(--color-text-2)
  }

  .access-row__actions {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 12px;
  }
}
</style>
