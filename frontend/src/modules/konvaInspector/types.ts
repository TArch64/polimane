import type {
  CustomInspectorState,
  DevToolsV6PluginAPIHookPayloads,
  PluginSetupFunction,
} from '@vue/devtools-kit';

export type DevtoolsPluginApi = Parameters<PluginSetupFunction>[0];
export type CustomInspectorStateSection = CustomInspectorState[string][number];
export type EditInspectorStatePayload = DevToolsV6PluginAPIHookPayloads['editInspectorState'];
