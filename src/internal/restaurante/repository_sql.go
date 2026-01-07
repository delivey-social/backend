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
	var items []CardapioItem
	return &items, nil
}
func (r SQLRestauranteRepository) CreateMenuItem(restauranteId uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	return uuid.New(), nil
}
func (r SQLRestauranteRepository) UpdateMenuItem(restauranteId uuid.UUID, id uuid.UUID, data MenuItemParams) error {
	return nil
}
func (r SQLRestauranteRepository) DeleteMenuItem(restauranteId uuid.UUID, id uuid.UUID) error {
	return nil
}
