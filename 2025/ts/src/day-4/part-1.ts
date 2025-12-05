import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-4/part-1/input.txt');

function checkNeighbors(lines: string[], i: number, j: number): number {
  const minI = 0;
  const maxI = lines.length - 1;
  const minJ = 0;
  const maxJ = lines[0].length - 1;
  let neighborCount = 0;

  for (let y = i - 1; y <= i + 1; y++) {
    if (y < minI || y > maxI) {
      continue;
    }
    for (let x = j - 1; x <= j + 1; x++) {
      if (x < minJ || x > maxJ) {
        continue;
      }
      if (y === i && x === j) {
        continue;
      }
      if (lines[y][x] === '@') {
        neighborCount++;
      }
    }
  }

  return neighborCount;
}

function checkLine(lines: string[], checkIdx: number): number {
  let canAccessCount = 0;
  for (let i = 0; i < lines.length; i++) {
    if (i !== checkIdx) {
      continue;
    }
    for (let j = 0; j < lines[i].length; j++) {
      if (lines[i][j] !== '@') {
        continue;
      }
      const neighborsCount = checkNeighbors(lines, i, j);
      if (neighborsCount < 4) {
        canAccessCount++;
      }
    }
  }
  return canAccessCount;
}

export async function day4Part1(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    let solution = 0;
    let lines: string[] = [];

    for await (const line of rl) {
      const trimmedLine = line.trim();
      lines.push(trimmedLine);

      if (lines.length === 2) {
        solution += checkLine(lines, 0);
      }

      if (lines.length === 3) {
        solution += checkLine(lines, 1);
        lines = lines.slice(1);
      }
    }

    solution += checkLine(lines, 1);

    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}