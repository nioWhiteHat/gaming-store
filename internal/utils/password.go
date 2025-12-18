package utils

import (

	"golang.org/x/crypto/bcrypt"
)
type Password struct{
	Plainpsw string
	Hashed []byte
}

func (p *Password ) Set(txtpsw string) error {
	hash,err := bcrypt.GenerateFromPassword([]byte(txtpsw),12)
	if err!=nil{
		return err
	}
	p.Hashed = hash
	p.Plainpsw = txtpsw
	return nil 

}

func (p *Password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hashed, []byte(plaintextPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func Hash(txt string) (string,error){
	hash,err := bcrypt.GenerateFromPassword([]byte(txt),12)
	if err!= nil{
		return "", err
	}
	return string(hash),nil
}