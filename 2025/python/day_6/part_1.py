import sys

FILE_PATH = "../inputs/day-6/part-1/input.txt"


def map_transform_nums(slice: list[str]) -> list[int]:
    result = []
    for v in slice:
        if v == "":
            continue
        try:
            num = int(v)
            result.append(num)
        except ValueError:
            continue
    return result


def map_clean_symbols(slice: list[str]) -> list[str]:
    result = []
    for v in slice:
        if v == "":
            continue
        result.append(v)
    return result


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


def calculate_total(nums_list: list[list[int]], symbols: list[str], line_len: int) -> int:
    total = 0
    for i in range(line_len):
        curr_nums = []
        for nums in nums_list:
            curr_nums.append(nums[i])

        if symbols[i] == "+":
            total += addition(curr_nums)
        elif symbols[i] == "*":
            total += multiplication(curr_nums)
        else:
            print(f"Unknown symbol: {symbols[i]}")
            sys.exit(1)

    return total


def day6_part1() -> None:
    try:
        with open(FILE_PATH, "r") as f:
            nums_list = []
            symbols = []
            line_len = 0

            for line in f:
                line = line.strip()

                parsed = line.split(" ")
                transformed = map_transform_nums(parsed)
                if len(transformed) != 0:
                    nums_list.append(transformed)
                else:
                    cleaned_symbols = map_clean_symbols(parsed)
                    symbols = cleaned_symbols
                    line_len = len(cleaned_symbols)

        solution = calculate_total(nums_list, symbols, line_len)

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