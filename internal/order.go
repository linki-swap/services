package internal

import "time"

type Order struct {
	SourceAsset string      `json:"sourceasset"`
	TargetAsset string      `json:"targetasset"`
	Status      string      `json:"Status"`
	Paid        bool        `json:"paid"`
	SourceValue float64     `json:"sourcevalue"`
	TargetValue float64     `json:"targetvalue"`
	Rate        float64     `json:"rate"`
	Stamp       time.Time   `json:"string"`
	Lane        Lane        `json:"lane"`
	Network     NetworkType `json:"networktype"`
}

type Status string

const (
	Pending    Status = "Pending"
	Confirming Status = "Confirming"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
