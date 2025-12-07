import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-6/part-1/input.txt');

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

async function readLines(filePath: string): Promise<string[]> {
  const fileStream = fs.createReadStream(filePath);
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  const lines: string[] = [];
  for await (const line of rl) {
    lines.push(line);
  }
  return lines;
}

function padLines(lines: string[]): string[] {
  let maxLen = 0;
  for (const line of lines) {
    if (line.length > maxLen) {
      maxLen = line.length;
    }
  }

  const padded: string[] = [];
  for (const line of lines) {
    if (line.length < maxLen) {
      padded.push(line + ' '.repeat(maxLen - line.length));
    } else {
      padded.push(line);
    }
  }
  return padded;
}

function isColumnEmpty(lines: string[], col: number): boolean {
  for (let row = 0; row < lines.length; row++) {
    if (col < lines[row].length && lines[row][col] !== ' ') {
      return false;
    }
  }
  return true;
}

function readNumberVertically(lines: string[], col: number, startRow: number): [number, number, boolean] {
  let numStr = '';
  let row = startRow;
  while (row < lines.length) {
    if (col >= lines[row].length) {
      break;
    }
    const ch = lines[row][col];
    if (ch >= '0' && ch <= '9') {
      numStr += ch;
      row++;
    } else {
      break;
    }
  }

  if (numStr === '') {
    return [0, startRow, false];
  }

  const num = parseInt(numStr, 10);
  if (isNaN(num)) {
    return [0, startRow, false];
  }

  return [num, row - 1, true];
}

function processProblem(lines: string[], startCol: number): [number[], string, number] {
  const numbers: number[] = [];
  let operator = '';
  let col = startCol;

  while (col >= 0) {
    if (isColumnEmpty(lines, col)) {
      break;
    }

    for (let row = 0; row < lines.length; row++) {
      if (col >= lines[row].length) {
        continue;
      }

      const char = lines[row][col];
      if (char === ' ') {
        continue;
      }

      if (char === '+' || char === '*') {
        operator = char;
      } else if (char >= '0' && char <= '9') {
        const [num, lastRow, success] = readNumberVertically(lines, col, row);
        if (success) {
          numbers.push(num);
          row = lastRow;
        }
      }
    }
    col--;
  }

  return [numbers, operator, col];
}

function calculateProblem(numbers: number[], operator: string): number {
  if (numbers.length === 0 || operator === '') {
    return 0;
  }

  switch (operator) {
    case '+':
      return addition(numbers);
    case '*':
      return multiplication(numbers);
    default:
      return 0;
  }
}

export async function day6Part2(): Promise<void> {
  try {
    const lines = await readLines(PATH);

    if (lines.length === 0) {
      console.log('No lines found');
      return;
    }

    const paddedLines = padLines(lines);
    let grandTotal = 0;
    let col = paddedLines[0].length - 1;

    while (col >= 0) {
      while (col >= 0 && isColumnEmpty(paddedLines, col)) {
        col--;
      }

      if (col < 0) {
        break;
      }

      const [numbers, operator, nextCol] = processProblem(paddedLines, col);
      const result = calculateProblem(numbers, operator);
      grandTotal += result;
      col = nextCol;
    }

    console.log(`Solution: ${grandTotal}`);
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}