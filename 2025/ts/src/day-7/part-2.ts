import * as fs from 'fs';
import * as readline from 'readline';
import { PATH } from './common';

function generatePathCounter(line: string): number[] {
  const result: number[] = [];
  for (const char of line) {
    if (char === 'S') {
      result.push(1);
    } else {
      result.push(0);
    }
  }
  return result;
}

function progressSplittersDP(line: string, pathCounter: number[]): number[] {
  const newPathCounter: number[] = new Array(line.length).fill(0);

  for (let i = 0; i < line.length; i++) {
    if (pathCounter[i] === 0) {
      continue;
    }

    const char = line[i];
    if (char === '^') {
      if (i - 1 >= 0) {
        newPathCounter[i - 1] += pathCounter[i];
      }
      if (i + 1 < line.length) {
        newPathCounter[i + 1] += pathCounter[i];
      }
    } else {
      newPathCounter[i] += pathCounter[i];
    }
  }

  return newPathCounter;
}

export async function day7Part2(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    let pathCounter: number[] = [];
    let idx = 0;

    for await (const line of rl) {
      const trimmedLine = line.trim();

      if (idx === 0) {
        pathCounter = generatePathCounter(trimmedLine);
      } else if (idx % 2 === 0) {
        pathCounter = progressSplittersDP(trimmedLine, pathCounter);
      }

      idx++;
    }

    const solution = pathCounter.reduce((sum, count) => sum + count, 0);
    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}