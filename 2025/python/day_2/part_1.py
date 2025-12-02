import sys
from typing import List, Tuple

# Constants
FILE_PATH = "../inputs/day-2/part-1/input.txt"
COMMA_SEPARATOR = ','
RANGE_SEPARATOR = '-'
LEADING_ZERO_CHAR = '0'
EXPECTED_RANGE_PARTS = 2

def check_if_invalid_p1(num: int) -> bool:
    """
    Check if a number is invalid in Part 1.
    Invalid IDs are made of exactly two identical halves (e.g., 11, 1010, 123123).
    """
    num_str = str(num)
    length = len(num_str)

    # Invalid IDs must have even length
    if length % 2 != 0:
        return False

    # No leading zeros allowed
    if num_str[0] == LEADING_ZERO_CHAR:
        return False

    half = length // 2
    first_half = num_str[0:half]
    second_half = num_str[half:]

    return first_half == second_half

def find_invalid_nums_in_range_p1(min_num: int, max_num: int) -> List[int]:
    """Find invalid numbers in a range."""
    invalid_nums = []

    for num in range(min_num, max_num + 1):
        if check_if_invalid_p1(num):
            invalid_nums.append(num)

    return invalid_nums

def parse_line_p1(line: str) -> Tuple[int, int]:
    """Parse a line into min and max numbers."""
    nums = line.split(RANGE_SEPARATOR)
    if len(nums) != EXPECTED_RANGE_PARTS:
        raise ValueError(f"Invalid line format: {line}")

    min_num = int(nums[0])
    max_num = int(nums[1])

    return min_num, max_num

def day2_part1() -> None:
    try:
        solution = 0
        with open(FILE_PATH, "r") as f:
            for line in f:
                # Split by comma separator
                ranges = [r.strip() for r in line.strip().split(COMMA_SEPARATOR) if r.strip()]
                for range_str in ranges:
                    min_num, max_num = parse_line_p1(range_str)
                    invalid_list = find_invalid_nums_in_range_p1(min_num, max_num)
                    solution += sum(invalid_list)
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