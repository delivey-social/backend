package valueobject

type Preco struct {
	preco_itens  int
	taxa_app     int
	taxa_entrega int
}

func NewPreco(preco_itens int, taxa_entrega int) Preco {
	var taxa_app = int(float64(preco_itens) * 0.1)

	return Preco{
		preco_itens:  preco_itens,
		taxa_entrega: taxa_entrega,
		taxa_app:     taxa_app,
	}
}
