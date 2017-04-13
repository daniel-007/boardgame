package users

//Factored into a sub-package so we don't get a cycle of dependencies between
//bolt and user db.

type StorageRecord struct {
	//The Firebase user id
	Id string
}