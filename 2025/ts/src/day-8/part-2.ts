import * as fs from 'fs';
import * as readline from 'readline';
import {
  PATH,
  Point,
  Segment,
  parseLineIntoPoint,
  parsePointsIntoSegments,
  sortSegmentsByLength,
  islandSizesP2,
} from './common';

function connectAllPairs(segments: Segment[], points: Point[]): number {
  for (let i = 0; i < segments.length; i++) {
    const segment = segments[i];
    segment.start.connect(segment.end);

    const islands = islandSizesP2(points);

    if (islands.length === 1) {
      return segment.start.x * segment.end.x;
    }
  }

  return -1;
}

export async function day8Part2(): Promise<void> {
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
    const solution = connectAllPairs(segments, points);

    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}
