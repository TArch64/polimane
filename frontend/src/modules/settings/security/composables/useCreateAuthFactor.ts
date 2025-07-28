import { useModal } from '@/components/modal';
import { useAsyncAction } from '@/composables';
import AuthFactorAddModal from '../components/authFactors/AuthFactorAddModal.vue';
import { useAuthFactorsStore } from '../stores';

export function useCreateAuthFactor() {
  const addModal = useModal(AuthFactorAddModal);
  const authFactorsStore = useAuthFactorsStore();

  return useAsyncAction(async () => {
    const init = await authFactorsStore.initNewFactor();
    await addModal.open({ init });
  });
}
