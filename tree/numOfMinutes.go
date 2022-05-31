package tree

import (
	"fmt"
	"strings"
)

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	employees := make([]*Employee, n)
	for i := 0; i < n; i++ {
		employees[i] = &Employee{id: i, t: informTime[i], reports: make([]int, 0)}
	}

	for eid, mgrId := range manager {
		if mgrId > -1 {
			employees[mgrId].reports = append(employees[mgrId].reports, eid)
		}
	}

	// fmt.Printf("Employees: %v\n\n", employees)
	// return 0
	return minTimeToPropagate(employees[headID], employees)
}

type Employee struct {
	id, t   int
	reports []int
}

func (Employee *Employee) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{Id: %d, T: %d, Reports[", Employee.id, Employee.t))
	for _, id := range Employee.reports {
		sb.WriteString(fmt.Sprintf("%d, ", id))
	}
	sb.WriteString("]}\n")
	return sb.String()
}

func minTimeToPropagate(emp *Employee, employees []*Employee) int {

	if len(emp.reports) == 0 {
		return 0
	}

	// fmt.Printf("Employee: %v\n", emp)
	var min *int
	for _, reportId := range emp.reports {
		t := minTimeToPropagate(employees[reportId], employees) + emp.t
		if min == nil {
			min = new(int)
			*min = t
		} else if *min < t {
			*min = t
		}
	}

	return *min
}
