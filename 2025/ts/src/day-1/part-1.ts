import * as fs from "fs";
import * as path from "path";
import * as readline from "readline";

const PATH = path.join(__dirname, "../../..", "inputs/day-1/part-1/input.txt");
const CIRCLE_SIZE = 100;
const START_POSITION = 50;

function parseLine(line: string): [string, number] {
  if (line.length < 2) {
    return ["", 0];
  }

  return [line[0], parseInt(line.slice(1), 10)];
}

function wrapPosition(position: number): number {
  return ((position % CIRCLE_SIZE) + CIRCLE_SIZE) % CIRCLE_SIZE;
}

export async function part1(): Promise<void> {
  const startTs = performance.now()
  const fileStream = fs.createReadStream(PATH);
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  try {
    let position = START_POSITION;
    let crossingsAtZero = 0;

    for await (const line of rl) {
      const [direction, distance] = parseLine(line.trim());

      if (direction === "L") {
        position -= distance;
      } else if (direction === "R") {
        position += distance;
      }

      position = wrapPosition(position);

      if (position === 0) {
        crossingsAtZero++;
      }
    }

    const endTs = performance.now()
    const elapsed = (endTs - startTs) * 1000;
    console.log(`Solution: ${crossingsAtZero} [${elapsed.toFixed(0)} (μs)]`);
  } finally {
    rl.close();
  }
}
