import type { FabricObjectProps, TextProps } from 'fabric';

export const OBJECT_DEFAULTS: Partial<FabricObjectProps> = {
  selectable: false,
};

export const TEXT_OBJECT_DEFAULTS: Partial<TextProps> = {
  ...OBJECT_DEFAULTS,
  fontFamily: 'Arial',
};
