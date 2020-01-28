package ch2

import "fmt"

type User struct {
	Name string
}

func (u User) IsAdmin() bool { return false }
func (u User) DisplayName() string {
	return u.Name
}

type Admin struct {
	User
}

func (a Admin) IsAdmin() bool { return true }
func (a Admin) DisplayName() string {
	// Add [Admin] in front of the embedded user’s display name
	return "[Admin] " + a.User.DisplayName()
}

// ============================================================================
// ============================================================================

func EmbeddedTypes() {
	u := User{"Normal User"}
	fmt.Println(u.Name)          // Normal User
	fmt.Println(u.DisplayName()) // Normal User
	fmt.Println(u.IsAdmin())     // false
	a := Admin{User{"Admin User"}}
	fmt.Println(a.Name)          // Admin User
	fmt.Println(a.User.Name)     // Admin User
	fmt.Println(a.DisplayName()) // [Admin] Admin User
	fmt.Println(a.IsAdmin())     // true
}
