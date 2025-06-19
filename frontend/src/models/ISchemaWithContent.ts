import type { ISchemaObject } from './ISchemaObject';

export interface ISchemaWithContent<C extends ISchemaObject = ISchemaObject> extends ISchemaObject {
  content: C[];
}

export type InferSchemaContent<O> = O extends ISchemaWithContent<infer C extends ISchemaObject> ? C : never;

export function isSchemaWithContent(obj: ISchemaObject): obj is ISchemaWithContent {
  return 'content' in obj;
}
