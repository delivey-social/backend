package enums

type PedidoStatus string

const (
	StatusAguardandoPagamento PedidoStatus = "AGUARDANDO_PAGAMENTO"
	StatusProduzindo          PedidoStatus = "PRODUZINDO"
	StausAguardandoRetirada   PedidoStatus = "AGUARDANDO_RETIRADA"
	StatusEntregando          PedidoStatus = "ENTREGANDO"
	StatusFinalizado          PedidoStatus = "FINALIZADO"
	StatusCancelado           PedidoStatus = "CANCELAR"
)
