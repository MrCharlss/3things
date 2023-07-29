package types

type CreteThingDTO struct {
	Win     string `json:"win"`
	Unhappy string `json:"unhappy"`
	Lesson  string `json:"lesson"`
    DayScore uint `json:"dayScore"`
    OtherScore uint `json:"otherScore"`
}

type Thing struct {
	Win     string `json:"win"`
	Unhappy string `json:"unhappy"`
	Lesson  string `json:"lesson"`
    DayScore uint `json:"dayScore"`
    OtherScore uint `json:"otherScore"`
}

type ThingResponse struct {
    ID uint `json:"id"`
 	Win     string `json:"win"`
	Unhappy string `json:"unhappy"`
	Lesson  string `json:"lesson"`
    DayScore uint `json:"dayScore"`
    OtherScore uint `json:"otherScore"`  
}
type GetThingsResponse struct {
    Things *[]ThingResponse `json:"things"`
}

type ThingCreateResponse struct {
    Thing *ThingResponse
}
