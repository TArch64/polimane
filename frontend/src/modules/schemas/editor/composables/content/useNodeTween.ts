import { type MaybeRefOrGetter, onWatcherCleanup, toValue, watch, type WatchSource } from 'vue';
import Konva from 'konva';

type InferWatchValue<S> = S extends WatchSource<infer T> ? T : never;
type NodeTweenFactory<S extends WatchSource> = (value: InferWatchValue<S>) => Omit<Konva.TweenConfig, 'node'> | null;

export function useNodeTween<S extends WatchSource>(
  nodeRef: MaybeRefOrGetter<Konva.Node | null>,
  source: S,
  buildTweenConfig: NodeTweenFactory<S>,
): void {
  let tween: Konva.Tween | null = null;

  watch(source, (value) => {
    if (tween) {
      tween.destroy();
    }

    const node = toValue(nodeRef);

    if (!node) {
      return;
    }

    const tweenConfig = buildTweenConfig(value);

    if (!tweenConfig) {
      return;
    }

    tween = new Konva.Tween({
      node,
      ...tweenConfig,

      onFinish() {
        tween = null;
        tweenConfig.onFinish?.();
      },
    });

    tween.play();

    onWatcherCleanup(() => {
      tween?.destroy();
      tween = null;
    });
  });
}
