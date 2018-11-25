package models

// MediaDynamoRecord --
type MediaDynamoRecord struct {
	AbrBucket      string   `json:"abrBucket"`
	Dash           []int    `json:"dash"`
	DashPlaylist   string   `json:"dashPlaylist"`
	DashURL        string   `json:"dashUrl"`
	EcodeJobID     string   `json:"ecodeJobId"`
	EndTime        string   `json:"EndTime"`
	FrameCapture   bool     `json:"frameCapture"`
	FrameHeight    int      `json:"frameHeight"`
	FrameWdith     int      `json:"frameWdith"`
	GUID           string   `json:"guid"`
	Hls            []int    `json:"hls"`
	HlsPlaylist    string   `json:"hlsPlaylist"`
	HlsURL         string   `json:"hlsUrl"`
	Mp4            []int    `json:"mp4"`
	Mp4Bucket      string   `json:"mp4Bucket"`
	Mp4Outputs     []string `json:"mp4Outputs"`
	SrcBucket      string   `json:"srcBucket"`
	SrcHeight      int      `json:"srcHeight"`
	SrcMediainfo   string   `json:"srcMediainfo"`
	SrcVideo       string   `json:"srcVideo"`
	SrcWidth       int      `json:"srcWidth"`
	StartTime      string   `json:"startTime"`
	WorkflowStatus string   `json:"workflowStatus"`
}
