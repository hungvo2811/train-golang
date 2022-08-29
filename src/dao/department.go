package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"test/src/database"
	"test/src/models"
	"test/src/models/query"
)

func CreateDepartment(ctx context.Context, department models.Department) (models.Department, error) {
	var (
		departmentCol = database.DepartmentCollection()
	)
	_, err := departmentCol.InsertOne(ctx, department)

	return department, err
}

func UpdateDepartment(ctx context.Context, id primitive.ObjectID, department models.Department) (models.Department, error) {
	var (
		departmentCol = database.DepartmentCollection()
		filter        = bson.D{{"_id", id}}
	)

	_, err := departmentCol.UpdateOne(ctx, filter, department)

	return department, err
}

func DeleteDepartment(ctx context.Context, id primitive.ObjectID) error {
	var (
		departmentCol = database.DepartmentCollection()
		filter        = bson.D{{"_id", id}}
	)

	_, err := departmentCol.DeleteOne(ctx, filter)

	return err
}

func GetDepartment(ctx context.Context, id primitive.ObjectID) (existingDepartment models.Department, err error) {
	var (
		departmentCol = database.DepartmentCollection()
		filter        = bson.D{{"_id", id}}
	)

	err = departmentCol.FindOne(ctx, filter).Decode(&existingDepartment)
	fmt.Println(err)
	return existingDepartment, err
}

func GetDepartments(ctx context.Context, query query.DepartmentQuery) (departments []models.Department, err error) {
	var departmentCol = database.DepartmentCollection()
	var filter = bson.M{}

	if len(query.Keyword) > 0 {
		filter["name"] = bson.M{
			"$regex": query.Keyword,
		}
	}

	opts := options.Find().SetLimit(query.Limit).SetSkip((query.Page - 1) * query.Limit)
	result, err := departmentCol.Find(ctx, filter, opts)
	if err != nil {
		return departments, err
	}
	for result.Next(context.Background()) {
		var department models.Department
		err = result.Decode(&department)
		if err != nil {
			log.Fatal(err)
		}
		departments = append(departments, department)
	}

	return departments, err
}

func GetDepartmentByName(name string) bool {
	var (
		departmentCol = database.DepartmentCollection()
		ctx           context.Context
	)
	var department models.Department
	var filter = bson.D{{"name", name}}
	var err = departmentCol.FindOne(ctx, filter).Decode(&department)
	if err != nil {
		return false
	}
	return true
}
