package model

type imp struct {
	impId string
	dspId string
	price string
}

type okResponse struct {
	Status string
}

type dspResponse struct {
	AdId     string
	DspId    int
	ImpId    string
	Price    int
	Creative string
}
