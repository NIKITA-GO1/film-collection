package film

type Film struct {
	ID          int    `json: "id"`
	Name        string `json: "name" valid:"required~Film name is required, length(1|150)~Film name must be between 1 and 150 characters"`
	Discription string `json "discription" valid:"required~Film discription is required, length(0|1000)~ film discription cant be more than 1000"`
	ReleaseDate string `json: "release_date`
	Rate        int    `json: "rate" valid:"int, range(0|10)~ Film rate must be betweeo 0 and 10"`
	Actors      []int  `json: "actors"`
}
