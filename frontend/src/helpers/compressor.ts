export async function compressObject(obj: object): Promise<Blob> {
  const blob = new Blob([JSON.stringify(obj)], { type: 'application/json' });
  const compressionStream = new CompressionStream('deflate-raw');
  const compressedStream = blob.stream().pipeThrough(compressionStream);
  return new Response(compressedStream).blob();
}

export function decompressObject<O extends object>(blob: Blob): Promise<O> {
  const decompressionStream = new DecompressionStream('deflate-raw');
  const decompressedStream = blob.stream().pipeThrough(decompressionStream);
  return new Response(decompressedStream).json() as Promise<O>;
}
