package whiteboard

type SaveWhiteboardRequest struct {
	Scene map[string]any `json:"scene"`
}