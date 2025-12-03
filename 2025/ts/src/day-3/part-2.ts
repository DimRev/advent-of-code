import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';


const PATH = path.join(__dirname, '../../..', 'inputs/day-3/part-1/input.txt');

function maxPairD2(line:string): number {
  let maxIdx = -1;
  let result = 0;

  for (let digitPos = 11; digitPos >= 0; digitPos--) {
    let maxDigit = -1;

    for (let i = maxIdx + 1; i < line.length - digitPos; i++) {
      const num = parseInt(line[i], 10);
      if (num > maxDigit) {
        maxDigit = num;
        maxIdx = i;
      }
    }

    result = result * 10 + maxDigit;
  }

  return result;
}

export async function day3Part2(): Promise<void> {
  const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    try {
      let solution = 0;
      for await (const line of rl) {
        const maxPair = maxPairD2(line.trim());
        solution += maxPair;
      }
      console.log(`Solution: ${solution}`);
    } finally {
      rl.close();
    }
}