import sys
from typing import TypedDict

FILE_PATH = "../inputs/day-5/part-1/input.txt"

class Range(TypedDict):
    start: int
    end: int

def parse_range_p1(range_str: str) -> Range:
    parts = range_str.split("-")
    if len(parts) != 2:
        raise ValueError(f"Invalid range format: {range_str}")
    start = int(parts[0])
    end = int(parts[1])
    return Range(start=start, end=end)

def is_id_in_ranges(ranges: list[Range], id: int) -> bool:
    for r in ranges:
        if id >= r["start"] and id <= r["end"]:
            return True
    return False

def day5_part1() -> None:
    try:
        solution = 0
        is_range_section = True
        ranges: list[Range] = []
        ids: list[int] = []

        with open(FILE_PATH, "r") as f:
            for line in f:
                line = line.strip()
                if line == "":
                    is_range_section = False
                    continue
                if is_range_section:
                    try:
                        r = parse_range_p1(line)
                        ranges.append(r)
                    except ValueError as e:
                        print(f"Error parsing range: {e}")
                        sys.exit(1)
                else:
                    try:
                        id = int(line)
                        ids.append(id)
                    except ValueError as e:
                        print(f"Error parsing ID: {e}")
                        sys.exit(1)

        sorted_ranges = sorted(ranges, key=lambda r: r["start"])

        for id in ids:
            if is_id_in_ranges(sorted_ranges, id):
                solution += 1

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