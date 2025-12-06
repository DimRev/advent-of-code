import sys
import time
from enum import Enum
from typing import Callable, Dict

from day_1 import day1_part1, day1_part2
from day_2 import day2_part1, day2_part2
from day_3 import day3_part1, day3_part2
from day_4 import day4_part1, day4_part2
from day_5 import day5_part1, day5_part2
from day_6 import day6_part1, day6_part2
from utils import populate_renderer

class Command(str, Enum):
    D1P1 = "d1p1"
    D1P2 = "d1p2"
    D2P1 = "d2p1"
    D2P2 = "d2p2"
    D3P1 = "d3p1"
    D3P2 = "d3p2"
    D4P1 = "d4p1"
    D4P2 = "d4p2"
    D5P1 = "d5p1"
    D5P2 = "d5p2"
    D6P1 = "d6p1"
    D6P2 = "d6p2"


CMD_MAP: Dict[Command, Callable[[], None]] = {
    Command.D1P1: day1_part1,
    Command.D1P2: day1_part2,
    Command.D2P1: day2_part1,
    Command.D2P2: day2_part2,
    Command.D3P1: day3_part1,
    Command.D3P2: day3_part2,
    Command.D4P1: day4_part1,
    Command.D4P2: day4_part2,
    Command.D5P1: day5_part1,
    Command.D5P2: day5_part2,
    Command.D6P1: day6_part1,
    Command.D6P2: day6_part2,
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

    start_ts = time.perf_counter_ns()
    CMD_MAP[cmd]()
    end_ts = time.perf_counter_ns()
    elapsed = (end_ts - start_ts) / 1000
    print(f"Finished running {cmd_str} in {elapsed} (μs)")
    populate_renderer(cmd_str, elapsed)

if __name__ == "__main__":
    main()