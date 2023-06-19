package schema

type MathOperations struct {
	CheckPrime   int   `json:"CheckPrime"`
	SumArray     []int `json:"SumArray"`
	ProductArray []int `json:"ProductArray"`
}

type Response struct {
	Prime      string `json:"Prime"`
	Sum        int    `json:"Sum"`
	Product    int    `json:"Product"`
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}
