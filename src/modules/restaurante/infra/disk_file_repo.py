from pathlib import Path

import anyio

from src.constants import IMAGE_PATH

from ..domain.image_repository import FileRepository


class DiskFileRepository(FileRepository):
    async def upload(self, data: bytes, relative_path: Path):
        print("[DiskFileRepository] Uploading file to disk...")

        path = IMAGE_PATH / relative_path

        path.parent.mkdir(parents=True, exist_ok=True)

        async with await anyio.open_file(path, "wb") as f:
            await f.write(data)
