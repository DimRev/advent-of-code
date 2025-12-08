import sys
from .common import FILE_PATH

def generate_beam_list(line: str) -> list[bool]:
    result = []
    for char in line:
        result.append(char == 'S')
    return result

def progress_splitters(line: str, beam_list: list[bool]) -> tuple[list[bool], int]:
    new_beam_list = [False] * len(line)
    split_count = 0

    for i, char in enumerate(line):
        if char == '^' and beam_list[i]:
            if i - 1 >= 0:
                new_beam_list[i - 1] = True
            if i + 1 < len(line):
                new_beam_list[i + 1] = True
            new_beam_list[i] = False
            split_count += 1
            continue
        new_beam_list[i] = new_beam_list[i] or beam_list[i]

    return new_beam_list, split_count

def format_beam_list(beam_list: list[bool]) -> str:
    result = ""
    for b in beam_list:
        result += "|" if b else "."
    return result

def day7_part1() -> None:
    try:
        should_print = False
        split_count = 0
        beam_list = []
        idx = 0

        with open(FILE_PATH, "r") as f:
            for line in f:
                line = line.strip()

                if idx == 0:
                    beam_list = generate_beam_list(line)
                    if should_print:
                        print(line)
                        print(format_beam_list(beam_list))
                elif idx % 2 == 0:
                    new_beam_list, new_split_count = progress_splitters(line, beam_list)
                    split_count += new_split_count
                    beam_list = new_beam_list
                    if should_print:
                        print(line)
                        print(format_beam_list(beam_list))

                idx += 1

        solution = split_count
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