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

function isIDInRanges(ranges: Range[], id: number): boolean {
  for (const r of ranges) {
    if (id >= r.start && id <= r.end) {
      return true;
    }
  }
  return false;
}

export async function day5Part1(): Promise<void> {
  try {
      const fileStream = fs.createReadStream(PATH);
      const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
      });

      let solution = 0;
      const ranges: Range[] = [];
      const ids: number[] = [];
      let isRangeSection = true;


      for await (const line of rl) {
        const trimmedLine = line.trim();
        if (trimmedLine === '') {
          isRangeSection = false;
          continue;
        }
        if (isRangeSection) {
          const range = parseRange(trimmedLine);
          ranges.push(range);
        } else {
          const id = parseInt(trimmedLine);
          ids.push(id);
        }
      }

      ranges.sort((a, b) => a.start - b.start);

      for (const id of ids) {
        if (isIDInRanges(ranges, id)) {
          solution++;
        }
      }


      console.log(`Solution: ${solution}`);
    } catch (error) {
      console.error('Error:', error);
      process.exit(1);
    }
}