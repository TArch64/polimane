import { FabricObject, type FabricObjectProps } from 'fabric';

export type InferObjectProps<O> = O extends FabricObject<infer P extends FabricObjectProps> ? P : never;
