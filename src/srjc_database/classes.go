package srjc_database

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type sections struct {
	gorm.Model

	section_id int32
	short_name string
	long_name string
	description string
	units float32
	status string
	current_enrolled int32
	seats_remaining int32
	start_date int64
	end_date int64
	final_date int64

	times   []times
}

type times struct {
	gorm.Model
	id int32

	monday   bool
	tuesday   bool
	wednesday   bool
	thursday   bool
	friday   bool
	saturday   bool
	sunday   bool

	start_time int32
	end_time int32
	campus string
	room string
	section_id int32
}


func AllSections() {
	temps := database.Find(sections{})
	fmt.Print(temps)
}
