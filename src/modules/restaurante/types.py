from uuid import UUID

from pydantic import BaseModel


class CreateRestauranteRequest(BaseModel):
    nome: str


class UpdateRestauranteRequest(BaseModel):
    id: UUID
    nome: str


class CreateMenuItemRequest(BaseModel):
    nome: str
    preco: int
