from abc import ABC, abstractmethod


class FileRepository(ABC):
    @abstractmethod
    async def upload(self, data: bytes): ...
