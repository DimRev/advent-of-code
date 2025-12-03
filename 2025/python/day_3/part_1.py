import sys
from typing import List, Tuple

FILE_PATH = "../inputs/day-3/part-1/input.txt"

def max_pair(line: str) -> int:
    first_max_digit = -1
    first_max_digit_index = -1

    for i, v in enumerate(line):
        if i == len(line) - 1:
            break
        if v.isdigit():
            digit = int(v)
            if digit > first_max_digit:
                first_max_digit = digit
                first_max_digit_index = i

    second_max_digit = -1
    for _, u in enumerate(line[first_max_digit_index+1:]):
        if u.isdigit():
            digit = int(u)
            if digit > second_max_digit:
                second_max_digit = digit

    return (first_max_digit * 10) + second_max_digit

def day3_part1() -> None:
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


