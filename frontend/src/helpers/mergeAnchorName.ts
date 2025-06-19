type StyleInput = Record<string, string> | Record<string, string>[];

export function mergeAnchorName(...styles: StyleInput[]): Record<string, string>[] {
  const anchorNames: string[] = [];

  const transformedStyles = styles.flat().map(({ anchorName, ...style }) => {
    if (anchorName) {
      anchorNames.push(anchorName);
    }

    return style;
  });

  return anchorNames.length
    ? [{ anchorName: anchorNames.join(', ') }, ...transformedStyles]
    : transformedStyles;
}
