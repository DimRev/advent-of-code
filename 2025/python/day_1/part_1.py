import sys
from typing import Tuple
FILE_PATH = '../inputs/day-1/part-1/example.txt'

def parseLine(line: str)-> Tuple[str, int]:
    if len(line) < 2:
        return ""

    return line[0], int(line[1:])

def day1_part1():
    try:
        with open(FILE_PATH, 'r') as f:
            n = 50
            for line in f:

                direction, num = parseLine(line.strip())
    except FileNotFoundError:
        print('File not found')
        sys.exit(1)
    except Exception as e:
        print(e)
        sys.exit(1)
