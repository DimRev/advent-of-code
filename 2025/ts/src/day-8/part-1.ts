import * as fs from 'fs';
import * as readline from 'readline';
import {
  PATH,
  Point,
  Segment,
  parseLineIntoPoint,
  parsePointsIntoSegments,
  sortSegmentsByLength,
  islandSizes,
} from './common';

function connectPairs(segments: Segment[], numOfConnections: number): void {
  for (let i = 0; i < numOfConnections; i++) {
    const segment = segments[i];
    segment.start.connect(segment.end);
  }
}

function sortIslandsBySize(islandSizes: number[]): number[] {
  return islandSizes.sort((a, b) => b - a);
}

function productLargestIslands(islandSizes: number[], num: number): number {
  let product = 1;
  for (let i = 0; i < num; i++) {
    const size = islandSizes[i];
    product *= size;
  }
  return product;
}

export async function day8Part1(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    let runningId = 0;
    const points: Point[] = [];

    for await (const line of rl) {
      const trimmedLine = line.trim();
      const point = parseLineIntoPoint(trimmedLine, runningId);
      points.push(point);
      runningId++;
    }

    const segments = sortSegmentsByLength(parsePointsIntoSegments(points));
    connectPairs(segments, 1000);
    const islands = sortIslandsBySize(islandSizes(points));
    const solution = productLargestIslands(islands, 3);

    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}
