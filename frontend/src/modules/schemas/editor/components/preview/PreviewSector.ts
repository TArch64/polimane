import { type FunctionalComponent, h } from 'vue';
import type { ISchema } from '@/models';
import type { IBeadsGridItem } from '../../composables';
import { PreviewBead } from './PreviewBead';

export interface IPreviewSectorProps {
  schema: ISchema;
  grid: IBeadsGridItem[];
}

export const PreviewSector: FunctionalComponent<IPreviewSectorProps> = (props) => (
  props.grid.map((bead) => h(PreviewBead, {
    key: bead.coord,
    offset: bead.offset,
    color: props.schema.beads[bead.coord]!,
  }))
);
