export class ObjectAnimationFrame {
  private frameId: number | null = null;

  request(callback: () => void): void {
    if (this.frameId) {
      cancelAnimationFrame(this.frameId);
    }

    this.frameId = requestAnimationFrame(() => {
      callback();
      this.frameId = null;
    });
  }

  dispose(): void {
    if (this.frameId) {
      cancelAnimationFrame(this.frameId);
      this.frameId = null;
    }
  }
}
