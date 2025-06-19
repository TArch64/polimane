export type WaiterRelease = () => void;

export interface IWaiter {
  add: () => WaiterRelease;
  wait: () => Promise<void>;
}

export function createWaiter(): IWaiter {
  const waiters: Promise<void>[] = [];

  function add(): () => void {
    const { promise, resolve } = Promise.withResolvers<void>();
    waiters.push(promise);
    return resolve;
  }

  async function wait(): Promise<void> {
    if (waiters.length) await Promise.all(waiters);
  }

  return { add, wait };
}
