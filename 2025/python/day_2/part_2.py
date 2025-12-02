import sys
from typing import List, Tuple

FILE_PATH = "../inputs/day-2/part-1/input.txt"
COMMA_SEPARATOR = ','
RANGE_SEPARATOR = '-'
LEADING_ZERO_CHAR = '0'
EXPECTED_RANGE_PARTS = 2
MIN_PATTERN_LEN = 1

def is_repeating_pattern(num_str: str, pattern_len: int) -> bool:
    """Check if the given string is made entirely of a pattern of the specified length."""
    pattern = num_str[0:pattern_len]

    for i in range(0, len(num_str), pattern_len):
        if num_str[i:i + pattern_len] != pattern:
            return False
    return True

def check_if_invalid_p2(num: int) -> bool:
    """
    Check if a number is invalid in Part 2.
    Invalid IDs are made of any pattern repeated throughout the entire number.
    """
    num_str = str(num)
    length = len(num_str)

    if num_str[0] == LEADING_ZERO_CHAR:
        return False


    max_pattern_len = length // 2
    for pattern_len in range(MIN_PATTERN_LEN, max_pattern_len + 1):
        if length % pattern_len != 0:
            continue

        if is_repeating_pattern(num_str, pattern_len):
            return True

    return False

def find_invalid_nums_in_range_p2(min_num: int, max_num: int) -> List[int]:
    """Find all invalid numbers in a range."""
    invalid_nums = []

    for num in range(min_num, max_num + 1):
        if check_if_invalid_p2(num):
            invalid_nums.append(num)

    return invalid_nums

def parse_line_p2(line: str) -> Tuple[int, int]:
    """Parse a line into min and max numbers."""
    nums = line.split(RANGE_SEPARATOR)
    if len(nums) != EXPECTED_RANGE_PARTS:
        raise ValueError(f"Invalid line format: {line}")

    min_num = int(nums[0])
    max_num = int(nums[1])

    return min_num, max_num

def day2_part2() -> None:
    try:
        solution = 0
        with open(FILE_PATH, "r") as f:
            for line in f:
                ranges = [r.strip() for r in line.strip().split(COMMA_SEPARATOR) if r.strip()]
                for range_str in ranges:
                    min_num, max_num = parse_line_p2(range_str)
                    invalid_list = find_invalid_nums_in_range_p2(min_num, max_num)
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