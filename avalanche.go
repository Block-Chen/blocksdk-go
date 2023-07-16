package blocksdk_go

type Avalanche struct {
	*BaseEthereum
}

func NewAvalanche(apiToken string, endpoint string) *Avalanche {
	return &Avalanche{
		BaseEthereum: NewBaseEthereum("avax", apiToken, endpoint, "v3"),
	}
}
