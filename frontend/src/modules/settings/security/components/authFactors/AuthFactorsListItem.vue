<template>
  <CardListItem :title>
    <template #actions>
      <Button
        icon
        danger
        :style="deleteConfirm.anchorStyle"
        @click="deleteFactor"
      >
        <TrashIcon />
      </Button>
    </template>
  </CardListItem>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { CardListItem } from '@/components/card';
import { useAsyncAction, useDateFormatter } from '@/composables';
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import type { IAuthFactor } from '../../models';

const props = defineProps<{
  factor: IAuthFactor;
}>();

const createdAt = useDateFormatter(() => props.factor.createdAt);
const title = computed(() => `TOTP доданий ${createdAt.value}`);

const deleteConfirm = useConfirm({
  message: 'Ви впевнені, що хочете видалити цю двухфакторну автентифікацію?',
  acceptButton: 'Видалити',
  danger: true,
  control: false,
});

const deleteFactor = useAsyncAction(async () => {
  if (!await deleteConfirm.ask()) {
    return;
  }
});
</script>
