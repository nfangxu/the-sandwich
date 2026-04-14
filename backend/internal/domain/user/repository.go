package user

// UserRepository defines the interface for user persistence
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}
