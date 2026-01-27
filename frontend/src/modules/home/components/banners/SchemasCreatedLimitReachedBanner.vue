<template>
  <Banner type="danger" :prepend-icon="LockIcon" v-if="counter.isReached">
    <p>Ліміт схем перевищено: <b>{{ counter.current }}</b> / <b>{{ counter.max }}</b></p>
    <p>Оновіть підписку або видаліть схеми</p>

    <template #actions>
      <Button variant="primary" @click="upgradePlanModal.open()">
        Оновити Підписку
      </Button>
    </template>
  </Banner>
</template>

<script setup lang="ts">
import { useSchemasCreatedCounter } from '@/composables/subscription';
import { Banner } from '@/components/banner';
import { LockIcon } from '@/components/icon';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { UpgradePlanModal } from '@/components/subscription';

const upgradePlanModal = useModal(UpgradePlanModal);
const counter = useSchemasCreatedCounter();
</script>
