from pathlib import Path
from uuid import UUID, uuid4

from fastapi import UploadFile

from src.core.errors import DomainError

from .domain.image_repository import FileRepository
from .domain.menu_repository import MenuRepository
from .domain.restaurante import Restaurante
from .domain.restaurante_repository import RestauranteRepository
from .types import (
    CreateRestauranteRequest,
    UpdateRestauranteRequest,
)


class RestauranteService:
    def __init__(
        self,
        repo: RestauranteRepository,
        menu_repo: MenuRepository,
        file_repo: FileRepository,
    ):
        self.repo = repo
        self.menu_repo = menu_repo
        self.file_repo = file_repo

    async def list_restaurantes(self) -> list[Restaurante]:
        print("[RestauranteService] Listando restaurantes...")

        return await self.repo.list_()

    async def get_restaurante(self, restaurante_id: UUID) -> Restaurante:
        print(f"[RestauranteService] Obtendo restaurante com ID: {restaurante_id}")

        restaurant = await self.repo.get_by_id(restaurante_id)

        if not restaurant:
            raise DomainError(
                "Restaurante não encontrado",
                404,
            )

        return restaurant

    async def create_restaurante(
        self,
        data: CreateRestauranteRequest,
    ) -> Restaurante:
        print("[RestauranteService] Criando novo restaurante...")

        restaurante = Restaurante(
            id=uuid4(),
            nome=data.nome,
        )

        await self.repo.save(restaurante)
        await self.menu_repo.save(restaurante.id)

        return restaurante

    async def update_restaurante(
        self,
        data: UpdateRestauranteRequest,
    ) -> Restaurante:
        print(f"[RestauranteService] Atualizando restaurante com ID: {data.id}")

        exists = bool(await self.repo.get_by_id(data.id))

        if not exists:
            raise DomainError("Restaurante não encontrado", 404)

        restaurante = Restaurante(
            id=data.id,
            nome=data.nome,
        )

        await self.repo.update(restaurante)

        return restaurante

    async def delete_restaurante(self, restaurante_id: UUID):
        print(f"[RestauranteService] Excluindo restaurante com ID: {restaurante_id}")

        exists = bool(await self.repo.get_by_id(restaurante_id))

        if not exists:
            raise DomainError("Restaurante não encontrado", 404)

        await self.repo.delete(restaurante_id)

    async def upload_restaurante_image(
        self,
        restaurante_id: UUID,
        file: UploadFile,
    ):
        print(
            f"[RestauranteService] Uploading image for restaurante with ID: {restaurante_id}"
        )

        if file.size is None or file.size == 0:
            raise DomainError("Arquivo de imagem vazio", 400)

        if file.size > 5 * 1024 * 1024:  # 5MB
            raise DomainError(
                "Arquivo de imagem muito grande (máximo 5MB)",
                400,
            )

        if not file.filename:
            raise DomainError("Arquivo de imagem não fornecido", 400)

        extension = file.filename.split(".")[-1].lower()

        if extension not in ["jpg", "jpeg", "png", "gif", "webp"]:
            raise DomainError("Formato de imagem não suportado", 400)

        path = Path("logos") / f"{restaurante_id}.{extension}"

        await self.file_repo.upload(
            data=await file.read(),
            relative_path=path,
        )
