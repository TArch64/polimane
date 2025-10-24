export type PascalToKebab<S extends string> = S extends `${infer First}${infer Rest}`
  ? Rest extends Uncapitalize<Rest>
    ? `${Lowercase<First>}${PascalToKebab<Rest>}`
    : `${Lowercase<First>}-${PascalToKebab<Uncapitalize<Rest>>}`
  : S;
