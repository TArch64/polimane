import { Rect } from 'fabric';

export const enum EditorObjectType {
  PATTERN = 'pattern',
}

export const EditorObjectTypeList = [
  EditorObjectType.PATTERN,
] as const;

export type EditorObjectTypeMap = {
  [EditorObjectType.PATTERN]: Rect;
};
