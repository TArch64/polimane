import Konva from 'konva';

export interface IStackUpdatePayload {
  parent: Konva.Group;
  child: Konva.Node;
  next: number;
  isInitial: boolean;
}

export interface IStackUpdate {
  next: number;
  tween?: Konva.Tween;
}

export type StackUpdateFn = (payload: IStackUpdatePayload) => IStackUpdate;
