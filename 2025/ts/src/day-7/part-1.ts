import * as fs from 'fs';
import * as readline from 'readline';
import { PATH } from './common';

function generateBeamList(line: string): boolean[] {
  const result: boolean[] = [];
  for (const char of line) {
    result.push(char === 'S');
  }
  return result;
}

function progressSplitters(line: string, beamList: boolean[]): [boolean[], number] {
  const newBeamList: boolean[] = new Array(line.length).fill(false);
  let splitCount = 0;

  for (let i = 0; i < line.length; i++) {
    const char = line[i];
    if (char === '^' && beamList[i]) {
      if (i - 1 >= 0) {
        newBeamList[i - 1] = true;
      }
      if (i + 1 < line.length) {
        newBeamList[i + 1] = true;
      }
      newBeamList[i] = false;
      splitCount++;
      continue;
    }
    newBeamList[i] = newBeamList[i] || beamList[i];
  }

  return [newBeamList, splitCount];
}

function formatBeamList(beamList: boolean[]): string {
  let result = '';
  for (const b of beamList) {
    result += b ? '|' : '.';
  }
  return result;
}

export async function day7Part1(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    const shouldPrint = false;
    let splitCount = 0;
    let beamList: boolean[] = [];
    let idx = 0;

    for await (const line of rl) {
      const trimmedLine = line.trim();

      if (idx === 0) {
        beamList = generateBeamList(trimmedLine);
        if (shouldPrint) {
          console.log(trimmedLine);
          console.log(formatBeamList(beamList));
        }
      } else if (idx % 2 === 0) {
        const [newBeamList, newSplitCount] = progressSplitters(trimmedLine, beamList);
        splitCount += newSplitCount;
        beamList = newBeamList;
        if (shouldPrint) {
          console.log(trimmedLine);
          console.log(formatBeamList(beamList));
        }
      }

      idx++;
    }

    const solution = splitCount;
    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}