import { computed } from 'vue';
import { type ColorObject, contrastWCAG21, parse, serialize, to } from 'colorjs.io/fn';
import { useEditorStore } from '@editor/stores';
import { MIN_CONTRAST_AA } from './useContrast';

const srgb = (value: number): ColorObject => ({
  space: 'srgb',
  coords: [value, value, value],
});

function sum(nums: number[]): number {
  return nums.reduce((a, b) => a + b, 0);
}

export function useBackgroundAccessibleColor() {
  const editorStore = useEditorStore();

  const bg = computed(() => parse(editorStore.schema.backgroundColor));
  const bgSrgb = computed(() => to(bg.value, 'srgb'));
  const isDarkBg = computed(() => sum(bgSrgb.value.coords) / 3 < 0.5);

  return computed(() => {
    let low = isDarkBg.value ? 0.5 : 0;
    let high = isDarkBg.value ? 1 : 0.5;
    let result = isDarkBg.value ? 1 : 0;

    while (high - low > 0.001) {
      const mid = (low + high) / 2;
      const contrastValue = contrastWCAG21(bg.value, srgb(mid));

      if (contrastValue >= MIN_CONTRAST_AA) {
        result = mid;
        if (isDarkBg.value) {
          high = mid;
        } else {
          low = mid;
        }
      } else {
        if (isDarkBg.value) {
          low = mid;
        } else {
          high = mid;
        }
      }
    }

    return serialize(srgb(result), {
      format: 'hex',
    });
  });
}
