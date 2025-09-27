<template>
  <Teleport to="body" v-if="previewSchema">
    <div inert hidden>
      <PreviewSchema ref="previewRef" :schema="previewSchema" />
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { nextTick, ref, toRaw } from 'vue';
import type { ISchema } from '@/models';
import { useDomRef } from '@/composables';
import { useEditorStore } from '../stores';
import { PreviewSchema } from './preview';

const editorStore = useEditorStore();

const previewSchema = ref<ISchema | null>(null);
const previewRef = useDomRef<SVGSVGElement>();
const serializer = new XMLSerializer();

async function generateScreenshot(): Promise<string | null> {
  previewSchema.value = structuredClone(toRaw(editorStore.schema));
  await nextTick();

  const content = serializer.serializeToString(previewRef.value);
  previewSchema.value = null;
  return content;
}

function needScreenshot(): boolean {
  if (!editorStore.schema.screenshotedAt) {
    return true;
  }

  const lastSaved = new Date(editorStore.schema.screenshotedAt);
  const now = new Date();
  const diff = now.getTime() - lastSaved.getTime();

  return diff > 30 * 60 * 1000;
}

editorStore.onSaved(async () => {
  const source = needScreenshot() ? await generateScreenshot() : null;
  if (source) await editorStore.updateScreenshot(source);
});
</script>
