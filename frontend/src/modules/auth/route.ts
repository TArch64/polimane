import { defineRoute, type InferWrapperRouteInfo } from '@/router/define';
import { authLoginRoute, type AuthLoginRouteInfo } from './login';
import { authCompleteRoute, type AuthCompleteRouteInfo } from './complete';

export const authRoute = defineRoute({
  path: '/auth',
  children: [authLoginRoute, authCompleteRoute],
});

export type AuthRouteInfo = InferWrapperRouteInfo<typeof authRoute, AuthLoginRouteInfo & AuthCompleteRouteInfo>;
