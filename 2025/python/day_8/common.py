import math
from typing import Dict

FILE_PATH = "../inputs/day-8/part-1/input.txt"


class Point:
    def __init__(self, x: int, y: int, z: int, id: int):
        self.id = id
        self.x = x
        self.y = y
        self.z = z
        self.connections_map: Dict[int, 'Point'] = {}

    def __str__(self) -> str:
        connections = self.list_connections_ids()
        return f"Point[{self.id}]({self.x}, {self.y}, {self.z}) Connections: {connections}"

    def distance(self, target: 'Point') -> int:
        diff_x_squared = math.pow(target.x - self.x, 2)
        diff_y_squared = math.pow(target.y - self.y, 2)
        diff_z_squared = math.pow(target.z - self.z, 2)
        return int(abs(math.sqrt(diff_x_squared + diff_y_squared + diff_z_squared)))

    def connect(self, target: 'Point') -> None:
        if target.id not in self.connections_map:
            self.connections_map[target.id] = target

        if self.id not in target.connections_map:
            target.connections_map[self.id] = self

    def list_connections_ids(self) -> list[int]:
        return list(self.connections_map.keys())


class Segment:
    def __init__(self, start: Point, end: Point, length: int):
        self.start = start
        self.end = end
        self.length = length

    def __str__(self) -> str:
        return f"Segment(Point[{self.start.id}], Point[{self.end.id}], {self.length})"


def generate_point(x: int, y: int, z: int, id: int) -> Point:
    return Point(x, y, z, id)


def parse_line_into_point(line: str, id: int) -> Point:
    parts = line.split(",")
    if len(parts) != 3:
        raise ValueError(f"Error parsing line: {line}")

    x = int(parts[0])
    y = int(parts[1])
    z = int(parts[2])

    return generate_point(x, y, z, id)


def parse_points_into_segments(points: list[Point]) -> list[Segment]:
    segments = []
    for i in range(len(points) - 1):
        for j in range(i + 1, len(points)):
            segment = Segment(
                start=points[i],
                end=points[j],
                length=points[j].distance(points[i])
            )
            segments.append(segment)

    return segments


def sort_segments_by_length(segments: list[Segment]) -> list[Segment]:
    return sorted(segments, key=lambda s: s.length)


def traverse_connections(point: Point, prev_count: int, visited_ids: dict[int, bool]) -> int:
    if point.id in visited_ids:
        return 0

    if len(point.connections_map) == 0:
        return 1

    count = 1
    visited_ids[point.id] = True

    for p in point.connections_map.values():
        count += traverse_connections(p, prev_count + 1, visited_ids)

    return count


def island_sizes(points: list[Point]) -> list[int]:
    visited_ids = {}
    sizes = []
    for p in points:
        size = traverse_connections(p, 0, visited_ids)
        if size > 0:
            sizes.append(size)
    return sizes


def island_sizes_p2(points: list[Point]) -> list[int]:
    visited_ids = {}
    sizes = []
    for p in points:
        size = traverse_connections(p, 0, visited_ids)
        if size > 0:
            sizes.append(size)
        if len(sizes) > 1:
            return []
    return sizes
