import os

from dotenv import load_dotenv
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse

from src.constants import ROOT_PATH
from src.core.errors import DomainError
from src.core.infra.lifespan import lifespan
from src.modules.pedido.infra.dependencies import get_pedido_service
from src.modules.pedido.presentation import PedidoRouter
from src.modules.restaurante.infra.dependencies import (
    get_menu_service,
    get_restaurante_service,
)
from src.modules.restaurante.presentation import MenuRouter, RestauranteRouter

load_dotenv(ROOT_PATH / ".env")


def get_app_version() -> str:
    return "0.0.0"


def create_app() -> FastAPI:
    app = FastAPI(title="Comida App Backend", lifespan=lifespan)

    app.include_router(PedidoRouter(get_pedido_service))
    app.include_router(MenuRouter(get_menu_service))
    app.include_router(RestauranteRouter(get_restaurante_service))

    @app.get("/")
    async def ping() -> dict[str, str]:
        return {"message": "comida.app backend is alive!"}

    @app.get("/version")
    async def version() -> dict[str, str]:
        return {"version": get_app_version()}

    @app.exception_handler(DomainError)
    async def domain_error_handler(request: Request, exc: DomainError):
        return JSONResponse(
            status_code=exc.http_status_code,
            content={"message": str(exc)},
        )

    return app


def start_server():
    import uvicorn

    HOST = os.getenv("HOST", "0.0.0.0")
    PORT = int(os.getenv("PORT", "8000"))

    uvicorn.run(app, host=HOST, port=PORT)


app = create_app()
