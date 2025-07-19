export function newArray<I>(length: number, filler: (index: number) => I): I[] {
  return Array.from({ length }, (_, index) => filler(index));
}
