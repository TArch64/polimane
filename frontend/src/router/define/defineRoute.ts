import type { Component } from 'vue';
import type { RouteRecordInfo } from 'vue-router';

export interface IAppRoute {
  path: string;
}

export interface IAppViewRoute extends IAppRoute {
  name: string;
  component: () => Promise<Component>;
}

export interface IAppWrapperRoute extends IAppRoute {
  children: AppRoute[];
}

export type AppRoute = IAppViewRoute | IAppWrapperRoute;

export function defineRoute<const R extends IAppViewRoute>(route: R): R;
export function defineRoute<const R extends IAppWrapperRoute>(route: R): R;
export function defineRoute<const R extends AppRoute>(route: R): R {
  return { ...route, props: true };
}

export type InferViewRouteInfo<R extends IAppViewRoute> = Record<R['name'], RouteRecordInfo<
  R['name'],
  R['path'],
  Record<never, never>,
  Record<never, never>
>>;

export type InferWrapperRouteInfo<R extends IAppWrapperRoute, M extends Record<string, RouteRecordInfo>> = {
  [N in keyof M]: Omit<M[N], 'path'> & { path: `${R['path']}/${M[N]['path']}` }
};
