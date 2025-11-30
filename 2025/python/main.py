import sys
from enum import Enum
from typing import Callable, Dict

from day_1 import day1_part1, day1_part2


class Command(str, Enum):
    D1P1 = "d1p1"
    D1P2 = "d1p2"


CMD_MAP: Dict[Command, Callable[[], None]] = {
    Command.D1P1: day1_part1,
    Command.D1P2: day1_part2,
}


def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <command>")
        print("Available commands:")
        for cmd in Command:
            print(f"\t- {cmd.value}")
        sys.exit(1)

    cmd_str = sys.argv[1]

    try:
        cmd = Command(cmd_str)
    except ValueError:
        print(f"Unknown command: {cmd_str}")
        print("Available commands:")
        for cmd_enum in Command:
            print(f"\t- {cmd_enum.value}")
        sys.exit(1)

    CMD_MAP[cmd]()


if __name__ == "__main__":
    main()