import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-6/part-1/input.txt');

function mapTransformNums(slice: string[]): number[] {
  const result: number[] = [];
  for (const v of slice) {
    if (v === '') {
      continue;
    }
    const num = parseInt(v, 10);
    if (isNaN(num)) {
      continue;
    }
    result.push(num);
  }
  return result;
}

function mapCleanSymbols(slice: string[]): string[] {
  const result: string[] = [];
  for (const v of slice) {
    if (v === '') {
      continue;
    }
    result.push(v);
  }
  return result;
}

function addition(nums: number[]): number {
  let total = 0;
  for (const num of nums) {
    total += num;
  }
  return total;
}

function multiplication(nums: number[]): number {
  let total = 1;
  for (const num of nums) {
    total *= num;
  }
  return total;
}

function calculateTotal(numsList: number[][], symbols: string[], lineLen: number): number {
  let total = 0;
  for (let i = 0; i < lineLen; i++) {
    const currNums: number[] = [];
    for (const nums of numsList) {
      currNums.push(nums[i]);
    }
    switch (symbols[i]) {
      case '+':
        total += addition(currNums);
        break;
      case '*':
        total += multiplication(currNums);
        break;
      default:
        console.error(`Unknown symbol: ${symbols[i]}`);
        process.exit(1);
    }
  }
  return total;
}

export async function day6Part1(): Promise<void> {
  try {
    const fileStream = fs.createReadStream(PATH);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    const numsList: number[][] = [];
    let symbols: string[] = [];
    let lineLen = 0;

    for await (const line of rl) {
      const trimmedLine = line.trim();

      const parsed = trimmedLine.split(' ');
      const transformed = mapTransformNums(parsed);
      if (transformed.length !== 0) {
        numsList.push(transformed);
      } else {
        const cleanedSymbols = mapCleanSymbols(parsed);
        symbols = cleanedSymbols;
        lineLen = cleanedSymbols.length;
      }
    }

    const solution = calculateTotal(numsList, symbols, lineLen);

    console.log(`Solution: ${solution}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}