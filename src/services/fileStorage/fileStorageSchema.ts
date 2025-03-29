import {
  array,
  check,
  endsWith,
  type InferInput,
  type InferOutput,
  integer,
  maxLength,
  minLength,
  number,
  object,
  pipe,
  string,
  transform,
  value,
} from 'valibot';

const fileStorageTimestamp = () => pipe(
  number(),
  integer(),
  transform((timestamp) => new Date(timestamp)),
  check((date) => !Number.isNaN(date) && date.getTime() > 0),
);

export const fileStorageEntrySchema = () => object({
  filename: pipe(
    string(),
    minLength(1, 'invalid filename'),
    endsWith('.plmn', 'invalid filename'),
  ),

  title: pipe(
    string(),
    minLength(1),
    maxLength(256),
  ),

  updatedAt: fileStorageTimestamp(),
});

export type FileStorageEntrySchema = ReturnType<typeof fileStorageEntrySchema>;
export type FileStorageEntryRaw = InferInput<FileStorageEntrySchema>;
export type FileStorageEntryData = InferOutput<FileStorageEntrySchema>;

export const fileStorageSchema = () => object({
  version: pipe(number(), integer(), value(1, 'Available versions: 1')),
  updatedAt: fileStorageTimestamp(),
  entries: array(fileStorageEntrySchema()),
});

export type FileStorageSchema = ReturnType<typeof fileStorageSchema>;
export type FileStorageRaw = InferInput<FileStorageSchema>;
export type FileStorageData = InferOutput<FileStorageSchema>;
