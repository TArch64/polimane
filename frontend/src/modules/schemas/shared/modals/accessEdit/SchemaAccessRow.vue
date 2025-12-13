<template>
  <li class="access-row">
    <p class="access-row__name">
      {{ displayName }}

      <span v-if="!editable" class="access-row__name-label">
        (Ви)
      </span>

      <span class="access-row__name-label" v-else-if="unevenAccess">
        *
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
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { AccessLevel } from '@/enums';
import SchemaAccessField from './SchemaAccessField.vue';

withDefaults(defineProps<{
  displayName: string;
  unevenAccess: boolean;
  editable?: boolean;
}>(), {
  editable: true,
});

const access = defineModel<AccessLevel>('access', { required: true });

const emit = defineEmits<{
  delete: [];
}>();

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете заборонити доступ до схему цьому користувачу?',
  acceptButton: 'Заборонити',
});

async function deleteUserIntent(): Promise<void> {
  const confirmed = await deleteConfirm.ask();
  if (confirmed.isAccepted) emit('delete');
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
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .access-row__name-label {
    font-weight: 400;
    color: var(--color-text-3)
  }

  .access-row__actions {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 12px;
  }
}
</style>
