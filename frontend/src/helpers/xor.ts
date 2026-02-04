export function decryptXor(key: Uint8Array, encrypted: Uint8Array): Uint8Array {
  const result = new Uint8Array(encrypted.length);
  for (let i = 0; i < encrypted.length; i++) {
    result[i] = encrypted[i]! ^ key[i % key.length]!;
  }
  return result;
}
