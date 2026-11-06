from pathlib import Path

import anyio

from src.constants import IMAGE_PATH

from ..domain.image_repository import FileRepository


class DiskFileRepository(FileRepository):
    async def upload(self, data: bytes, relative_path: Path):
        print("[DiskFileRepository] Uploading file to disk...")

        path = IMAGE_PATH / relative_path
        path_folder = path.parent

        path_folder.mkdir(parents=True, exist_ok=True)

        for image in path_folder.iterdir():
            if image.is_file() and image.stem == path.stem:
                image.unlink()

        async with await anyio.open_file(path, "wb") as f:
            await f.write(data)
