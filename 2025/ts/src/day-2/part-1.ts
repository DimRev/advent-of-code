import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const PATH = path.join(__dirname, '../../..', 'inputs/day-2/part-1/input.txt');
const COMMA_SEPARATOR = ',';
const RANGE_SEPARATOR = '-';
const LEADING_ZERO_CHAR = '0';

function checkIfInvalid(num: number): boolean {
  const numStr = num.toString();
  const length = numStr.length;

  if (length % 2 !== 0) {
    return false;
  }

  if (numStr[0] === LEADING_ZERO_CHAR) {
    return false;
  }

  const half = Math.floor(length / 2);
  const firstHalf = numStr.slice(0, half);
  const secondHalf = numStr.slice(half);

  return firstHalf === secondHalf;
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

export async function day2Part1(): Promise<void> {
  const fileStream = fs.createReadStream(PATH);
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  try {
    let solution = 0;
    for await (const line of rl) {
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
