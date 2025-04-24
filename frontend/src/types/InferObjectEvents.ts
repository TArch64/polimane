import type { FabricObject, ObjectEvents } from 'fabric';
import type { SafeAny } from '@/types/SafeAny';

export type InferObjectEvents<O> = O extends FabricObject<SafeAny, SafeAny, infer E extends ObjectEvents> ? E : never;
