import * as path from 'path';

export const PATH = path.join(__dirname, '../../..', 'inputs/day-8/part-1/input.txt');

export class Point {
  id: number;
  x: number;
  y: number;
  z: number;
  connectionsMap: Map<number, Point>;

  constructor(x: number, y: number, z: number, id: number) {
    this.id = id;
    this.x = x;
    this.y = y;
    this.z = z;
    this.connectionsMap = new Map();
  }

  toString(): string {
    const connections = this.listConnectionsIds();
    return `Point[${this.id}](${this.x}, ${this.y}, ${this.z}) Connections: ${connections}`;
  }

  distance(target: Point): number {
    const diffXSquared = Math.pow(target.x - this.x, 2);
    const diffYSquared = Math.pow(target.y - this.y, 2);
    const diffZSquared = Math.pow(target.z - this.z, 2);
    return Math.floor(Math.abs(Math.sqrt(diffXSquared + diffYSquared + diffZSquared)));
  }

  connect(target: Point): void {
    if (!this.connectionsMap.has(target.id)) {
      this.connectionsMap.set(target.id, target);
    }

    if (!target.connectionsMap.has(this.id)) {
      target.connectionsMap.set(this.id, this);
    }
  }

  listConnectionsIds(): number[] {
    return Array.from(this.connectionsMap.keys());
  }
}

export class Segment {
  start: Point;
  end: Point;
  length: number;

  constructor(start: Point, end: Point, length: number) {
    this.start = start;
    this.end = end;
    this.length = length;
  }

  toString(): string {
    return `Segment(Point[${this.start.id}], Point[${this.end.id}], ${this.length})`;
  }
}

export function generatePoint(x: number, y: number, z: number, id: number): Point {
  return new Point(x, y, z, id);
}

export function parseLineIntoPoint(line: string, id: number): Point {
  const parts = line.split(',');
  if (parts.length !== 3) {
    throw new Error(`Error parsing line: ${line}`);
  }

  const x = parseInt(parts[0], 10);
  const y = parseInt(parts[1], 10);
  const z = parseInt(parts[2], 10);

  if (isNaN(x) || isNaN(y) || isNaN(z)) {
    throw new Error(`Error parsing line: ${line}`);
  }

  return generatePoint(x, y, z, id);
}

export function parsePointsIntoSegments(points: Point[]): Segment[] {
  const segments: Segment[] = [];
  for (let i = 0; i < points.length - 1; i++) {
    for (let j = i + 1; j < points.length; j++) {
      const segment = new Segment(
        points[i],
        points[j],
        points[j].distance(points[i])
      );
      segments.push(segment);
    }
  }
  return segments;
}

export function sortSegmentsByLength(segments: Segment[]): Segment[] {
  return segments.sort((a, b) => a.length - b.length);
}

export function traverseConnections(
  point: Point,
  prevCount: number,
  visitedIds: Map<number, boolean>
): number {
  if (visitedIds.has(point.id)) {
    return 0;
  }

  if (point.connectionsMap.size === 0) {
    return 1;
  }

  let count = 1;
  visitedIds.set(point.id, true);

  for (const p of point.connectionsMap.values()) {
    count += traverseConnections(p, prevCount + 1, visitedIds);
  }

  return count;
}

export function islandSizes(points: Point[]): number[] {
  const visitedIds = new Map<number, boolean>();
  const sizes: number[] = [];
  for (const p of points) {
    const size = traverseConnections(p, 0, visitedIds);
    if (size > 0) {
      sizes.push(size);
    }
  }
  return sizes;
}

export function islandSizesP2(points: Point[]): number[] {
  const visitedIds = new Map<number, boolean>();
  const sizes: number[] = [];
  for (const p of points) {
    const size = traverseConnections(p, 0, visitedIds);
    if (size > 0) {
      sizes.push(size);
    }
    if (sizes.length > 1) {
      return [];
    }
  }
  return sizes;
}
