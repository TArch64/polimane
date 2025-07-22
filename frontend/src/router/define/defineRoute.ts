import type { Component } from 'vue';
import type { RouteParamValueRaw, RouteRecordInfo } from 'vue-router';
import type { InferComponentProps, SafeAny } from '@/types';

export interface IAppRoute {
  path: string;
}

export interface IAppViewRoute extends IAppRoute {
  name: string;
  component: () => Promise<{ default: Component }>;
}

export interface IAppWrapperRoute extends IAppRoute {
  name?: string;
  children: Array<IAppViewRoute | IAppWrapperRoute | IAppRedirectRoute>;
}

export interface IAppRedirectRoute extends IAppRoute {
  redirect: SafeAny;
}

export function defineViewRoute<const R extends IAppViewRoute>(route: R): R {
  return { ...route, props: true };
}

export function defineWrapperRoute<const R extends IAppWrapperRoute>(route: R): R {
  return { ...route, props: true };
}

export function defineRedirectRoute(path: string, redirect: string): IAppRedirectRoute {
  return { path, redirect: { name: redirect } };
}

type InferRouteProps<P> = {
  [K in keyof P]: P[K] extends RouteParamValueRaw ? P[K] : never;
};

type InferRouteNormalizedProps<P> = {
  [K in keyof P]: P[K] extends RouteParamValueRaw ? string : never;
};

export type InferViewRouteInfo<
  R extends IAppViewRoute,
  P = InferComponentProps<Awaited<ReturnType<R['component']>>['default']>,
> = Record<R['name'], RouteRecordInfo<
  R['name'],
  R['path'],
  InferRouteProps<P>,
  InferRouteNormalizedProps<P>
>>;

export type InferWrapperRouteInfo<R extends IAppWrapperRoute, M extends Record<string, RouteRecordInfo>> = {
  [N in keyof M]: Omit<M[N], 'path'> & { path: `${R['path']}/${M[N]['path']}` }
};
