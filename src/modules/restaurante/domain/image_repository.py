from abc import ABC, abstractmethod
from pathlib import Path


class FileRepository(ABC):
    @abstractmethod
    async def upload(self, data: bytes, relative_path: Path): ...
