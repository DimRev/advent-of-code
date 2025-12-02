import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-2/part-1/input.txt');
const COMMA_SEPARATOR = ',';
const RANGE_SEPARATOR = '-';
const LEADING_ZERO_CHAR = '0';
const MIN_PATTERN_LEN = 1;

function isRepeatingPattern(numStr: string, patternLen: number): boolean {
  const pattern = numStr.slice(0, patternLen);

  for (let i = 0; i < numStr.length; i += patternLen) {
    if (numStr.slice(i, i + patternLen) !== pattern) {
      return false;
    }
  }
  return true;
}

function checkIfInvalid(num: number): boolean {
  const numStr = num.toString();
  const length = numStr.length;

  if (numStr[0] === '0') {
    return false;
  }

  const maxPatternLen = Math.floor(length / 2);
  for (
    let patternLen = MIN_PATTERN_LEN;
    patternLen <= maxPatternLen;
    patternLen++
  ) {
    if (length % patternLen !== 0) {
      continue;
    }

    if (isRepeatingPattern(numStr, patternLen)) {
      return true;
    }
  }

  return false;
}

function findInvalidNumsInRange(minNum: number, maxNum: number): number[] {
  const invalidNums = [];

  for (let num = minNum; num <= maxNum; num++) {
    if (checkIfInvalid(num)) {
      invalidNums.push(num);
    }
  }

  return invalidNums;
}

function parseRange(line: string): [number, number] {
  const nums = line.split(RANGE_SEPARATOR);
  if (nums.length !== 2) {
    return [0, 0];
  }

  return [parseInt(nums[0], 10), parseInt(nums[1], 10)];
}

export async function day2Part2(): Promise<void> {
  const fileStream = fs.createReadStream(PATH);
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  try {
    let solution = 0;
    for await (const line of rl) {
      // Split by comma separator
      for (const range of line.trim().split(COMMA_SEPARATOR)) {
        if (range.length <= 1) {
          continue;
        }
        const [minNum, maxNum] = parseRange(range);
        const invalidNums = findInvalidNumsInRange(minNum, maxNum);
        solution += invalidNums.reduce((acc, num) => acc + num, 0);
      }
    }
    console.log(`Solution: ${solution}`);
  } finally {
    rl.close();
  }
}
