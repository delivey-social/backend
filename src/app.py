from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse

from src.core.errors import DomainError
from src.core.infra.lifespan import lifespan
from src.modules.pedido.infra.dependencies import get_pedido_service
from src.modules.pedido.presentation import PedidoRouter
from src.modules.restaurante.infra.dependencies import (
    get_restaurante_service,
    menu_service,
)
from src.modules.restaurante.presentation import MenuRouter, RestauranteRouter


def create_app() -> FastAPI:
    app = FastAPI(title="Comida App Backend", lifespan=lifespan)

    app.include_router(PedidoRouter(get_pedido_service))
    app.include_router(MenuRouter(menu_service))
    app.include_router(RestauranteRouter(get_restaurante_service))

    @app.get("/")
    async def ping() -> dict[str, str]:
        return {"message": "comida.app backend is alive!"}

    @app.exception_handler(DomainError)
    async def domain_error_handler(request: Request, exc: DomainError):
        return JSONResponse(
            status_code=exc.http_status_code,
            content={"message": str(exc)},
        )

    return app


app = create_app()
