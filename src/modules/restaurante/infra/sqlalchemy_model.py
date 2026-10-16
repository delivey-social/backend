from uuid import UUID

from sqlalchemy import String
from sqlalchemy.orm import Mapped, mapped_column

from src.core.infra.sqlalchemy_base import Base


class RestauranteModel(Base):
    __tablename__ = "restaurante"

    id: Mapped[UUID] = mapped_column(primary_key=True)
    nome: Mapped[str] = mapped_column(String(30), nullable=False)
