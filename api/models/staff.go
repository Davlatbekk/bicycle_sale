package models

type Staff struct {
	StaffId     int    `json:"staff_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Active      int    `json:"active"`
	StoreId     int    `json:"store_id"`
	StoreData   *Store `json:"store_data"`
	ManagerId   int    `json:"manager_id"`
	ManagerData *Staff `json:"manager_data"`
}

type StaffPrimaryKey struct {
	StaffId int `json:"staff_id"`
}

type CreateStaff struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Active    int    `json:"active"`
	StoreId   int    `json:"store_id"`
	ManagerId int    `json:"manager_id"`
}

type UpdateStaff struct {
	StaffId   int    `json:"staff_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Active    int    `json:"active"`
	StoreId   int    `json:"store_id"`
	ManagerId int    `json:"manager_id"`
}

type GetListStaffRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStaffResponse struct {
	Count  int      `json:"count"`
	Staffs []*Staff `json:"staffs"`
}

type GetListReportStaffRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type Report struct {
	FullName  string  `json:"full_name"`
	Category  string  `json:"category"`
	Product   string  `json:"product"`
	Count     int     `json:"count"`
	TotalSumm float64 `json:"total_summ"`
	Date      string  `json:"date"`
}
type GetListReportStaffResponse struct {
	Count   int       `json:"count"`
	Reports []*Report `json:"reports"`
}
