import sys
from typing import List

FILE_PATH = "../inputs/day-4/part-1/input.txt"

def check_neighbors(lines: List[str], i: int, j: int) -> int:
    min_i = 0
    max_i = len(lines) - 1
    min_j = 0
    max_j = len(lines[0]) - 1
    neighbor_count = 0

    for y in range(i-1, i+2):
        if y < min_i or y > max_i:
            continue
        for x in range(j-1, j+2):
            if x < min_j or x > max_j:
                continue
            if y == i and x == j:
                continue
            if lines[y][x] == "@":
                neighbor_count += 1

    return neighbor_count

def check_line(lines: List[str], check_idx: int) -> int:
    can_access_count = 0
    for i in range(len(lines)):
        if i != check_idx:
            continue
        for j, char in enumerate(lines[i]):
            if char != "@":
                continue
            neighbors_count = check_neighbors(lines, i, j)
            if neighbors_count < 4:
                can_access_count += 1
    return can_access_count

def day4_part1() -> None:
    try:
        solution = 0
        lines = []

        with open(FILE_PATH, "r") as f:
            for line in f:
                line = line.strip()
                lines.append(line)

                if len(lines) == 2:
                    solution += check_line(lines, 0)

                if len(lines) == 3:
                    solution += check_line(lines, 1)
                    lines = lines[1:]

        solution += check_line(lines, 1)

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
