import { FabricText, Group, Rect } from 'fabric';
import type { ISchemaPattern } from '@/models';
import { TEXT_OBJECT_DEFAULTS } from '../objectDefaults';

export class PatternTitleObject extends Group {
  private static PADDING_VERTICAL = 5;
  private static PADDING_HORIZONTAL = 8;

  private static MIN_TEXT_WIDTH = 50;

  private readonly border: Rect;
  private readonly text: FabricText;

  constructor(pattern: ISchemaPattern) {
    super();

    this.text = this.createText(pattern);
    this.border = this.createBorder();

    this.add(this.border, this.text);
  }

  private createText(pattern: ISchemaPattern): FabricText {
    const text = new FabricText(pattern.name, {
      ...TEXT_OBJECT_DEFAULTS,
      top: PatternTitleObject.PADDING_VERTICAL,
      left: PatternTitleObject.PADDING_HORIZONTAL,
      fontSize: 14,
      textAlign: 'center',
      lineHeight: 1.7,
    });

    if (text.width < PatternTitleObject.MIN_TEXT_WIDTH) {
      text.set({ width: PatternTitleObject.MIN_TEXT_WIDTH });
    }

    return text;
  }

  private createBorder(): Rect {
    return new Rect({
      ...this.borderSize,
      rx: 4,
      ry: 4,
      stroke: 'rgba(0, 0, 0, 0.2)',
      fill: 'white',
    });
  }

  update(pattern: ISchemaPattern) {
    this.text.set({ text: pattern.name });
    this.border.set(this.borderSize);
  }

  private get borderSize() {
    return {
      width: this.text.width + PatternTitleObject.PADDING_HORIZONTAL * 2,
      height: this.text.height + PatternTitleObject.PADDING_VERTICAL * 2,
    };
  }
}
