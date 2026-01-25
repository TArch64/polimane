<template>
  <Card title="Видалити Акаунт">
    <p>
      Видалення акаунту є необоротною дією. Вся ваша особиста інформація буде видалена з нашої
      системи через 30 днів.
      Будь ласка, переконайтеся, що ви розумієте наслідки цієї дії перед тим, як продовжити
    </p>

    <Button
      danger
      variant="secondary"
      class="delete-account__button"
      :style="deleteConfirm.anchorStyle"
      :loading="deleteAccount.isActive"
      @click="deleteIntent"
    >
      Видалити Акаунт
    </Button>
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { Button } from '@/components/button';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction } from '@/composables';
import { useProfileStore } from '../stores';

const profileStore = useProfileStore();

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити свій акаунт?',
  acceptButton: 'Видалити',
});

const deleteAccount = useAsyncAction(async () => {
  await profileStore.deleteAccount();
});

async function deleteIntent(): Promise<void> {
  const confirmation = await deleteConfirm.ask();
  if (confirmation.isAccepted) await deleteAccount();
}
</script>

<style scoped>
@layer page {
  .delete-account__button {
    margin-top: 8px;
    margin-left: auto;
  }
}
</style>
