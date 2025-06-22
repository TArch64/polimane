export interface ISelectOption<V extends string> {
  value: V;
  label: string;
  disabled?: boolean;
}
