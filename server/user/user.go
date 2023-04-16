package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"-"`
	Bio            string    `json:"bio"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Verification
type CreateUserOpts struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser creates brand new user with new ID!
func CreateUser(opts CreateUserOpts) (User, error) {
	now := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(opts.Password), 6)
	return User{
		ID:             uuid.New(),
		Username:       opts.Username,
		Email:          opts.Email,
		HashedPassword: hashedPassword,
		Bio:            "",
		CreatedAt:      now,
		UpdatedAt:      now,
	}, err
}

func (u *User) UpdatePassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return err
	}
	u.HashedPassword = hashedPassword
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdateBio(bio string) {
	u.Bio = bio
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}
