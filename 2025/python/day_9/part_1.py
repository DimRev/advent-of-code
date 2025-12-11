import sys
from .common import FILE_PATH

def day9_part1() -> None:
    try:

        with open(FILE_PATH, "r") as f:
            for line in f:
                line = line.strip()

        solution = 0

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