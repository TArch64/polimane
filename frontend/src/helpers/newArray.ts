export function newArray<I>(size: number, filler: (index: number) => I): I[] {
  return new Array(size).fill(0).map((_, index) => filler(index));
}
