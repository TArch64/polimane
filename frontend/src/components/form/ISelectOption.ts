export type SelectValue = string;

export interface ISelectOption<V extends SelectValue> {
  value: V;
  label: string;
  disabled?: boolean;
}

export type SelectOptions<V extends SelectValue> = ISelectOption<V>[];
