import { FabricImage, type ImageProps } from 'fabric';
import type { SafeAny } from '@/types';

export interface IIconObjectProps extends Partial<ImageProps> {
  source: SafeAny;
  color?: string;
}

export class IconObject extends FabricImage {
  private static buildSvgUrl(input: string, color?: string): string {
    let svg = color ? input.replace('currentColor', color) : input;
    svg = svg.replace('\n', '');

    if (!svg.includes('xmlns="http://www.w3.org/2000/svg"')) {
      svg = svg.replace('<svg', '<svg xmlns="http://www.w3.org/2000/svg"');
    }

    return `data:image/svg+xml;base64,${btoa(svg)}`;
  }

  constructor({ source, color, ...props }: IIconObjectProps) {
    const imageEl = document.createElement('img');
    imageEl.src = IconObject.buildSvgUrl(source, color);

    super(imageEl, {
      ...props,
      width: imageEl.width,
      height: imageEl.height,
    });

    if (props.width) this.scaleToWidth(props.width);
    if (props.height) this.scaleToHeight(props.height);
  }
}
