import sys
from typing import List, Tuple

FILE_PATH = "../inputs/day-3/part-1/input.txt"

MAX_DIGITS = 12

def max_pair(line: str) -> int:
    max_idx = -1
    result = 0

    for digit_pos in range(MAX_DIGITS - 1, -1, -1):
        max_digit = -1
        for i in range(max_idx + 1, len(line) - digit_pos):
            num = int(line[i])
            if num > max_digit:
                max_digit = num
                max_idx = i

        result = result * 10 + max_digit

    return result

def day3_part2() -> None:
    try:
        solution = 0
        with open(FILE_PATH, "r") as f:
            for line in f:
                pair = max_pair(line.strip())
                solution += pair
        print(f"Solution: {solution}")

    except FileNotFoundError:
        print("File not found")
        sys.exit(1)
    except ValueError as e:
        print(f"Parsing error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)


