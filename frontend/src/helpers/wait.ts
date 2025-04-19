export function wait(time: number): Promise<void> {
  return new Promise((resolve) => {
    setTimeout(() => resolve(), time);
  });
}

/**
 * Wait for a click event to complete
 * Required when you want to add click listener in another click listeners
 * Otherwise without timeout new listener will be called immediately
 */
export const waitClickComplete = () => wait(33);
