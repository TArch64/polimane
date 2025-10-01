<template>
  <div
    class="schema-export__preview"
    :style="styles"
    v-html="source"
  />
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import type { ISchema } from '@/models';
import { useHttpClient } from '@/composables';
import type { UrlPath } from '@/helpers';
import { collectUniqColors } from './collectUniqColors';

const props = defineProps<{
  schema: ISchema;
}>();

const http = useHttpClient();

const width = computed(() => props.schema.size.left + props.schema.size.right + 2);
const height = computed(() => props.schema.size.top + props.schema.size.bottom + 2);
const aspectRatio = computed(() => width.value / height.value);

const source = ref('');

const colors = computed(() => collectUniqColors(props.schema));

const styles = computed(() => {
  const raw = [
    ['background', props.schema.backgroundColor],
    ...colors.value.entries(),
  ];

  const vars = raw.map(([key, value]) => [`--ps-${key}`, value]);
  return Object.fromEntries(vars);
});

onMounted(async () => {
  const url: UrlPath = ['/schemas', props.schema.id, 'preview'];

  const svg = await http.get<string>(url, {}, {
    responseType: 'text',
  });

  const document = new DOMParser().parseFromString(svg, 'image/svg+xml');

  document.querySelector<SVGRectElement>('rect')!.setAttribute('fill', 'var(--ps-background)');

  for (const [index, color] of colors.value.entries()) {
    const beadEls = document.querySelectorAll<SVGCircleElement>(`[fill="${color}"]`);
    for (const el of beadEls) {
      el.setAttribute('fill', `var(--ps-${index})`);
    }
  }

  source.value = document.documentElement.outerHTML;
});

defineExpose({
  getSource: () => source.value,
});
</script>

<style scoped>
@layer page {
  .schema-export__preview {
    width: 100%;
    height: auto;
    border-radius: var(--rounded-md);
    border: var(--divider);
    overflow: clip;
    aspect-ratio: v-bind("aspectRatio");
    background-color: v-bind("schema.backgroundColor");

    :deep(svg) {
      width: 100%;
      height: 100%;
    }
  }
}
</style>
