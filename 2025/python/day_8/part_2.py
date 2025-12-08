import sys
from .common import (
    FILE_PATH,
    Point,
    Segment,
    parse_line_into_point,
    parse_points_into_segments,
    sort_segments_by_length,
    island_sizes_p2,
)


def connect_all_pairs(segments: list[Segment], points: list[Point]) -> int:
    for i in range(len(segments)):
        segment = segments[i]
        segment.start.connect(segment.end)

        islands = island_sizes_p2(points)

        if len(islands) == 1:
            return segment.start.x * segment.end.x

    return -1


def day8_part2() -> None:
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
        solution = connect_all_pairs(segments, points)

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
