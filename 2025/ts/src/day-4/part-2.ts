import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-4/part-1/input.txt');

function checkNeighborsP2(lines: string[], i: number, j: number): number {
  const minI = 0;
  const maxI = lines.length - 1;
  const minJ = 0;
  let neighborCount = 0;

  for (let y = i - 1; y <= i + 1; y++) {
    if (y < minI || y > maxI) {
      continue;
    }
    const maxJY = lines[y].length - 1;
    for (let x = j - 1; x <= j + 1; x++) {
      if (x < minJ || x > maxJY) {
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

function checkLineP2(lines: string[]): [number, string[]] {
  let canAccessCount = 0;
  const copy = lines.slice();

  const toRemove = new Set<string>();
  for (let i = 0; i < lines.length; i++) {
    for (let j = 0; j < lines[i].length; j++) {
      if (lines[i][j] !== '@') {
        continue;
      }
      const neighborsCount = checkNeighborsP2(lines, i, j);
      if (neighborsCount < 4) {
        toRemove.add(`${i},${j}`);
        canAccessCount++;
      }
    }
  }

  for (const coord of toRemove) {
    const [i, j] = coord.split(',').map(Number);
    const lineChars = copy[i].split('');
    lineChars[j] = 'x';
    copy[i] = lineChars.join('');
  }

  return [canAccessCount, copy];
}

export async function day4Part2(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    let lines: string[] = [];

    for await (const line of rl) {
      const trimmedLine = line.trim();
      lines.push(trimmedLine);
    }

    // Filter out empty lines
    lines = lines.filter(line => line !== '');

    let totalRemoved = 0;
    while (true) {
      const [validCount, newLines] = checkLineP2(lines);
      if (validCount === 0) {
        break;
      }
      totalRemoved += validCount;
      lines = newLines;
    }

    console.log(`Solution: ${totalRemoved}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}