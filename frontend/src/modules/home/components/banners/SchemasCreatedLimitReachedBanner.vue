<template>
  <Banner type="danger" :prepend-icon="LockIcon" v-if="counter.isReached">
    <p>Ліміт схем перевищено: <b>{{ counter.current }}</b> / <b>{{ counter.max }}</b></p>
    <p v-if="canUpgrade">Оновіть підписку або видаліть схеми</p>

    <template #actions v-if="canUpgrade">
      <Button variant="primary" @click="upgradePlanModal.open()">
        Оновити Підписку
      </Button>
    </template>
  </Banner>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Banner } from '@/components/banner';
import { LockIcon } from '@/components/icon';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { useSchemasCreatedCounter } from '@/composables/subscription';
import { UpgradePlanModal } from '@/components/subscription';
import { useSessionStore } from '@/stores';
import { isMaxSubscriptionPlan } from '@/enums';

const sessionStore = useSessionStore();

const counter = useSchemasCreatedCounter();
const upgradePlanModal = useModal(UpgradePlanModal);

const canUpgrade = computed(() => !isMaxSubscriptionPlan(sessionStore.plan.id));
</script>
