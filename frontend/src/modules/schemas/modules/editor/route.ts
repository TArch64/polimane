import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const schemaEditorRoute = defineViewRoute({
  name: 'schema-editor',
  path: 'editor',
  component: () => import('./Route.vue'),
});

export type SchemaEditorRoute = InferViewRouteInfo<typeof schemaEditorRoute>;
