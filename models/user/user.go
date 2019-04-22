package user

import (
	"errors"
	"time"

	"github.com/ostisense/api/postgres"
	stringUtils "github.com/ostisense/api/utils/string_utils"
	"golang.org/x/crypto/bcrypt"
)

type Email string
type PlainPassword string
type BCryptedPassword string
type SecureToken string

type dbUser struct {
	ID               postgres.BigSerial `db:"id"`
	Email            Email              `db:"email"`
	BcryptedPassword BCryptedPassword   `db:"bcrypted_password"`
	Token            SecureToken        `db:"token"`
	CreatedAt        time.Time          `db:"created_at"`
}

var ErrInvalidEmail = errors.New("invalid email")
var ErrInvalidPassword = errors.New("invalid password")
var ErrMismatchedEmailAndPassword = errors.New("email and password do not match")

// returns non-nil err unless password matches bcrypted password
func (self *dbUser) matchesPassword(password PlainPassword) error {
	passwordBytes := []byte(password)
	bcryptedPasswordBytes := []byte(self.BcryptedPassword)
	return bcrypt.CompareHashAndPassword(bcryptedPasswordBytes, passwordBytes)
}

func fetchDBUserByToken(token SecureToken) (*dbUser, error) {
	var user dbUser
	queryString := "select * from users where token = $1"
	err := postgres.SharedDB().QueryRowx(queryString, token).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func fetchDBUserByEmail(email Email) (*dbUser, error) {
	var user dbUser
	queryString := "select * from users where email = $1"
	err := postgres.SharedDB().QueryRowx(queryString, email).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getBcryptCost() int {
	return 16
}

func getBcryptedPasswordFromPassword(password PlainPassword) (bcryptedPassword BCryptedPassword, err error) {
	passwordBytes := []byte(password)
	bcryptedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, getBcryptCost())
	if err != nil {
		return "", err
	}
	bcryptedPassword = BCryptedPassword(bcryptedPasswordBytes[:])
	return bcryptedPassword, nil
}

func getNewToken() (token SecureToken, err error) {
	tokenString, err := stringUtils.GenerateSecureRandomString(32)
	if err != nil {
		return "", err
	}
	return SecureToken(tokenString), nil
}

func CreateUser(email Email, password PlainPassword) (*User, error) {
	isValidEmail := stringUtils.IsValidEmail(string(email))
	if !isValidEmail {
		return nil, ErrInvalidEmail
	}

	isValidPassword := stringUtils.IsValidPassword(string(password))
	if !isValidPassword {
		return nil, ErrInvalidPassword
	}

	bcryptedPassword, err := getBcryptedPasswordFromPassword(password)
	if err != nil {
		return nil, err
	}

	token, err := getNewToken()
	if err != nil {
		return nil, err
	}

	createdAt := time.Now()
	_, err = postgres.SharedDB().Exec(`
		insert into users
			(email, bcrypted_password, token, created_at)
		values ($1, $2, $3, $4)
	`, email, bcryptedPassword, token, createdAt)
	if err != nil {
		return nil, err
	}

	user, err := FetchUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FetchUserByToken(token SecureToken) (*User, error) {
	dbUser, err := fetchDBUserByToken(token)
	if err != nil {
		return nil, err
	}
	return makeUserFromDBUser(dbUser), nil
}

func FetchUserByEmail(email Email) (*User, error) {
	dbUser, err := fetchDBUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return makeUserFromDBUser(dbUser), nil
}

func FetchUserByEmailMatchingPassword(email Email, password PlainPassword) (*User, error) {
	dbUser, err := fetchDBUserByEmail(email)
	if err != nil {
		return nil, ErrMismatchedEmailAndPassword
	}
	err = dbUser.matchesPassword(password)
	if err != nil {
		return nil, ErrMismatchedEmailAndPassword
	}
	return makeUserFromDBUser(dbUser), nil
}

// public-facing type that only exposes the public fields
type User struct {
	ID    postgres.BigSerial `json:"id"`
	Email Email              `json:"email"`
	Token SecureToken        `json:"token"`
}

func (self *User) Validate() error {
	if len(self.Email) == 0 {
		return errors.New("User Invalid: no email")
	}
	return nil
}

func makeUserFromDBUser(user *dbUser) *User {
	return &User{
		ID:    user.ID,
		Email: user.Email,
		Token: user.Token,
	}
}
