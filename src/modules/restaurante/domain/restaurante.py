from uuid import UUID

from pydantic import BaseModel


class Restaurante(BaseModel):
    id: UUID
    nome: str
