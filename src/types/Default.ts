export type Default<V, D extends V> = V extends undefined ? D : V;
