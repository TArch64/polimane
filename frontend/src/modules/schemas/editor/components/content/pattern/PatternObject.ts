import { Group, Rect } from 'fabric';
import type { ISchemaPattern } from '@/models';
import type { IObjectOnUpdate } from '@/modules/schemas/editor/composables';
import { ObjectHover } from '../ObjectHover';
import { OBJECT_DEFAULTS } from '../objectDefaults';
import { ObjectAnimationFrame } from '../ObjectAnimationFrame';
import { PatternTitleObject } from './PatternTitleObject';

export class PatternObject extends Group implements IObjectOnUpdate<ISchemaPattern> {
  private static CONTAINER_MIN_WIDTH = 1000;
  private static CONTAINER_MIN_HEIGHT = 100;

  private readonly border;
  private readonly hover;
  private readonly title;
  private readonly offListeners;
  private readonly containerPositionAnimationFrame = new ObjectAnimationFrame();

  readonly container;

  constructor(pattern: ISchemaPattern) {
    super([], {
      ...OBJECT_DEFAULTS,
    });

    this.hover = new ObjectHover(this);
    this.border = this.createBorder();
    this.title = this.createTitle(pattern);

    this.container = new Group([], { left: 0, top: 0 });
    this.offListeners = this.container.on('layout:after', this.onContainerResize.bind(this));

    this.add(this.border, this.title, this.container);
  }

  private createBorder(): Rect {
    const border = new Rect({
      rx: 8,
      ry: 8,
      strokeDashArray: [10, 5],
      fill: 'transparent',
    });

    this.hover.apply(border, {
      default: { stroke: 'rgba(0, 0, 0, 0.1)' },
      hover: { stroke: 'rgba(0, 0, 0, 0.5)' },
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

  private onContainerResize() {
    this.border.set(this.borderSize);

    this.containerPositionAnimationFrame.request(() => {
      const left = Math.max((this.border.width - this.container.width) / 2, 0);

      if (this.container.left !== left) {
        this.container.set({ left: this.border.left + left });
      }

      const top = Math.max((this.border.height - this.container.height) / 2, 0);

      if (this.container.top !== top) {
        this.container.set({ top: this.border.top + top });
      }

      this.canvas?.requestRenderAll();
    });
  }

  private get borderSize() {
    return {
      width: Math.max(this.container.width, PatternObject.CONTAINER_MIN_WIDTH),
      height: Math.max(this.container.height, PatternObject.CONTAINER_MIN_HEIGHT),
    };
  }

  dispose() {
    super.dispose();
    this.hover.dispose();
    this.containerPositionAnimationFrame.dispose();
    this.offListeners();
  }
}
