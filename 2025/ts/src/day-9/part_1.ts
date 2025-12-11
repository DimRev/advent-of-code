import * as fs from 'fs';
import * as readline from 'readline';
import { PATH } from './common';

export async function day9Part1(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });


    for await (const line of rl) {
      const trimmedLine = line.trim();
      console.log(trimmedLine);
    }

    const solution = 0;
    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}