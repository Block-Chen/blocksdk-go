package blocksdk_go

type BinanceSmart struct {
	*BaseEthereum
}

func NewBinanceSmart(apiToken string, endpoint string) *BinanceSmart {
	return &BinanceSmart{
		BaseEthereum: NewBaseEthereum("bsc", apiToken, endpoint, "v3"),
	}
}
