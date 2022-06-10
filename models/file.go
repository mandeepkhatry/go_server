package models

import "encoding/json"

type FileData struct {
	ContentType string `json:"content_type"`
	Size        int    `json:"file_size"`
	FileName    string `json:"file_name"`
}

func (f *FileData) EncodeData() ([]byte, error) {
	data, err := json.Marshal(f)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func DecodeData(data []byte) (FileData, error) {
	var f FileData
	err := json.Unmarshal(data, &f)
	if err != nil {
		return FileData{}, err
	}
	return f, nil
}
