import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';


const PATH = path.join(__dirname, '../../..', 'inputs/day-3/part-1/input.txt');

function maxPairD1(line:string): number {
  let firstMaxDigit = -1;
  let firstMaxDigitIndex = -1;

  for (let i = 0; i < line.length - 1; i++) {
    const num = parseInt(line[i], 10);
    if (num > firstMaxDigit) {
      firstMaxDigit = num;
      firstMaxDigitIndex = i;
    }
  }

  let secondMaxDigit = -1;

  for (let i = firstMaxDigitIndex + 1; i < line.length; i++) {
    const num = parseInt(line[i], 10);
    if (num > secondMaxDigit) {
      secondMaxDigit = num;
    }
  }

  return (firstMaxDigit * 10) + secondMaxDigit;
}

export async function day3Part1(): Promise<void> {
  const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    try {
      let solution = 0;
      for await (const line of rl) {
        const maxPair = maxPairD1(line.trim());
        solution += maxPair;
      }
      console.log(`Solution: ${solution}`);
    } finally {
      rl.close();
    }
}