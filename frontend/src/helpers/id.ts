export const newId = (): string => crypto.randomUUID();

export function parseIdBytes(id: string): Uint8Array {
  return new Uint8Array(
    id
      .replace(/-/g, '')
      .match(/.{2}/g)!
      .map((b) => parseInt(b, 16)),
  );
}
