import { type FunctionPlugin, watch } from 'vue';
import type { PluginDescriptor } from '@vue/devtools-kit';
import { setupDevtoolsPlugin } from '@vue/devtools-api';
import { Inspector } from './Inspector';

export const konvaInspectorPlugin: FunctionPlugin = (app) => {
  const DESCRIPTOR: PluginDescriptor = {
    app,
    id: 'konva-inspector',
    packageName: '@tarch64/konva-inspector',
    label: 'Konva',
  };

  setupDevtoolsPlugin(DESCRIPTOR, (api) => {
    const inspector = new Inspector(api);

    watch(window.__KONVA_STAGE_REF__, (stage) => {
      inspector.useStage(stage);
    }, { immediate: true });
  });
};
