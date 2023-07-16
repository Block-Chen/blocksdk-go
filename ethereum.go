package blocksdk_go

type Ethereum struct {
	*BaseEthereum
}

func NewEthereum(apiToken string, endpoint string) *Ethereum {
	return &Ethereum{
		BaseEthereum: NewBaseEthereum("eth", apiToken, endpoint, "v3"),
	}
}
