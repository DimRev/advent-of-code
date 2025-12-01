import sys
import time
from typing import Tuple

FILE_PATH = "../inputs/day-1/part-1/input.txt"
CIRCLE_SIZE = 100
START_POSITION = 50


def parse_line(line: str) -> Tuple[str, int]:
    """Parse a line into direction and distance."""
    if len(line) < 2:
        raise ValueError(f"Invalid line format: {line}")

    return line[0], int(line[1:])


def wrap_position(position: int) -> int:
    """Wrap position around the circle."""
    return ((position % CIRCLE_SIZE) + CIRCLE_SIZE) % CIRCLE_SIZE


def day1_part1() -> None:
    """Calculate crossings at position 0 for part 1."""
    try:
        start_ts = time.perf_counter_ns()
        position = START_POSITION
        crossings_at_zero = 0

        with open(FILE_PATH, "r") as f:
            for line in f:
                direction, distance = parse_line(line.strip())

                if direction == "L":
                    position -= distance
                elif direction == "R":
                    position += distance

                position = wrap_position(position)

                if position == 0:
                    crossings_at_zero += 1

        end_ts = time.perf_counter_ns()
        elapsed = (end_ts - start_ts) / 1000
        print(f"Solution: {crossings_at_zero}  [{elapsed} (μs)]")

    except FileNotFoundError:
        print("File not found")
        sys.exit(1)
    except ValueError as e:
        print(f"Parsing error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)

