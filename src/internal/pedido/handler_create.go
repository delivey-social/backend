package pedido

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PedidoHandler) create(c *gin.Context) {
	var body CreatePedidoDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "Requisição Inválida",
		})
		return
	}

	usuario, err := createUserVO(body.User.Email, body.User.Phone, body.User.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = createAddressVO(body.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := h.service.Create(body.RestaurantID, body.Items, *usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Pedido criado com sucesso!",
		"id":       id,
	})
}

func createUserVO(email string, phone string, name string) (*Usuario, error) {
	userEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	userPhone, err := NewTelefone(phone)
	if err != nil {
		return nil, err
	}

	usuario := NewUsuario(userEmail, userPhone, name)

	return &usuario, nil
}

func createAddressVO(data AddressDTO) (*Endereco, error) {
	cep, err := NewCEP(data.CEP)
	if err != nil {
		return nil, err
	}

	address, err := NewEndereco(cep, data.Street, data.Neighborhood, data.Number, data.Observation)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
