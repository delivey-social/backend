from pathlib import Path
from uuid import uuid4

import pytest
from httpx import AsyncClient

from src.modules.restaurante.service import MAX_IMAGE_SIZE, MAX_IMAGE_SIZE_MB
from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_upload_logo(
    client: AsyncClient, create_restaurante: CreateRestauranteResponse, tmp_path: str
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file = b"fake image data"  # Simulating a fake image file

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    assert res.status_code == 200

    expected_file_path = Path(tmp_path) / "logos" / f"{id}.png"

    assert expected_file_path.exists()
    assert expected_file_path.read_bytes() == file


@pytest.mark.anyio
async def test_upload_logo_multiple_files(
    client: AsyncClient, create_restaurante: CreateRestauranteResponse, tmp_path: str
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file1 = b"fake image data 1"
    file2 = b"fake image data 2"

    res = await client.put(
        f"/restaurante/{id}/logo",
        files=[
            ("file", ("logo1.png", file1, "image/png")),
            ("file", ("logo2.png", file2, "image/png")),
        ],
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Apenas um arquivo de imagem deve ser enviado",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_no_file(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={},
    )

    assert res.status_code == 422

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_invalid_restaurante_id(
    client: AsyncClient,
    tmp_path: str,
):
    invalid_id = "invalid-uuid"

    file = b"fake image data"

    res = await client.put(
        f"/restaurante/{invalid_id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    assert res.status_code == 422

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_restaurante_not_found(
    client: AsyncClient,
    tmp_path: str,
):
    non_existent_id = uuid4()

    file = b"fake image data"

    res = await client.put(
        f"/restaurante/{non_existent_id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    assert res.status_code == 404
    assert res.json() == {
        "message": "Restaurante não encontrado",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_unsupported_format(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file = b"fake image data"

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.txt", file, "text/plain")},
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Formato de imagem não suportado",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_empty_file(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    empty_file = b""

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", empty_file, "image/png")},
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Arquivo de imagem vazio",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_too_large_file(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    large_file = b"a" * (MAX_IMAGE_SIZE + 1)

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", large_file, "image/png")},
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": f"Arquivo de imagem muito grande (máximo {MAX_IMAGE_SIZE_MB}MB)",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_replace_existing_logo(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file1 = b"fake image data 1"
    file2 = b"fake image data 2"

    # Upload the first logo
    res1 = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", file1, "image/png")},
    )

    assert res1.status_code == 200

    expected_file_path = Path(tmp_path) / "logos" / f"{id}.png"
    assert expected_file_path.exists()
    assert expected_file_path.read_bytes() == file1

    # Upload the second logo to replace the first one
    res2 = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.jpg", file2, "image/png")},
    )

    new_expected_file_path = Path(tmp_path) / "logos" / f"{id}.jpg"

    assert res2.status_code == 200

    assert not expected_file_path.exists()  # The old file should be removed
    assert new_expected_file_path.exists()
    assert new_expected_file_path.read_bytes() == file2
