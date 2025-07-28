<template>
  <Modal
    :width="ModalWidth.LG"
    title="Двохфакторна автентифікація"
    save-button="Увімкнути"
    @save="create"
  >
    <p class="new-auth-factor__start">
      Для активації двохфакторної автентифікації, будь ласка,
      відскануйте QR-код за допомогою вашого додатку для автентифікації
      (наприклад, Google Authenticator, 1Password, Apple Passwords тощо)
    </p>

    <div class="new-auth-factor__code-row">
      <div class="new-auth-factor__image-container">
        <img
          :src="init.qrCode"
          class="new-auth-factor__image"
          alt="QR код для двохфакторної автентифікації"
        />
      </div>

      <div>
        <p class="new-auth-factor__image-hint">
          Не виходить скористатися QR-кодом? Скопіюйте цей код та вставте його у ваш додаток для
          автентифікації
        </p>

        <CopyBadge :text="init.secret" />
      </div>
    </div>

    <p class="new-auth-factor__enter-code">
      Після сканування QR-коду, ваш додаток для автентифікації згенерує код,
      який потрібно ввести для завершення активації двохфакторної автентифікації
    </p>

    <TextField
      required
      variant="control"
      placeholder="Введіть код з вашого додатку"
      v-model="form.code"
    />
  </Modal>
</template>

<script setup lang="ts">
import { nextTick, reactive } from 'vue';
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { CopyBadge } from '@/components/badge';
import { TextField } from '@/components/form';
import { useAsyncAction } from '@/composables';
import type { IAuthFactorInit } from '../../models';
import { useAuthFactorsStore } from '../../stores';

const props = defineProps<{
  init: IAuthFactorInit;
}>();

const modal = useActiveModal();
const authFactorsStore = useAuthFactorsStore();

const form = reactive({
  code: '',
});

const create = useAsyncAction(async () => {
  const factor = await authFactorsStore.createFactor(props.init.challengeId, form.code);

  modal.close(null, () => {
    authFactorsStore.addFactor(factor);
    return nextTick();
  });
});
</script>

<style scoped>
@layer page {
  .new-auth-factor__start {
    margin-bottom: 16px;
  }

  .new-auth-factor__code-row {
    display: flex;
    align-items: center;
    padding-left: 8px;
    margin-bottom: 16px;
  }

  .new-auth-factor__image-container {
    display: flex;
    border: var(--divider);
    padding: 12px;
    border-radius: var(--rounded-md);
    margin-right: 24px;
  }

  .new-auth-factor__image {
    width: 120px;
    aspect-ratio: 1;
  }

  .new-auth-factor__image-hint {
    margin-top: -12px;
    margin-bottom: 6px;
  }

  .new-auth-factor__enter-code {
    margin-bottom: 12px;
  }
}
</style>
