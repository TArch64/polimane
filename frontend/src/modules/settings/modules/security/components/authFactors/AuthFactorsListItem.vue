<template>
  <CardListItem :title>
    <template #actions>
      <Button
        icon
        danger
        title="Видалити"
        :style="deleteConfirm.anchorStyle"
        @click="deleteIntent"
      >
        <TrashIcon />
      </Button>
    </template>
  </CardListItem>
</template>

<script setup lang="ts">
import { computed, nextTick } from 'vue';
import { CardListItem } from '@/components/card';
import {
  useAsyncAction,
  useDateFormatter,
  useProgressBar,
  useRouteTransition,
} from '@/composables';
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import type { IAuthFactor } from '../../models';
import { useAuthFactorsStore } from '../../stores';

const props = defineProps<{
  factor: IAuthFactor;
}>();

const authFactorsStore = useAuthFactorsStore();
const routeTransition = useRouteTransition();

const createdAt = useDateFormatter(() => props.factor.createdAt);
const title = computed(() => `TOTP доданий ${createdAt.value}`);

const deleteConfirm = useConfirm({
  message: 'Ви впевнені, що хочете видалити цю двухфакторну автентифікацію?',
  acceptButton: 'Видалити',
  danger: true,
  control: false,
});

const deleteFactor = useAsyncAction(async () => {
  await authFactorsStore.deleteFactor(props.factor);
  routeTransition.start(() => nextTick());
});

useProgressBar(deleteFactor);

async function deleteIntent() {
  const confirmation = await deleteConfirm.ask();
  if (confirmation.isAccepted) await deleteFactor();
}
</script>
