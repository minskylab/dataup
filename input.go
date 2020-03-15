package dataup

type DiseasesPayload struct {
	MusclePain      float64 `json:"muscle_pain"`
	Breath          float64 `json:"breath"`
	NasalCongestion float64 `json:"nasal_congestion"`
	Headache        float64 `json:"head_ache"`
	BoneAche        float64 `json:"bone_ache"`
	SoreThroat      float64 `json:"sore_throat"`
	Tiredness       float64 `json:"tiredness"`
	DryCaught       float64 `json:"dry_cought"`
}

type DiseasesWeight struct {
	MusclePain      float64 `json:"muscle_pain"`
	Breath          float64 `json:"breath"`
	NasalCongestion float64 `json:"nasal_congestion"`
	Headache        float64 `json:"head_ache"`
	BoneAche        float64 `json:"bone_ache"`
	SoreThroat      float64 `json:"sore_throat"`
	Tiredness       float64 `json:"tiredness"`
	DryCaught       float64 `json:"dry_cought"`
}

