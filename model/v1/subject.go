package model

type Subject struct {
	SubUrl  string `json:"sub_url"`
	SubType string `json:"sub_type"`
}

type SubjectList struct {
	ListMeta `json:",inline"`
	Items    []Subject `json:"items"`
}
