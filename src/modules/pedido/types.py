from abc import ABC, abstractmethod
from typing import Annotated
from uuid import UUID

from pydantic import BaseModel, Field


class UserGateway(ABC):
    @abstractmethod
    async def exists(self, user_id: UUID) -> bool: ...


class RestaurantItem(BaseModel):
    product_id: UUID
    name: str
    price: int


class RestaurantGateway(ABC):
    @abstractmethod
    async def exists(self, restaurant_id: UUID) -> bool: ...

    @abstractmethod
    async def get_item(
        self, restaurant_id: UUID, product_id: UUID
    ) -> RestaurantItem: ...


class CreatePedidoRequestItem(BaseModel):
    product_id: UUID
    quantity: int = Field(..., gt=0)


class CreatePedidoRequest(BaseModel):
    user_id: UUID
    restaurant_id: UUID
    items: Annotated[list[CreatePedidoRequestItem], Field(..., min_length=1)]


class UpdatePedidoRequest(BaseModel):
    user_id: UUID
    restaurant_id: UUID
    items: Annotated[list[CreatePedidoRequestItem], Field(..., min_length=1)]
