package restaurante

import "comida.app/src/shared"

type Restaurante struct {
	nome string
	// TODO: CNPJ should be value object
	cnpj        string
	endereco    shared.Endereco
	responsavel shared.Usuario
	imagem_url  string
}

func NewRestaurante(nome string, cnpj string, cep string, responsavel shared.Usuario, imagem_url string) (Restaurante, error) {
	endereco := shared.Endereco{
		CEP:         cep,
		Rua:         "",
		Bairro:      "",
		Numero:      "",
		Complemento: "",
		Cidade:      "",
		UF:          "",
		Latitude:    0,
		Longitude:   0,
	}

	return Restaurante{
		nome:        nome,
		cnpj:        cnpj,
		endereco:    endereco,
		responsavel: responsavel,
		imagem_url:  imagem_url,
	}, nil
}
