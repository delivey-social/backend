import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_list_empty_restaurants(client: AsyncClient):
    response = await client.get("/restaurante/")

    assert response.status_code == 200
    assert response.json() == []


@pytest.mark.anyio
async def test_list_restaurants(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.get("/restaurante/")

    assert response.status_code == 200

    restaurants = response.json()

    assert len(restaurants) == 1
    assert restaurants[0]["nome"] == "Restaurante Teste"


@pytest.mark.anyio
async def test_list_restaurants_after_update(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    restaurant_id = create_restaurante.id

    updated_name = "Restaurante Atualizado"

    await client.put(
        f"/restaurante/{restaurant_id}",
        json={"nome": updated_name},
    )

    response = await client.get("/restaurante/")

    assert response.status_code == 200

    restaurants = response.json()

    assert len(restaurants) == 1
    assert restaurants[0]["nome"] == updated_name


@pytest.mark.anyio
async def test_list_restaurants_after_delete(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    restaurant_id = create_restaurante.id

    await client.delete(f"/restaurante/{restaurant_id}")

    response = await client.get("/restaurante/")

    assert response.status_code == 200
    assert response.json() == []


@pytest.mark.anyio
async def test_list_restaurants_after_creating_multiple(
    client: AsyncClient,
):
    restaurant_names = ["Restaurante A", "Restaurante B", "Restaurante C"]

    for name in restaurant_names:
        await client.post(
            "/restaurante/",
            json={"nome": name},
        )

    response = await client.get("/restaurante/")

    assert response.status_code == 200

    restaurants = response.json()

    assert len(restaurants) == len(restaurant_names)

    returned_names = [restaurant["nome"] for restaurant in restaurants]
    assert set(returned_names) == set(restaurant_names)
