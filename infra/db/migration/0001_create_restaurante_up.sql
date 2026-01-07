CREATE TABLE restaurantes(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	cnpj VARCHAR(14) NOT NULL,
	nome VARCHAR(100) NOT NULL
)