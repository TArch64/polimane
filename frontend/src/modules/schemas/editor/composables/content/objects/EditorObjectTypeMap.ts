import { EditorObjectType } from '../../../enums';
import type { PatternObject } from './pattern';

export type EditorObjectTypeMap = {
  [EditorObjectType.PATTERN]: PatternObject;
};
