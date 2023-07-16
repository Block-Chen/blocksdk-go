package blocksdk_go

type Klaytn struct {
	*BaseEthereum
}

func NewKlaytn(apiToken string, endpoint string) *Klaytn {
	return &Klaytn{
		BaseEthereum: NewBaseEthereum("klay", apiToken, endpoint, "v3"),
	}
}
