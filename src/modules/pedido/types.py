from typing import Annotated
from uuid import UUID

from pydantic import BaseModel, Field


class CreatePedidoRequestItem(BaseModel):
    product_id: UUID
    quantity: int = Field(..., gt=0)


class CreatePedidoRequest(BaseModel):
    user_id: UUID
    restaurant_id: UUID
    items: Annotated[list[CreatePedidoRequestItem], Field(..., min_length=1)]
