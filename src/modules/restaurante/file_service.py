from pathlib import Path

from fastapi import UploadFile

from src.constants import IMAGE_PATH
from src.core.errors import DomainError

from .domain.image_repository import FileRepository

MAX_IMAGE_SIZE_MB = 5
MAX_IMAGE_SIZE = MAX_IMAGE_SIZE_MB * 1024 * 1024  # 5MB
ACCEPTED_IMAGE_FORMATS = ["jpg", "jpeg", "png", "gif", "webp"]


class FileService:
    def __init__(self, repo: FileRepository):
        self.repo = repo

    async def upload_image(
        self,
        file: UploadFile,
        base_path: Path,
        *,
        filename_override: str | None = None,
    ) -> None:
        print(f"[FileService] Uploading image to {base_path}...")

        if not file.filename:
            raise DomainError("Arquivo de imagem não fornecido", 400)

        extension = file.filename.split(".")[-1].lower()

        if extension not in ACCEPTED_IMAGE_FORMATS:
            raise DomainError("Formato de imagem não suportado", 400)

        if file.size is None or file.size == 0:
            raise DomainError("Arquivo de imagem vazio", 400)

        if file.size > MAX_IMAGE_SIZE:
            raise DomainError(
                f"Arquivo de imagem muito grande (máximo {MAX_IMAGE_SIZE_MB}MB)", 400
            )

        name = f"{filename_override or file.filename.split('.')[0]}.{extension}"

        final_path = base_path / name

        await self.repo.upload(
            data=await file.read(),
            relative_path=final_path,
        )

        print(f"[FileService] Image uploaded to {final_path}")

    async def get_image(self, base_path: Path, filename: str) -> Path:
        print(f"[FileService] Getting image from {base_path}...")

        for ext in ACCEPTED_IMAGE_FORMATS:
            file_path = IMAGE_PATH / base_path / f"{filename}.{ext}"

            if file_path.exists():
                return file_path

        raise DomainError("Imagem não encontrada", 404)
