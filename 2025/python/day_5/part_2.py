import sys
from typing import TypedDict

FILE_PATH_P2 = "../inputs/day-5/part-1/input.txt"

class Range(TypedDict):
    start: int
    end: int

def parse_range_p2(range_str: str) -> Range:
    parts = range_str.split("-")
    if len(parts) != 2:
        raise ValueError(f"Invalid range format: {range_str}")
    start = int(parts[0])
    end = int(parts[1])
    return Range(start=start, end=end)

def squash_ranges_p2(ranges: list[Range]) -> list[Range]:
    if len(ranges) == 0:
        return ranges

    sorted_ranges = sorted(ranges, key=lambda r: r["start"])

    result: list[Range] = [sorted_ranges[0]]

    for r in sorted_ranges[1:]:
        last = result[-1]
        if r["start"] <= last["end"] + 1:
            last["end"] = max(last["end"], r["end"])
        else:
            result.append(r)

    return result

def check_total_ids_in_ranges_p2(ranges: list[Range]) -> int:
    total = 0
    for r in ranges:
        total += r["end"] - r["start"] + 1
    return total

def day5_part2() -> None:
    try:
        is_range_section = True
        ranges: list[Range] = []

        with open(FILE_PATH_P2, "r") as f:
            for line in f:
                line = line.strip()
                if line == "":
                    is_range_section = False
                    break
                if is_range_section:
                    try:
                        r = parse_range_p2(line)
                        ranges.append(r)
                    except ValueError as e:
                        print(f"Error parsing range: {e}")
                        sys.exit(1)

        squashed = squash_ranges_p2(ranges)
        total_count = check_total_ids_in_ranges_p2(squashed)

        print(f"Solution: {total_count}")

    except FileNotFoundError:
        print("File not found")
        sys.exit(1)
    except ValueError as e:
        print(f"Parsing error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)