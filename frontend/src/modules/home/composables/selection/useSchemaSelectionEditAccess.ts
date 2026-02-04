import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { LockIcon, PeopleIcon } from '@/components/icon';
import { SchemaAccessEditModal } from '@/modules/schemas/shared/modals/accessEdit';
import { useModal } from '@/components/modal';
import { SubscriptionLimit } from '@/enums';
import { useSessionStore } from '@/stores';
import { UpgradePlanModal } from '@/components/subscription';

export function useSchemaSelectionEditAccess(actionIds: Ref<string[]>): Ref<MaybeContextMenuAction> {
  const sessionStore = useSessionStore();

  const sharedAccessLimit = computed(() => sessionStore.getLimit(SubscriptionLimit.SHARED_ACCESS));
  const isAvailable = computed(() => sharedAccessLimit.value! > 1);

  const accessEditModal = useModal(SchemaAccessEditModal);
  const upgradePlanModal = useModal(UpgradePlanModal);

  return computed((): MaybeContextMenuAction => !!actionIds.value.length && {
    title: 'Редагувати Доступ',
    icon: isAvailable.value ? PeopleIcon : LockIcon,

    async onAction() {
      if (!isAvailable.value) {
        const isUpgraded = await upgradePlanModal.open();
        if (!isUpgraded) return;
      }

      void accessEditModal.open({
        schemaIds: actionIds.value,
      });
    },
  });
}
