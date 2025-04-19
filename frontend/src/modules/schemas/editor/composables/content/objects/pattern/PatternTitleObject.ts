import { FabricText, Group, Rect } from 'fabric';
import type { ISchemaPattern } from '@/models';

export class PatternTitleObject extends Group {
  static PADDING_VERTICAL = 4;
  static PADDING_HORIZONTAL = 6;

  private readonly border: Rect;
  private readonly text: FabricText;

  constructor(pattern: ISchemaPattern) {
    super();

    this.text = new FabricText(pattern.name, {
      top: PatternTitleObject.PADDING_VERTICAL - 1,
      left: PatternTitleObject.PADDING_HORIZONTAL,
      fontSize: 14,
      fontFamily: 'Arial',
    });

    this.border = new Rect({
      rx: 4,
      ry: 4,
      height: this.text.height + PatternTitleObject.PADDING_VERTICAL * 2,
      width: this.text.width + PatternTitleObject.PADDING_HORIZONTAL * 2,
      stroke: 'rgba(0, 0, 0, 0.2)',
      fill: 'white',
    });

    this.add(this.border, this.text);
  }
}
