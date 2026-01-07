package restaurante

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type SQLRestauranteRepository struct {
	db *sql.DB
}

func NewSQLRestauranteRepository(db *sql.DB) RestauranteRepository {
	return SQLRestauranteRepository{
		db: db,
	}
}

func (r SQLRestauranteRepository) List() []Restaurante {
	return make([]Restaurante, 0)
}
func (r SQLRestauranteRepository) Create(cnpj CNPJ, name string) (uuid.UUID, error) {
	var id uuid.UUID

	err := r.db.QueryRow(`
        INSERT INTO restaurantes (cnpj, nome)
        VALUES ($1, $2) 
        RETURNING id
    `, cnpj.String(), name).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}
func (r SQLRestauranteRepository) GetMenu(restauranteId uuid.UUID) (*Cardapio, error) {
	var cardapio Cardapio
	return &cardapio, nil
}
func (r SQLRestauranteRepository) GetItemsByIDs(restauranteId uuid.UUID, ids []uuid.UUID) (*[]CardapioItem, error) {
	var res []CardapioItem = make([]CardapioItem, 0, len(ids))

	rows, err := r.db.Query(`
        SELECT id, nome, preco, categoria
        FROM cardapio_itens
        WHERE id = ANY($1)
    `, ids)
	if err != nil {
		return &res, fmt.Errorf("error fetching items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item CardapioItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Category); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return &res, nil
}
func (r SQLRestauranteRepository) CreateMenuItem(restauranteId uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	var id uuid.UUID

	err := r.db.QueryRow(`
        INSERT INTO cardapio_itens (nome, preco, categoria, restaurante_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, data.Name, data.Price, data.Category, restauranteId).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}
func (r SQLRestauranteRepository) UpdateMenuItem(restauranteId uuid.UUID, id uuid.UUID, data MenuItemParams) error {
	err := r.db.QueryRow(`
        UPDATE cardapio_itens
        SET nome = $1, preco = $2, categoria = $3
        WHERE id = $4 AND restaurante_id = $5
    `, data.Name, data.Price, data.Category, id, restauranteId).Scan()
	if err != nil {
		return fmt.Errorf("error updating menu item: %w", err)
	}

	return nil
}
func (r SQLRestauranteRepository) DeleteMenuItem(restauranteId uuid.UUID, id uuid.UUID) error {
	err := r.db.QueryRow(`
        DELETE FROM cardapio_itens
        WHERE id = $1 AND restaurante_id = $2
    `, id, restauranteId).Scan()
	if err != nil {
		return fmt.Errorf("error deleting menu item: %w", err)
	}

	return nil
}
