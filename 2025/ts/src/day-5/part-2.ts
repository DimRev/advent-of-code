import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-5/part-1/input.txt');

type Range = {
  start: number;
  end: number;
};

function parseRange(rangeStr: string): Range {
  const parts = rangeStr.split('-');
  if (parts.length !== 2) {
    throw new Error(`Invalid range format: ${rangeStr}`);
  }
  const start = parseInt(parts[0]);
  const end = parseInt(parts[1]);
  return { start, end };
}

function squashRanges(ranges: Range[]): Range[] {
  if (ranges.length === 0) {
    return ranges;
  }

  ranges.sort((a, b) => a.start - b.start);

  const result: Range[] = [ranges[0]];

  for (let i = 1; i < ranges.length; i++) {
    const r = ranges[i];
    const last = result[result.length - 1];
    if (r.start <= last.end + 1) {
      last.end = Math.max(last.end, r.end);
    } else {
      result.push(r);
    }
  }

  return result;
}

function countTotalIDs(ranges: Range[]): number {
  let total = 0;
  for (const r of ranges) {
    total += r.end - r.start + 1;
  }
  return total;
}

export async function day5Part2(): Promise<void> {
  try {
        const fileStream = fs.createReadStream(PATH);
        const rl = readline.createInterface({
          input: fileStream,
          crlfDelay: Infinity,
        });

        let solution = 0;
        const ranges: Range[] = [];


        for await (const line of rl) {
          const trimmedLine = line.trim();
          if (trimmedLine === '') {
            break;
          }
          const range = parseRange(trimmedLine);
          ranges.push(range);
        }

        const squashed = squashRanges(ranges);
        solution = countTotalIDs(squashed);

        console.log(`Solution: ${solution}`);
      } catch (error) {
        console.error('Error:', error);
        process.exit(1);
      }
}