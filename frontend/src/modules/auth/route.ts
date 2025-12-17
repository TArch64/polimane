import { defineWrapperRoute, type InferWrapperRouteInfo } from '@/router/define';
import { authLoginRoute, type AuthLoginRouteInfo } from './login';
import { authCompleteRoute, type AuthCompleteRouteInfo } from './complete';
import { authDeletedUserRoute, type AuthDeletedUserRouteInfo } from './deletedUser';

export const authRoute = defineWrapperRoute({
  path: '/auth',

  children: [
    authLoginRoute,
    authCompleteRoute,
    authDeletedUserRoute,
  ],
});

export type AuthRouteInfo = InferWrapperRouteInfo<typeof authRoute, AuthLoginRouteInfo & AuthCompleteRouteInfo & AuthDeletedUserRouteInfo>;
