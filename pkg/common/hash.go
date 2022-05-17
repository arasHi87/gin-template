package common

import "golang.org/x/crypto/bcrypt"

/* Hash password by bcrypt
@Params
	password(string): string need to be hash

@Return
	hash(string): hashed password
	error(error): error
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/* Compare password and hash by bcrypt
@Params
	password(string): string need to be hash
	hash(string): hash need to be compare

@Return
	result(bool): compare result
*/
func CheckPassword(password string, hash string) bool {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(hashedPassword))
	return err == nil
}
