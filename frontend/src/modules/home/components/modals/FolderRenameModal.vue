<template>
  <Modal
    title="Змінити Назву Директорії"
    :loading="save.isActive"
    @save="save"
  >
    <TextField
      required
      ref="nameRef"
      variant="control"
      placeholder="Назва"
      v-model="form.name"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import type { IFolder } from '@/models';
import { useHomeStore } from '@/modules/home/stores';
import { HttpError, HttpErrorReason, useAsyncAction } from '@/composables';

const props = defineProps<{
  folder: IFolder;
}>();

const homeStore = useHomeStore();

const nameRef = ref<ComponentExposed<typeof TextField>>(null!);
const modal = useActiveModal();

const form = reactive({
  name: props.folder.name,
});

const save = useAsyncAction(async () => {
  form.name = form.name.trim();

  if (form.name === props.folder.name) {
    return modal.close(null);
  }

  try {
    await homeStore.updateFolder.do(props.folder, {
      name: form.name,
    });

    modal.close(null);
  } catch (error) {
    if (HttpError.isReason(error, HttpErrorReason.ALREADY_IN_USE_NAME)) {
      nameRef.value.setError('Ця назва вже використовується');
    }
  }
});
</script>
