export function isEqual(v1: unknown, v2: unknown): boolean {
  return JSON.stringify(v1) === JSON.stringify(v2);
}
