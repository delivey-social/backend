from uuid import UUID

from pydantic import BaseModel


class MenuItem(BaseModel):
    id: UUID
    nome: str
    preco: int


class Menu(BaseModel):
    id: UUID
    restaurante_id: UUID
    itens: list[MenuItem]


class Restaurante(BaseModel):
    id: UUID
    nome: str
