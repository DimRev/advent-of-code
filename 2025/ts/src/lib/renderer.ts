import { execSync } from "child_process";
import * as path from "path";

const RENDERER_DIR = "../render";

export function PopulateRenderer(dayPart: string, value: number): void {
  const rendererPath = path.resolve(RENDERER_DIR);

  const cmd = `go run main.go populate ts ${dayPart} ${value}`;

  try {
    execSync(cmd, {
      cwd: rendererPath,
      stdio: "inherit",
    });
  } catch (err) {
    console.error(`Error running populate command: ${err}`);
    throw err;
  }
}