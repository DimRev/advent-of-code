import { day1Part1, day1Part2 } from "./day-1";
import { day2Part1, day2Part2 } from "./day-2";

const commands = ["d1p1", "d1p2", "d2p1", "d2p2"] as const;
type Cmds = typeof commands[number];
type CmdMap = { [key in Cmds]: () => Promise<void> };

const cmdMap: CmdMap = {
  d1p1: day1Part1,
  d1p2: day1Part2,
  d2p1: day2Part1,
  d2p2: day2Part2,
};

async function main(): Promise<void> {
  const args = process.argv.slice(2);

  if (args.length === 0) {
    console.log("Usage: pnpm start <command>");
    console.log(`Available commands:
${commands.map(c=>`\t- ${c}`).join("\n")}`);
    process.exit(1);
  }

  const unparsedCmd = args[0];

  if (!(unparsedCmd in cmdMap)) {
    console.log(`Unknown command: ${unparsedCmd}`);
    console.log(`Available commands:
${commands.map(c=>`\t- ${c}`).join("\n")}`);
    process.exit(1);
  }

  const parsedCmd = unparsedCmd as Cmds;
  const startTs = performance.now()
  await cmdMap[parsedCmd]();
  const endTs = performance.now()
  const elapsed = (endTs - startTs) * 1000;
  console.log(`Finished running ${unparsedCmd} in ${elapsed.toFixed(0)} (μs)`);
}

void main();
