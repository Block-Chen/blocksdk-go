package blocksdk_go

type EthereumClassic struct {
	*BaseEthereum
}

func NewEthereumClassic(apiToken string, endpoint string) *EthereumClassic {
	return &EthereumClassic{
		BaseEthereum: NewBaseEthereum("etc", apiToken, endpoint, "v3"),
	}
}
