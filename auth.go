package ldbauth1

import ()

type LDBAuth struct {
	Users map[string]string
}

func (db *LDBAuth) CheckPasswd(user, pass string) (bool, error) {

	for key, value := range db.Users {
		if key == user && pass == value {
			return true, nil
		}
	}
	return false, nil
}

func (db *LDBAuth) AddUser(user, pass string) error {
	db.Users[user] = pass
	return nil
}

func (db *LDBAuth) ChgPass(user, pass string) error {
	db.Users[user] = pass
	return nil
}
