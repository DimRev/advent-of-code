import sys
from typing import Tuple
FILE_PATH = '../inputs/day-1/part-1/input.txt'

def parseLine(line: str)-> Tuple[str, int]:
    if len(line) < 2:
        return ""

    return line[0], int(line[1:])


def calculatePassesOverZero(dir: str, start_pos: int, distance: int) -> int:
    if distance == 0:
        return 0

    if dir == 'R':
        distance_to_zero = 100 - start_pos
        if start_pos == 0:
            distance_to_zero = 100
        if distance >= distance_to_zero:
            return 1 + (distance - distance_to_zero) // 100
        return 0
    else:
        distance_to_zero = start_pos
        if start_pos == 0:
            distance_to_zero = 100
        if distance >= distance_to_zero:
            return 1 + (distance - distance_to_zero) // 100
        return 0

def day1_part2():
    try:
        with open(FILE_PATH, 'r') as f:
            n = 50
            z = 0
            for line in f:

                direction, dist = parseLine(line.strip())

                full_rotations_from_dist = dist // 100
                z += full_rotations_from_dist
                dist = dist % 100

                passes_over_zero = calculatePassesOverZero(direction, n, dist)
                z += passes_over_zero

                if direction == 'L':
                    n -= dist
                else:
                    n += dist

                n = ((n % 100) + 100) % 100

            print(f'Solution: {z}')
    except FileNotFoundError:
        print('File not found')
        sys.exit(1)
    except Exception as e:
        print(e)
        sys.exit(1)