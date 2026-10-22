from typing import Optional
from uuid import UUID, uuid4

from sqlalchemy import ForeignKey, Numeric, String
from sqlalchemy.orm import Mapped, mapped_column, relationship

from src.core.infra.sqlalchemy_base import Base


class RestauranteModel(Base):
    __tablename__ = "restaurante"

    id: Mapped[UUID] = mapped_column(primary_key=True, default=uuid4)
    nome: Mapped[str] = mapped_column(String(100), nullable=False)

    # One-to-one relationship with Menu (if a restaurant has only one menu)
    menu: Mapped[Optional["MenuModel"]] = relationship(
        back_populates="restaurante", uselist=False, cascade="all, delete-orphan"
    )


class MenuModel(Base):
    __tablename__ = "menu"

    id: Mapped[UUID] = mapped_column(primary_key=True, default=uuid4)
    # Link back to the restaurant table
    restaurante_id: Mapped[UUID] = mapped_column(
        ForeignKey("restaurante.id"), nullable=False, unique=True
    )

    # Relationships
    restaurante: Mapped["RestauranteModel"] = relationship(back_populates="menu")

    # This matches your domain logic: menu.itens.append(data)
    itens: Mapped[list["MenuItemModel"]] = relationship(
        back_populates="menu", cascade="all, delete-orphan", lazy="selectin"
    )


class MenuItemModel(Base):
    __tablename__ = "menu_item"

    id: Mapped[UUID] = mapped_column(primary_key=True, default=uuid4)
    menu_id: Mapped[UUID] = mapped_column(ForeignKey("menu.id"), nullable=False)

    nome: Mapped[str] = mapped_column(String(100), nullable=False)
    preco: Mapped[int] = mapped_column(Numeric(10, 2), nullable=False)

    # Relationship back to Menu
    menu: Mapped["MenuModel"] = relationship(back_populates="itens")
