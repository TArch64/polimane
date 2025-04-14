export const enum PatternType {
  SQUARE = 'square',
  DIAMOND = 'diamond',
}

export interface ISchemaPattern {
  id: string;
  type: PatternType;
}
