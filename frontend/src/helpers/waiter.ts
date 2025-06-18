export type WaiterRelease = () => void;

export interface IWaiter {
  add: () => WaiterRelease;
  wait: () => Promise<void>;
}

export function createWaiter(): IWaiter {
  const waiters: Promise<void>[] = [];

  function add(): () => void {
    let resolve: () => void;
    const promise = new Promise<void>((res) => resolve = res);
    waiters.push(promise);
    return resolve!;
  }

  async function wait(): Promise<void> {
    if (waiters.length) await Promise.all(waiters);
  }

  return { add, wait };
}
