package model

type FileUploadInfo struct {
	Filename       string `json:"filename"`
	FileNameOrigin string `json:"fileNameOrigin"`
	Path           string `json:"uploadFilePath"`
	Ext            string `json:"ext"`
}
