package entities

type NewclassesDetail struct {
	ID         int        `json:"id"`
	Address    string     `json:"address"`
	District   string     `json:"district"`
	Sobuoi     int        `json:"sobuoi"`
	Time       string     `json:"time"`
	Salary     int        `json:"salary"`
	Require    string     `json:"require"`
	Status     int        `json:"status"`
	Contact    string     `json:"contact"`
	Created_at string     `json:"created_at"`
	Subjects   []Subject  `json:"subjects"`
	Classes    []Class    `json:"classes"`
	Categories []Category `json:"categories"`
}

type ClassesOfNewclasses struct {
	ID_class    int `json:"id_class"`
	ID_newclass int `json:"id_newclass"`
}

type SubjectsOfNewclasses struct {
	ID_subject  int `json:"id_subject"`
	ID_newclass int `json:"id_newclass"`
}

type CategoriesOfNewclasses struct {
	ID_category int `json:"id_category"`
	ID_newclass int `json:"id_newclass"`
}

type NewclasssesSet struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	District   string `json:"district"`
	Sobuoi     int    `json:"sobuoi"`
	Time       string `json:"time"`
	Salary     int    `json:"salary"`
	Require    string `json:"require"`
	Status     int    `json:"status"`
	Contact    string `json:"contact"`
	Subjects   string `json:"subjects"`
	Classes    string `json:"classes"`
	Categories string `json:"categories"`
	Created_at string `json:"created_at"`
}

type NewClassesReq struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	District   string `json:"district"`
	Sobuoi     int    `json:"sobuoi"`
	Time       string `json:"time"`
	Salary     int    `json:"salary"`
	Require    string `json:"require"`
	Status     int    `json:"status"`
	Contact    string `json:"contact"`
	Created_at string `json:"created_at"`
	Subjects   []int  `json:"subjects"`
	Classes    []int  `json:"classes"`
	Categories []int  `json:"categories"`
}

type NewClassesDefault struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	District   string `json:"district"`
	Sobuoi     int    `json:"sobuoi"`
	Time       string `json:"time"`
	Salary     int    `json:"salary"`
	Require    string `json:"require"`
	Status     int    `json:"status"`
	Contact    string `json:"contact"`
	Created_at string `json:"created_at"`
}
