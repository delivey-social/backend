from uuid import UUID

from pydantic import BaseModel

from .domain.restaurante import Restaurante


class ListRestaurantesResponse(BaseModel):
    restaurantes: list[Restaurante]


class GetRestauranteResponse(Restaurante):
    pass


class CreateRestauranteRequest(BaseModel):
    nome: str


class UpdateRestauranteRequest(BaseModel):
    id: UUID
    nome: str


class CreateMenuItemRequest(BaseModel):
    nome: str
    preco: int


class UpdateMenuItemRequest(BaseModel):
    nome: str
    preco: int
