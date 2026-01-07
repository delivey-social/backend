CREATE TABLE IF NOT EXISTS restaurantes(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	cnpj VARCHAR(14) NOT NULL,
	nome VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS cardapio_itens(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	nome VARCHAR(100) NOT NULL,
	preco INTEGER NOT NULL,
	categoria VARCHAR(100) NOT NULL,
	restaurante_id UUID,
	FOREIGN KEY (restaurante_id) REFERENCES restaurantes(id)
);