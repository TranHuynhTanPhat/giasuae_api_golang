package entities

type Salaryinfo struct {
	ID             int    `json:"id"`
	Type_teacher   int    `json:"type_teacher"`
	Two_sessions   string `json:"two_sessions"`
	Three_sessions string `json:"three_sessions"`
	Four_sessions  string `json:"four_sessions"`
	Five_sessions  string `json:"five_sessions"`
	ID_category    int    `json:"id_category"`
	Created_at     string `json:"created_at"`
}
type SalaryinfoView struct {
	ID             int    `json:"id"`
	Type_teacher   int    `json:"type_teacher"`
	Two_sessions   string `json:"two_sessions"`
	Three_sessions string `json:"three_sessions"`
	Four_sessions  string `json:"four_sessions"`
	Five_sessions  string `json:"five_sessions"`
	ID_category    string `json:"id_category"`
	Created_at     string `json:"created_at"`
}
type SalaryinfoDetail struct {
	ID             int      `json:"id"`
	Type_teacher   int      `json:"type_teacher"`
	Two_sessions   string   `json:"two_sessions"`
	Three_sessions string   `json:"three_sessions"`
	Four_sessions  string   `json:"four_sessions"`
	Five_sessions  string   `json:"five_sessions"`
	ID_category    Category `json:"id_category"`
	Created_at     string   `json:"created_at"`
}
