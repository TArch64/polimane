import Konva from 'konva';
import type { CustomInspectorState } from '@vue/devtools-kit';
import type { InspectorHighlight } from './hightlight';
import type { CustomInspectorStateSection, EditInspectorStatePayload } from './types';

export class InspectorState {
  private SELECTOR_ATTRS = ['id', 'name'];
  private DIMENSIONS_ATTRS = ['x', 'y', 'width', 'height'];

  constructor(private readonly highlight: InspectorHighlight) {
  }

  getInspectorState(stage: Konva.Stage, nodeId: number): CustomInspectorState {
    const node = this.findNodeById(stage, nodeId);

    if (node) {
      this.highlight.show(node);
      return this.formatNodeState(node);
    } else {
      this.highlight.hide();
      return {};
    }
  }

  private findNodeById(stage: Konva.Stage, id: number): Konva.Node | null {
    if (stage._id === id) return stage;

    return stage.findOne((node: Konva.Node) => node._id === id) ?? null;
  }

  private formatNodeState(node: Konva.Node): CustomInspectorState {
    return {
      Selectors: this.SELECTOR_ATTRS.map((key) => this.formatProperty(node, key)),
      Dimensions: this.DIMENSIONS_ATTRS.map((key) => this.formatProperty(node, key)),
      Attrs: this.formatAttrs(node),
    };
  }

  private formatProperty(node: Konva.Node, key: string): CustomInspectorStateSection {
    const value = node.getAttr(key);
    const isEditable = typeof value === 'string' || typeof value === 'number' || Array.isArray(value);

    return {
      key,
      objectType: 'other',
      editable: isEditable,
      raw: value,
      value: value,
    };
  }

  private formatAttrs(node: Konva.Node): CustomInspectorStateSection[] {
    const excluded = this.SELECTOR_ATTRS.concat(this.DIMENSIONS_ATTRS);
    const attrs = Object.keys(node.getAttrs()).filter((key) => !excluded.includes(key));
    return attrs.map((key) => this.formatProperty(node, key));
  }

  editInspectorState(stage: Konva.Stage, payload: EditInspectorStatePayload) {
    const node = this.findNodeById(stage, +payload.nodeId);
    if (!node) return;

    const key = payload.path[0];
    const value = payload.state.value;

    // array / object edit
    if (payload.path.length > 1) {
      const object = node.getAttr(key);
      const subKey = payload.path[1];

      if (Array.isArray(object)) {
        const newValue = [...object];
        newValue[+subKey] = payload.state.value;
        return node.setAttr(key, newValue);
      }

      return node.setAttr(key, {
        ...object,
        [subKey]: value,
      });
    }

    node.setAttr(key, value);
  }
}
