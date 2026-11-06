from ..domain.image_repository import FileRepository


class DiskFileRepository(FileRepository):
    async def upload(self, data: bytes):
        print("[DiskFileRepository] Uploading file to disk...")
