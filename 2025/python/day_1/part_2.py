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


def calculate_passes_over_zero(direction: str, start_pos: int, distance: int) -> int:
    """Calculate number of times we pass through position 0 during a move."""
    if distance == 0:
        return 0

    if direction == "R":
        distance_to_zero = CIRCLE_SIZE if start_pos == 0 else CIRCLE_SIZE - start_pos
    else:
        distance_to_zero = CIRCLE_SIZE if start_pos == 0 else start_pos

    if distance >= distance_to_zero:
        return 1 + (distance - distance_to_zero) // CIRCLE_SIZE
    return 0


def day1_part2() -> None:
    """Calculate total crossings at position 0 for part 2."""
    try:
        position = START_POSITION
        crossings_at_zero = 0

        with open(FILE_PATH, "r") as f:
            for line in f:
                direction, distance = parse_line(line.strip())

                full_rotations = distance // CIRCLE_SIZE
                crossings_at_zero += full_rotations

                remaining_distance = distance % CIRCLE_SIZE
                additional_crossings = calculate_passes_over_zero(
                    direction, position, remaining_distance
                )
                crossings_at_zero += additional_crossings

                if direction == "L":
                    position -= remaining_distance
                elif direction == "R":
                    position += remaining_distance

                position = wrap_position(position)

        print(f"Solution: {crossings_at_zero}")

    except FileNotFoundError:
        print("File not found")
        sys.exit(1)
    except ValueError as e:
        print(f"Parsing error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)
