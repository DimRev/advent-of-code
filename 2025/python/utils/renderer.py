import subprocess
import os
from pathlib import Path

RENDERER_DIR = "../render"


def populate_renderer(day_part: str, value: int) -> None:
    """
    Populate the renderer data.json with benchmark results.

    Args:
        day_part: day and part identifier (e.g., d1p1, d1p2)
        value: microseconds value
    """
    renderer_path = os.path.abspath(RENDERER_DIR)

    cmd = ["go", "run", "main.go", "populate", "python", day_part, str(value)]

    try:
        result = subprocess.run(cmd, cwd=renderer_path, check=True, capture_output=False)
    except subprocess.CalledProcessError as e:
        print(f"Error running populate command: {e}")
        raise
