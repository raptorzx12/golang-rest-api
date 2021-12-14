package dto

//PirateUpdateDTO is a model that client use when updating a pirate
type PirateUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	DevilFruit string `json:"devil_fruit" form:"devil_fruit"`
	Crew       string `json:"crew" form:"crew" binding:"required"`
	Job        string `json:"job" form:"job" binding:"required"`
	UserID     uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

//PirateCreateDTO is a model that client use when create a new pirate
type PirateCreateDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	DevilFruit string `json:"devil_fruit" form:"devil_fruit"`
	Crew       string `json:"crew" form:"crew" binding:"required"`
	Job        string `json:"job" form:"job" binding:"required"`
	UserID     uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
