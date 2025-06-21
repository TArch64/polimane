import Konva from 'konva';
import { NodeRect } from '@/models';

export interface IStackUpdatePayload {
  parentRect: NodeRect;
  childRect: NodeRect;
  next: number;
}

export interface IStackUpdate<K extends keyof Konva.NodeConfig> {
  next: Konva.NodeConfig[K];
  property: K & string;
  extra?: Partial<Konva.NodeConfig>;
}

export type StackUpdateFn<K extends keyof Konva.NodeConfig> = (payload: IStackUpdatePayload) => IStackUpdate<K>;
