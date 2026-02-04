<template>
  <div class="beads-limit-reached-popover">
    <p>Ви досягли максимальної кількості бісеру для цієї схеми</p>

    <Button
      variant="primary"
      size="md"
      class="beads-limit-reached-popover__upgrade"
      @click="upgradeModal.open()"
    >
      Хочу Більше
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { UpgradePlanModal } from '@/components/subscription';

const upgradeModal = useModal(UpgradePlanModal);
</script>

<style scoped>
@layer page {
  .beads-limit-reached-popover {
    width: 300px;
  }

  .beads-limit-reached-popover__upgrade {
    align-self: flex-end;
    margin-top: 8px;
    margin-left: auto;
  }
}
</style>
