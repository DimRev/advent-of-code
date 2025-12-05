import sys
from typing import List, Tuple, Dict

FILE_PATH_P2 = "../inputs/day-4/part-1/input.txt"

def check_neighbors_p2(lines: List[str], i: int, j: int) -> int:
    min_i = 0
    max_i = len(lines) - 1
    min_j = 0
    neighbor_count = 0

    for y in range(i - 1, i + 2):
        if y < min_i or y > max_i:
            continue
        max_j_y = len(lines[y]) - 1
        for x in range(j - 1, j + 2):
            if x < min_j or x > max_j_y:
                continue
            if y == i and x == j:
                continue
            if lines[y][x] == "@":
                neighbor_count += 1

    return neighbor_count

def check_line_p2(lines: List[str]) -> Tuple[int, List[str]]:
    can_access_count = 0
    copy = lines.copy()

    to_remove = set()
    for i in range(len(lines)):
        for j in range(len(lines[i])):
            if lines[i][j] != "@":
                continue
            neighbors_count = check_neighbors_p2(lines, i, j)
            if neighbors_count < 4:
                to_remove.add((i, j))
                can_access_count += 1

    for i, j in to_remove:
        line_list = list(copy[i])
        line_list[j] = 'x'
        copy[i] = ''.join(line_list)

    return can_access_count, copy

def day4_part2() -> None:
    try:
        lines = []

        with open(FILE_PATH_P2, "r") as f:
            for line in f:
                line = line.strip()
                lines.append(line)

        # Filter out empty lines
        filtered_lines = [line for line in lines if line != ""]
        lines = filtered_lines

        total_removed = 0
        while True:
            valid_count, new_lines = check_line_p2(lines)
            if valid_count == 0:
                break
            total_removed += valid_count
            lines = new_lines

        print(f"Solution: {total_removed}")

    except FileNotFoundError:
        print("File not found")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)