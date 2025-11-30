import { part1, part2 } from "./day-1";

const commands = ["d1p1", "d1p2"] as const;
type Cmds = typeof commands[number];
type CmdMap = { [key in Cmds]: () => void };

const cmdMap: CmdMap = {
  d1p1: part1,
  d1p2: part2,
};

function main(): void {
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
  cmdMap[parsedCmd]();
}

main();
