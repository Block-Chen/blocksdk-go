package blocksdk_go

type Polygon struct {
	*BaseEthereum
}

func NewPolygon(apiToken string, endpoint string) *Polygon {
	return &Polygon{
		BaseEthereum: NewBaseEthereum("matic", apiToken, endpoint, "v3"),
	}
}
