package actor

type Actor struct {
	ID        int    `json: "id"`
	Name      string `json: "name" valid: "required~Actor name is required, length(1|50)~Actor name must be between 1 and 50 characters"`
	Gender    string `json: "gender" valid: "required~Actor gender is required, length(4|6)~Actor gender must be male or female"`
	BirthDate string `json: "birth_date" valid: "required~Actor bitrh_date is required, length(10|10)~Actor birth_date must be format dd.mm.yyyy"`
}
