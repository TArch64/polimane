<template>
  <p class="bottom-bar-metric">
    <span>{{ label }}</span>

    <span class="bottom-bar-metric__value">
      <span class="bottom-bar-metric__mirror-value">
        {{ maxValue }}
      </span>

      <span class="bottom-bar-metric__actual-value">
        <slot>{{ value }}</slot>
      </span>
    </span>
  </p>
</template>

<script setup lang="ts">
defineProps<{
  label: string;
  value: string;
  maxValue: string;
}>();
</script>

<style scoped>
@layer page {
  .bottom-bar-metric {
    display: flex;
    justify-content: space-between;
    gap: 4px;
    padding: 0 6px;
  }

  .bottom-bar-metric__value {
    position: relative;
    color: var(--metric-value-color, inherit);
  }

  .bottom-bar-metric__mirror-value {
    visibility: hidden;
  }

  .bottom-bar-metric__actual-value {
    position: absolute;
    top: 0;
    right: 0;
    display: flex;
  }
}
</style>
