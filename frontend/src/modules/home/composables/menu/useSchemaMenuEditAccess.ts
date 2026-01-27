import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useAccessPermissions } from '@/composables';
import type { ListSchema } from '@/modules/home/stores';
import { LockIcon, PeopleIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { SchemaAccessEditModal } from '@/modules/schemas/shared/modals/accessEdit';
import { useSessionStore } from '@/stores';
import { SubscriptionLimit } from '@/enums';
import { UpgradePlanModal } from '@/components/subscription';

export function useSchemaMenuEditAccess(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const sessionStore = useSessionStore();

  const permissions = useAccessPermissions(() => schema.value.access);
  const sharedAccessLimit = computed(() => sessionStore.getLimit(SubscriptionLimit.SHARED_ACCESS));
  const isAvailable = computed(() => sharedAccessLimit.value! > 1);

  const accessEditModal = useModal(SchemaAccessEditModal);
  const upgradePlanModal = useModal(UpgradePlanModal);

  return computed((): MaybeContextMenuAction => permissions.admin && {
    title: 'Редагувати Доступ',
    icon: isAvailable.value ? PeopleIcon : LockIcon,

    async onAction() {
      if (!isAvailable.value) {
        const isUpgraded = await upgradePlanModal.open();
        if (!isUpgraded) return;
      }

      void accessEditModal.open({
        schemaIds: [schema.value.id],
      });
    },
  });
}
