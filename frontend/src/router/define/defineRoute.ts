import type { Component } from 'vue';
import type { RouteParamValueRaw, RouteRecordInfo } from 'vue-router';
import type { ComponentProps } from '@/types';

export interface IAppRoute {
  path: string;
}

export interface IAppViewRoute extends IAppRoute {
  name: string;
  component: () => Promise<{ default: Component }>;
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

type InferRouteProps<P> = {
  [K in keyof P]: P[K] extends RouteParamValueRaw ? P[K] : never;
};

type InferRouteNormalizedProps<P> = {
  [K in keyof P]: P[K] extends RouteParamValueRaw ? string : never;
};

export type InferViewRouteInfo<
  R extends IAppViewRoute,
  P = ComponentProps<Awaited<ReturnType<R['component']>>['default']>,
> = Record<R['name'], RouteRecordInfo<
  R['name'],
  R['path'],
  InferRouteProps<P>,
  InferRouteNormalizedProps<P>
>>;

export type InferWrapperRouteInfo<R extends IAppWrapperRoute, M extends Record<string, RouteRecordInfo>> = {
  [N in keyof M]: Omit<M[N], 'path'> & { path: `${R['path']}/${M[N]['path']}` }
};
