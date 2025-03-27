import type { Component } from 'vue';
import type { IAppRoute } from './IAppRoute';

export interface IAppViewRoute extends IAppRoute {
  component: () => Promise<Component>;
}
