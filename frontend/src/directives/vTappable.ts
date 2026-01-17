import type { DirectiveBinding, ObjectDirective } from 'vue';

export interface ITappableValue {
  disabled?: boolean;
}

class VTappable implements ObjectDirective<HTMLElement, ITappableValue> {
  private abortController: AbortController | null = null;

  mounted = (el: HTMLElement, binding: DirectiveBinding<ITappableValue>) => {
    if (this.toEnabled(binding)) this.attach(el);
  };

  updated = (el: HTMLElement, binding: DirectiveBinding<ITappableValue>) => {
    if (this.isAttached) {
      this.detach();
    }
    if (this.toEnabled(binding)) {
      this.attach(el);
    }
  };

  get isAttached(): boolean {
    return !!this.abortController;
  }

  private toEnabled(binding: DirectiveBinding<ITappableValue>): boolean {
    return binding.value?.disabled !== true;
  }

  private attach(el: HTMLElement) {
    this.abortController = new AbortController();

    el.addEventListener('mousedown', (downEvent) => {
      if (downEvent.button !== 0) return;

      this.configureAnimation(el);
      this.setAnimationClasses(el, 'in');

      window.addEventListener('mouseup', (upEvent) => {
        upEvent.button === 0
          ? this.setAnimationClasses(el, 'out')
          : this.cleanUp(el);
      }, { once: true, capture: true });
    }, {
      capture: true,
      signal: this.abortController.signal,
    });

    el.addEventListener('animationend', (event) => {
      if (event.animationName === 'tap-out') {
        this.cleanUp(el);
      }
    });
  }

  private detach() {
    this.abortController?.abort();
    this.abortController = null;
  }

  private configureAnimation(el: HTMLElement) {
    el.style.setProperty('--tap-scale', this.getTapScale(el));
  }

  private getTapScale(el: HTMLElement): string {
    const width = el.offsetWidth;
    const tapScaleDiff = Math.min(width * 0.07, 4);
    const tapScale = (width - tapScaleDiff) / width;
    return tapScale.toString();
  }

  private setAnimationClasses(el: HTMLElement, state: 'in' | 'out') {
    el.classList.toggle('tap-animation-in', state === 'in');
    el.classList.toggle('tap-animation-out', state === 'out');
  }

  private cleanUp(el: HTMLElement) {
    el.classList.remove('tap-animation-out', 'tap-animation-in');
    el.style.removeProperty('--tap-scale');
  }
}

export const vTappable = new VTappable();
