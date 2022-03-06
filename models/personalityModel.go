package models

type Personality struct {
	Id        int
	Answer1   string
	Answer2   string
	Scoreline int
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
