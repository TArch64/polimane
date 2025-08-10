export class ObjectCompressor<O extends object> {
  static _instance: ObjectCompressor<object>;

  static typed<O extends object>(): ObjectCompressor<O> {
    this._instance ??= new ObjectCompressor<object>();
    return this._instance as ObjectCompressor<O>;
  }

  async compress(obj: O): Promise<Blob> {
    const blob = new Blob([JSON.stringify(obj)], { type: 'application/json' });
    const compressionStream = new CompressionStream('deflate-raw');
    const compressedStream = blob.stream().pipeThrough(compressionStream as ReadableWritablePair);
    return new Response(compressedStream).blob();
  }

  decompress(blob: Blob): Promise<O> {
    const decompressionStream = new DecompressionStream('deflate-raw');
    const decompressedStream = blob.stream().pipeThrough(decompressionStream as ReadableWritablePair);
    return new Response(decompressedStream).json() as Promise<O>;
  }
}
