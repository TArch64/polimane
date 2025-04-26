import { FabricText, Group, Rect } from 'fabric';
import type { IObjectOnAdded, ObjectParent } from '@/modules/schemas/editor/composables';
import { OBJECT_DEFAULTS, TEXT_OBJECT_DEFAULTS } from '../objectDefaults';
import { IconObject } from '../IconObject';
import plusIcon from '~icons/eva/plus-outline?raw';

export class PatternEmptyObject extends Group implements IObjectOnAdded {
  private static PADDING_HORIZONTAL = 16;
  private static PADDING_VERTICAL = 6;

  private readonly text: FabricText;
  private readonly icon: IconObject;
  private readonly background: Rect;

  constructor() {
    super();

    this.icon = this.createIcon();
    this.text = this.createText();

    this.add(this.icon, this.text);

    this.background = this.createBackground();
    this.add(this.background);
    this.sendObjectToBack(this.background);
  }

  private createIcon(): IconObject {
    return new IconObject({
      ...OBJECT_DEFAULTS,
      source: plusIcon,
      color: 'rgba(0, 0, 0, 0.8)',
      left: PatternEmptyObject.PADDING_HORIZONTAL,
      top: PatternEmptyObject.PADDING_VERTICAL,
      width: 24,
      height: 24,
    });
  }

  private createText(): FabricText {
    return new FabricText('Додати Рядок', {
      ...TEXT_OBJECT_DEFAULTS,
      fontSize: 16,
      left: 32 + PatternEmptyObject.PADDING_HORIZONTAL,
      top: 3 + PatternEmptyObject.PADDING_VERTICAL,
      fill: 'rgba(0, 0, 0, 0.8)',
    });
  }

  private createBackground(): Rect {
    return new Rect({
      top: 0,
      left: 0,
      rx: 8,
      ry: 8,
      width: this.width + 4 + PatternEmptyObject.PADDING_HORIZONTAL * 2,
      height: this.height + PatternEmptyObject.PADDING_VERTICAL * 2,
      fill: 'transparent', // 'rgba(255, 255, 255, 0.8)',
      stroke: 'rgba(0, 0, 0, 0.2)',
    });
  }

  onAdded(parent: ObjectParent) {
    this.setX((parent.width - this.width) / 2);
    this.setY(((parent.height - this.height) / 2) - 5);
  }
}
