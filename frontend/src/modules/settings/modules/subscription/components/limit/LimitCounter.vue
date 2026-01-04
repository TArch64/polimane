<template>
  <div class="limit-counter">
    <svg class="limit-counter__graph" viewBox="0 0 40 40">
      <ProgressCircle
        color="var(--color-background-3)"
        :value="100"
      />

      <ProgressCircle
        color="var(--color-primary)"
        :value="usedPercentage"
        v-if="usedPercentage"
      />
    </svg>

    <div class="limit-counter__info">
      <h3 class="limit-counter__name">
        {{ limit.title }}
      </h3>

      <p class="limit-counter__value">
        {{ limit.used }}/{{ limit.max }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, type FunctionalComponent, h } from 'vue';
import type { ISubscriptionLimit } from '../../stores';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

const usedPercentage = computed(() => {
  if (props.limit.max === null) {
    return 0;
  }

  const used = Math.min(props.limit.used!, props.limit.max);
  return (used / props.limit.max) * 100;
});

interface IProgressCircleProps {
  value: number;
  color: string;
}

const ProgressCircle: FunctionalComponent<IProgressCircleProps> = (props) => h('circle', {
  'cx': '20',
  'cy': '20',
  'r': '18.5',
  'fill': 'none',
  'stroke': props.color,
  'stroke-width': '3',
  'stroke-linecap': 'round',
  'pathLength': '100',
  'stroke-dasharray': `${props.value * 0.75} ${100 - props.value * 0.75}`,
  'transform': 'rotate(135 20 20)',
});
</script>

<style scoped>
@layer page {
  .limit-counter {
    position: relative;
    width: 140px;
    height: 120px;
  }

  .limit-counter__graph {
    position: absolute;
    z-index: 1;
    width: 140px;
    height: 140px;
  }

  .limit-counter__info {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    padding: 20px 20px 0;
    z-index: 2;
    height: 100%;
  }

  .limit-counter__name {
    font-size: 16px;
    font-weight: 450;
    text-align: center;
    text-wrap: balance;
    margin-bottom: 4px;
  }

  .limit-counter__value {
    font-size: 14px;
    text-align: center;
  }
}
</style>
