import { FabricObject } from 'fabric';
import { combineStopHandles } from '@/helpers';
import type { InferObjectProps } from '@/types';

type ObjectHoverState = 'default' | 'hover';
type ObjectHoverStates<O extends FabricObject> = Record<ObjectHoverState, Partial<InferObjectProps<O>>>;

interface IHoverTarget {
  object: FabricObject;
  states: ObjectHoverStates<FabricObject>;
}

export class ObjectHover {
  readonly dispose: VoidFunction;
  private readonly targets: IHoverTarget[] = [];

  constructor(private readonly object: FabricObject) {
    this.dispose = combineStopHandles(
      object.on('mouseover', () => this.renderState('hover')),
      object.on('mouseout', () => this.renderState('default')),
    );
  }

  private renderState(state: ObjectHoverState) {
    for (const { object, states } of this.targets) {
      object.set(states[state]);
    }

    this.requestRenderAll();
  }

  private requestRenderAll() {
    this.object.canvas?.requestRenderAll();
  }

  apply<O extends FabricObject>(object: O, states: ObjectHoverStates<O>): this {
    object.set(states.default);
    this.targets.push({ object, states });
    return this;
  }
}
