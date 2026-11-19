from enum import Enum
from uuid import UUID, uuid4

from pydantic import BaseModel, Field


class PedidoStatus(Enum):
    PENDING = "pending"
    CONFIRMED = "confirmed"
    DELIVERING = "delivering"
    DELIVERED = "delivered"
    CANCELED = "canceled"


class PedidoItem(BaseModel):
    product_id: UUID
    quantity: int
    price_snapshot: float


class Pedido(BaseModel):
    id: UUID = Field(default_factory=lambda _: uuid4())
    user_id: UUID
    restaurant_id: UUID
    items: list[PedidoItem]
    status: PedidoStatus = Field(default=PedidoStatus.PENDING)
