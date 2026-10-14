from abc import ABC, abstractmethod
from typing import Annotated
from uuid import UUID

from pydantic import BaseModel, Field


class UserGateway(ABC):
    @abstractmethod
    def exists(self, user_id: UUID) -> bool: ...


class RestaurantGateway(ABC):
    @abstractmethod
    def exists(self, restaurant_id: UUID) -> bool: ...

    @abstractmethod
    def item_exists(self, restaurant_id: UUID, product_id: UUID) -> bool: ...


class CreatePedidoRequestItem(BaseModel):
    product_id: UUID
    quantity: int = Field(..., gt=0)


class CreatePedidoRequest(BaseModel):
    user_id: UUID
    restaurant_id: UUID
    items: Annotated[list[CreatePedidoRequestItem], Field(..., min_length=1)]
