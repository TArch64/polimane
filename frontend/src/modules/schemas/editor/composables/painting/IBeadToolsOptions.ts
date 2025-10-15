import type { Ref } from 'vue';
import type { IBeadsGrid } from '@editor/composables';

export interface IBeadToolsOptions {
  backgroundRef: Ref<SVGRectElement>;
  beadsGrid: IBeadsGrid;
}
