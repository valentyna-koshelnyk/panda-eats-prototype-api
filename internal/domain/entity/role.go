package entity

type Role int

const (
	UserRole = iota
	AdminRole
)

// String returns the string representation of the Role
func (r Role) String() string {
	switch r {
	case UserRole:
		return "UserRole"
	case AdminRole:
		return "AdminRole"
	default:
		return "Unknown"
	}
}
