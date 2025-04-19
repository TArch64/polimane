type StyleInput = Record<string, string> | Record<string, string>[];

export function mergeAnchorName(...styles: StyleInput[]): Record<string, string>[] {
  const flatStyles = styles.flat();
  const anchorNames: string[] = [];

  for (const style of flatStyles) {
    if ('anchorName' in style) {
      anchorNames.push(style.anchorName);
      delete style['anchorName'];
    }
  }

  return anchorNames.length ? [{ anchorName: anchorNames.join(', ') }, ...flatStyles] : flatStyles;
}
