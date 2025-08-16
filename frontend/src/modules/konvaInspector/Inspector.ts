import Konva from 'konva';
import type { DevtoolsPluginApi } from './types';
import { InspectorHighlight } from './highlight';
import { InspectorTree } from './InspectorTree';
import { InspectorState } from './InspectorState';

const ID = 'konva-nodes';

export class Inspector {
  private stage: Konva.Stage | null = null;
  private readonly highlight = new InspectorHighlight();
  private readonly tree = new InspectorTree(this.highlight.layerId);
  private readonly state = new InspectorState(this.highlight, () => this.refreshInspectorState());

  constructor(private readonly api: DevtoolsPluginApi) {
    api.addInspector({
      id: ID,
      label: 'Konva Nodes',
    });

    api.on.getInspectorTree((payload) => {
      if (payload.inspectorId === ID) {
        payload.rootNodes = this.stage ? this.tree.collectInspectorTree(this.stage) : [];
      }
    });

    api.on.getInspectorState((payload) => {
      if (payload.inspectorId === ID) {
        payload.state = this.stage ? this.state.getInspectorState(this.stage, +payload.nodeId) : {};
      }
    });

    api.on.editInspectorState((payload) => {
      if (payload.inspectorId === ID && this.stage) {
        this.state.editInspectorState(this.stage, payload);
      }
    });

    this.refreshInspectorTree = this.refreshInspectorTree.bind(this);
  }

  useStage(stage: Konva.Stage | null): void {
    this.stage = stage;

    if (stage) {
      this.highlight.useStage(stage);
    }

    this.refreshInspectorState();
    this.refreshInspectorTree();
  }

  private refreshInspectorTree(): void {
    this.api.sendInspectorTree(ID);
  }

  private refreshInspectorState(): void {
    this.api.sendInspectorState(ID);
  }
}
