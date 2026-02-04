<template>
  <Component
    :is="tag"
    class="bottom-bar-metric"
    :class="classes"
    v-bind="attrs"
    v-on="listeners"
  >
    <span>{{ label }}</span>

    <span class="bottom-bar-metric__value">
      <span class="bottom-bar-metric__mirror-value">
        {{ maxValue }}
      </span>

      <span class="bottom-bar-metric__actual-value">
        <slot>{{ value }}</slot>
      </span>
    </span>
  </Component>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = withDefaults(defineProps<{
  label: string;
  value: string;
  maxValue: string;
  interactive?: boolean;
}>(), {
  interactive: false,
});

const emit = defineEmits<{
  click: [];
}>();

const tag = computed(() => props.interactive ? 'button' : 'p');

const classes = computed(() => ({
  'bottom-bar-metric--interactive': props.interactive,
}));

const attrs = computed(() => props.interactive
  ? { type: 'button' }
  : {});

const listeners = computed(() => props.interactive
  ? { click: () => emit('click') }
  : {},
);
</script>

<style scoped>
@layer page {
  .bottom-bar-metric {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 100%;
    gap: 4px;
    padding: 3px 6px;
    line-height: 1;
  }

  .bottom-bar-metric--interactive {
    cursor: pointer;
    color: inherit;
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
