import type { Component, IntrinsicElementAttributes } from 'vue';

export type ComponentTag = keyof IntrinsicElementAttributes & string;
export type ComponentAs = ComponentTag | Component;
