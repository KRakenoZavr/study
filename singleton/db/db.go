package db

type DB struct {
	name string
}

var instance *DB

func GetDB() *DB {
	if instance != nil {
		return instance
	}

	instance = &DB{}
	return instance
}
