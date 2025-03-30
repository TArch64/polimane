import type { IAppViewRoute } from './IAppViewRoute';

export const defineRoute = <const R extends IAppViewRoute>(route: R): R => route;
