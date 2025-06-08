package wasendcore


// Response for count operations.
type CountResponse struct {
	// The total count.
	Total *float64 `field:"required" json:"total" yaml:"total"`
	// Count breakdown by status.
	//
	// Example:
	//   {
	//     "ACTIVE": "10",
	//     "ARCHIVED": "5",
	//     "DELETED": "2"
	//   }
	//
	ByStatus *map[string]*string `field:"optional" json:"byStatus" yaml:"byStatus"`
}

