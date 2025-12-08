import sys
from .common import (
    FILE_PATH,
    Point,
    Segment,
    parse_line_into_point,
    parse_points_into_segments,
    sort_segments_by_length,
    island_sizes,
)


def connect_pairs(segments: list[Segment], num_of_connections: int) -> None:
    for i in range(num_of_connections):
        segment = segments[i]
        segment.start.connect(segment.end)


def sort_islands_by_size(island_sizes: list[int]) -> list[int]:
    return sorted(island_sizes, reverse=True)


def product_largest_islands(island_sizes: list[int], num: int) -> int:
    product = 1
    for i in range(num):
        size = island_sizes[i]
        product *= size
    return product


def day8_part1() -> None:
    try:
        running_id = 0
        points = []
        
        with open(FILE_PATH, "r") as f:
            for line in f:
                line = line.strip()
                point = parse_line_into_point(line, running_id)
                points.append(point)
                running_id += 1
        
        segments = sort_segments_by_length(parse_points_into_segments(points))
        connect_pairs(segments, 1000)
        islands = sort_islands_by_size(island_sizes(points))
        solution = product_largest_islands(islands, 3)
        
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
