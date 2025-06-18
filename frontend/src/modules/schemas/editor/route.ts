import { defineRoute, type InferViewRouteInfo } from '@/router/define';

export const schemaEditorRoute = defineRoute({
  name: 'schema-editor',
  path: 'editor',
  component: () => import('./Route.vue'),
});

export type SchemaEditorRoute = InferViewRouteInfo<typeof schemaEditorRoute>;
