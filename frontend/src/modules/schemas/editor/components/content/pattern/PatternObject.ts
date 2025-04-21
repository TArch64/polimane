import { Group, Rect } from 'fabric';
import type { ISchemaPattern } from '@/models';
import type { IUpdatableFabricObject } from '@/modules/schemas/editor/composables';
import { PatternTitleObject } from './PatternTitleObject';

export class PatternObject extends Group implements IUpdatableFabricObject<ISchemaPattern> {
  private readonly border: Rect;
  private readonly title: PatternTitleObject;

  constructor(pattern: ISchemaPattern) {
    super([], {
      selectable: false,
    });

    this.border = new Rect({
      rx: 8,
      ry: 8,
      width: 1000,
      height: 100,
      stroke: 'rgba(0, 0, 0, 0.2)',
      strokeDashArray: [10, 5],
      fill: 'transparent',
    });

    this.title = new PatternTitleObject(pattern);
    this.title.setX(8);
    this.title.setY(-this.title.height / 2);

    this.add(this.border, this.title);
  }

  update(pattern: ISchemaPattern) {
    this.title.update(pattern);
  }
}
