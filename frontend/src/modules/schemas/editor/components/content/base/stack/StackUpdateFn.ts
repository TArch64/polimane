import Konva from 'konva';

export interface IStackUpdatePayload {
  parent: Konva.Group;
  child: Konva.Node;
  next: number;
}

export interface IStackUpdate<K extends keyof Konva.NodeConfig> {
  next: Konva.NodeConfig[K];
  property: K & string;
  extra?: Partial<Konva.NodeConfig>;
}

export type StackUpdateFn<K extends keyof Konva.NodeConfig> = (payload: IStackUpdatePayload) => IStackUpdate<K>;
