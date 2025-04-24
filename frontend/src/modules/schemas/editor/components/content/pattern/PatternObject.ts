import { Group, Rect } from 'fabric';
import type { ISchemaPattern } from '@/models';
import type { IObjectOnUpdate } from '@/modules/schemas/editor/composables';
import { ObjectHover } from '../ObjectHover';
import { OBJECT_DEFAULTS } from '../objectDefaults';
import { PatternTitleObject } from './PatternTitleObject';

export class PatternObject extends Group implements IObjectOnUpdate<ISchemaPattern> {
  private readonly border: Rect;
  private readonly hover: ObjectHover;
  private readonly title: PatternTitleObject;

  constructor(pattern: ISchemaPattern) {
    super([], OBJECT_DEFAULTS);

    this.hover = new ObjectHover(this);
    this.border = this.createBorder();
    this.title = this.createTitle(pattern);

    this.add(this.border, this.title);
  }

  private createBorder(): Rect {
    const border = new Rect({
      rx: 8,
      ry: 8,
      width: 1000,
      height: 100,
      strokeDashArray: [10, 5],
      fill: 'transparent',
    });

    this.hover.apply(border, {
      default: { stroke: 'rgba(0, 0, 0, 0.2)' },
      hover: { stroke: '#000' },
    });

    return border;
  }

  private createTitle(pattern: ISchemaPattern): PatternTitleObject {
    const title = new PatternTitleObject(pattern);
    title.setX(8);
    title.setY(-title.height / 2);
    return title;
  }

  onUpdate(pattern: ISchemaPattern) {
    this.title.update(pattern);
  }

  dispose() {
    super.dispose();
    this.hover.dispose();
  }
}
