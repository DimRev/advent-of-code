import sys
from .common import FILE_PATH

def generate_path_counter(line: str) -> list[int]:
    result = []
    for char in line:
        if char == 'S':
            result.append(1)
        else:
            result.append(0)
    return result

def progress_splitters_dp(line: str, path_counter: list[int]) -> list[int]:
    new_path_counter = [0] * len(line)

    for i, char in enumerate(line):
        if path_counter[i] == 0:
            continue

        if char == '^':
            if i - 1 >= 0:
                new_path_counter[i - 1] += path_counter[i]
            if i + 1 < len(line):
                new_path_counter[i + 1] += path_counter[i]
        else:
            new_path_counter[i] += path_counter[i]

    return new_path_counter

def day7_part2() -> None:
    try:
        path_counter = []
        idx = 0

        with open(FILE_PATH, "r") as f:
            for line in f:
                line = line.strip()

                if idx == 0:
                    path_counter = generate_path_counter(line)
                elif idx % 2 == 0:
                    path_counter = progress_splitters_dp(line, path_counter)

                idx += 1

        solution = sum(path_counter)
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