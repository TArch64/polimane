import type { IPoint } from './Point';

export type BeadCoord = `${number}:${number}`;

export function serializeBeadCoord(x: number, y: number): BeadCoord {
  return `${x}:${y}`;
}

export function serializeBeadPoint(point: IPoint): BeadCoord {
  return serializeBeadCoord(point.x, point.y);
}

export function parseBeadCoord(coord: string): IPoint {
  const [x, y] = coord.split(':').map(Number);
  return { x: x!, y: y! };
}
