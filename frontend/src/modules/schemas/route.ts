import { schemaEditorRoute, type SchemaEditorRoute } from '@editor/route';
import { defineWrapperRoute, type InferWrapperRouteInfo } from '@/router/define';

export const schemasRoute = defineWrapperRoute({
  path: '/schemas/:schemaId',
  children: [schemaEditorRoute],
});

export type SchemasRouteInfo = InferWrapperRouteInfo<typeof schemasRoute, SchemaEditorRoute>;
