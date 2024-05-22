package elasticsearch

type ResposeES struct {
	Hits        Hits        `json:"hits"`
	Aggregation Aggregation `json:"aggregations"`
}

type Hits struct {
	Total interface{} `json:"total"`
	Hits  interface{} `json:"hits"`
}

type Aggregation struct {
	Data Data `json:"data"`
}

type Data struct {
	Buckets []interface{} `json:"buckets"`
}
