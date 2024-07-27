package request

type GeneratePresignedURLsRequest struct {
	ObjectNameList []string `json:"objectNameList"`
}
