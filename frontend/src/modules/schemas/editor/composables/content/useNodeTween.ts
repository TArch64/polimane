import { type MaybeRefOrGetter, toValue, watch, type WatchSource } from 'vue';
import Konva from 'konva';
import { isEqual } from '@/helpers';

type InferWatchValue<S> = S extends WatchSource<infer T> ? T : never;
type NodeTweenFactory<S extends WatchSource> = (value: InferWatchValue<S>) => Omit<Konva.TweenConfig, 'node'> | null;

export function useNodeTween<S extends WatchSource>(
  nodeRef: MaybeRefOrGetter<Konva.Node | null>,
  source: S,
  buildTweenConfig: NodeTweenFactory<S>,
): void {
  watch(source, (value, oldValue) => {
    if (isEqual(value, oldValue)) {
      return;
    }

    const tweenConfig = buildTweenConfig(value);

    if (tweenConfig) {
      toValue(nodeRef)?.to(tweenConfig);
    }
  });
}
