import sys

FILE_PATH = "../inputs/day-6/part-1/input.txt"


def addition(nums: list[int]) -> int:
    total = 0
    for num in nums:
        total += num
    return total


def multiplication(nums: list[int]) -> int:
    total = 1
    for num in nums:
        total *= num
    return total


def read_lines(file_path: str) -> list[str]:
    with open(file_path, "r") as f:
        lines = []
        for line in f:
            lines.append(line.rstrip('\n'))
    return lines


def pad_lines(lines: list[str]) -> list[str]:
    max_len = 0
    for line in lines:
        if len(line) > max_len:
            max_len = len(line)

    padded = []
    for line in lines:
        if len(line) < max_len:
            padded.append(line + " " * (max_len - len(line)))
        else:
            padded.append(line)
    return padded


def is_column_empty(lines: list[str], col: int) -> bool:
    for row in range(len(lines)):
        if col < len(lines[row]) and lines[row][col] != ' ':
            return False
    return True


def read_number_vertically(lines: list[str], col: int, start_row: int) -> tuple[int, int, bool]:
    num_str = ""
    row = start_row
    while row < len(lines):
        if col >= len(lines[row]):
            break
        ch = lines[row][col]
        if ch.isdigit():
            num_str += ch
            row += 1
        else:
            break

    if num_str == "":
        return 0, start_row, False

    try:
        num = int(num_str)
        return num, row - 1, True
    except ValueError:
        return 0, start_row, False


def process_problem(lines: list[str], start_col: int) -> tuple[list[int], str, int]:
    numbers = []
    operator = ""
    col = start_col

    while col >= 0:
        if is_column_empty(lines, col):
            break

        row = 0
        while row < len(lines):
            if col >= len(lines[row]):
                row += 1
                continue

            char = lines[row][col]
            if char == ' ':
                row += 1
                continue

            if char == '+' or char == '*':
                operator = char
                row += 1
            elif char.isdigit():
                num, last_row, success = read_number_vertically(lines, col, row)
                if success:
                    numbers.append(num)
                    row = last_row + 1
                else:
                    row += 1
            else:
                row += 1

        col -= 1

    return numbers, operator, col


def calculate_problem(numbers: list[int], operator: str) -> int:
    if len(numbers) == 0 or operator == "":
        return 0

    if operator == "+":
        return addition(numbers)
    elif operator == "*":
        return multiplication(numbers)
    else:
        return 0


def day6_part2() -> None:
    try:
        lines = read_lines(FILE_PATH)

        if len(lines) == 0:
            print("No lines found")
            return

        lines = pad_lines(lines)
        grand_total = 0
        col = len(lines[0]) - 1

        while col >= 0:
            while col >= 0 and is_column_empty(lines, col):
                col -= 1

            if col < 0:
                break

            numbers, operator, next_col = process_problem(lines, col)
            result = calculate_problem(numbers, operator)
            grand_total += result
            col = next_col

        print(f"Solution: {grand_total}")

    except FileNotFoundError:
        print("File not found")
        sys.exit(1)
    except ValueError as e:
        print(f"Parsing error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)