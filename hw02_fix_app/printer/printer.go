package printer

import (
	"fmt"

	"github.com/AndreiRubinJ/otusgo/hw02_fix_app/employee"
)

func PrintStaff(staff []employee.Employee) {
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID,
			staff[i].Age,
			staff[i].Name,
			staff[i].DepartmentID)
		fmt.Println(str)
	}
}
