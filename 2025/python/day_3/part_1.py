import sys
from typing import List, Tuple

FILE_PATH = "../inputs/day-3/part-1/example.txt"

def max_pair(line: str) -> int:
    max_first_digit = -1
    max_first_digit_idx = -1

    for [v, i] in enumerate(line):
        if i == len(line) - 1:
            break
        num = int(v)
        if num > max_first_digit:
            max_first_digit = num
            max_first_digit_idx = i

    max_second_digit = -1
    for [v, i] in enumerate(line[max_first_digit_idx+1:]):
        num = int(v)
        if num > max_first_digit:
            max_second_digit = num

    return (max_first_digit * 10) + max_second_digit

def day3_part1() -> None:
    try:
        solution = 0
        with open(FILE_PATH, "r") as f:
            for line in f:
                pair = max_pair(line)
                print(f"%d - %s", line, pair)
                solution += line
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