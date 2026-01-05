<template>
  <div class="subscription-limit">
    <svg class="subscription-limit__graph" viewBox="0 0 40 40">
      <ProgressCircle
        color="var(--color-background-3)"
        :value="100"
      />

      <ProgressCircle
        color="var(--color-primary)"
        :value="progress"
        v-if="progress"
      />
    </svg>

    <div class="subscription-limit__info">
      <h3 class="subscription-limit__title">
        {{ title }}
      </h3>

      <p class="subscription-limit__subtitle">
        {{ subtitle }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type FunctionalComponent, h } from 'vue';

defineProps<{
  title: string;
  subtitle: string;
  progress: number;
}>();

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
  .subscription-limit {
    position: relative;
    width: 140px;
    height: 120px;
  }

  .subscription-limit__graph {
    position: absolute;
    z-index: 1;
    width: 140px;
    height: 140px;
  }

  .subscription-limit__info {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    padding: 20px 20px 0;
    z-index: 2;
    height: 100%;
  }

  .subscription-limit__title {
    font-size: 16px;
    font-weight: 450;
    text-align: center;
    text-wrap: balance;
    margin-bottom: 4px;
  }

  .subscription-limit__subtitle {
    font-size: 14px;
    text-align: center;
  }
}
</style>
