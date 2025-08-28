package pedido

var level1Bairros = []string{
	"Água Verde",
	"Batel",
	"Centro",
	"Rebouças",
	"Parolin",
	"Guaíra",
	"Portão",
	"Vila Izabel",
	"Seminário",
	"Santa Quitéria",
	"Bigorrilho",
}

var level2Bairros = []string{
	"Mossunguê",
	"Mercês",
	"São Francisco",
	"Centro Cívico",
	"Alto da Glória",
	"Alto da XV",
	"Cristo Rei",
	"Jardim Botânico",
	"Prado Velho",
	"Hauer",
	"Fanny",
	"Lindóia",
	"Novo Mundo",
	"Fazendinha",
	"Campo Comprido",
	"Cabral",
	"Capão Raso",
}

var level3Bairros = []string{
	"Orleans",
	"Santo Inácio",
	"Cascatinha",
	"Vista Alegre",
	"Bom retiro",
	"Ahú",
	"Hugo Lange",
	"Jardim Social",
	"Tarumã",
	"Capão da Imbuia",
	"Cajuru",
	"Jardim das Américas",
	"Guabirotuba",
	"Uberaba",
	"Boqueirão",
	"Xaxim",
	"Pinheirinho",
	"Cidade Insustrial",
	"Santa Felicidade",
}

const LEVEL_1_BASE_TARIFF = 500
const LEVEL_2_BASE_TARIFF = 800
const LEVEL_3_BASE_TARIFF = 1300

func PopulateBairros(s *BairroService) {
	for _, bairro := range level1Bairros {
		s.CreateBairro(bairro, LEVEL_1_BASE_TARIFF)
	}

	for _, bairro := range level2Bairros {
		s.CreateBairro(bairro, LEVEL_2_BASE_TARIFF)
	}

	for _, bairro := range level3Bairros {
		s.CreateBairro(bairro, LEVEL_3_BASE_TARIFF)
	}
}
