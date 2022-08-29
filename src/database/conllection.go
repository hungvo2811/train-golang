package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	departments = "departments"
	staffs      = "staffs"
)

func StaffCollection() *mongo.Collection {
	return db.Collection(staffs)
}

func DepartmentCollection() *mongo.Collection {
	return db.Collection(departments)
}
