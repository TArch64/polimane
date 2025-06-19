export type Default<V, D extends V> = undefined extends V ? D : V;
