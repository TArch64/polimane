import { defineRoute, type InferWrapperRouteInfo } from '@/router/define';
import { schemaEditorRoute, type SchemaEditorRoute } from './editor';

export const schemasRoute = defineRoute({
  path: '/schemas/:schemaId',
  children: [schemaEditorRoute],
});

export type SchemasRouteInfo = InferWrapperRouteInfo<typeof schemasRoute, SchemaEditorRoute>;
