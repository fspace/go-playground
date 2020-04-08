package options

type User struct {
	Id, Sex int
	Name    string
}

func NewUser(f func(u *User)) *User {
	u := new(User)
	f(u)
	return u
}
func WithName(name string) func(u *User) {
	return func(u *User) {
		u.Name = name
	}
}

// NewUser2 多个可选的属性赋值参数
func NewUser2(opts ...func(u *User)) *User {
	u := new(User)
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(u)
		}
	}
	return u
}

// 每次写一个函数签名比较长 可以专门弄个类型 避免冗长的签名形式
type UserOption func(u *User)

func NewUser3(opts ...UserOption) *User {
	u := new(User)
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(u)
		}
	}
	return u
}

func WithSex(sex int) UserOption {
	return func(u *User) {
		u.Sex = sex
	}
}
