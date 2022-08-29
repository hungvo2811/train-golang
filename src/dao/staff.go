package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"test/src/database"
	"test/src/models"
	"test/src/models/query"
)

func CreateStaff(ctx context.Context, staff models.Staff) (models.Staff, error) {

	var (
		staffCol = database.StaffCollection()
	)
	_, err := staffCol.InsertOne(ctx, staff)

	return staff, err
}

func UpdateStaff(ctx context.Context, id primitive.ObjectID, staff models.Staff) (models.Staff, error) {
	var (
		staffCol = database.StaffCollection()
		filter   = bson.D{{"_id", id}}
	)

	_, err := staffCol.UpdateOne(ctx, filter, staff)

	return staff, err
}

func DeleteStaff(ctx context.Context, id primitive.ObjectID) error {
	var (
		staffCol = database.StaffCollection()
		filter   = bson.D{{"_id", id}}
	)

	_, err := staffCol.DeleteOne(ctx, filter)

	return err
}

func GetStaff(ctx context.Context, id primitive.ObjectID) (staff models.Staff, err error) {
	var (
		staffCol = database.StaffCollection()
		filter   = bson.D{{"_id", id}}
	)

	err = staffCol.FindOne(ctx, filter).Decode(&staff)

	return staff, err
}

func GetStaffs(ctx context.Context, query query.StaffQuery) (staffs []models.Staff, err error) {
	var staffCol = database.StaffCollection()
	var filter = bson.M{}

	if !query.Department.IsZero() {
		filter["departmentId"] = query.Department
	}

	if len(query.Keyword) > 0 {
		filter["name"] = bson.M{
			"$regex": query.Keyword,
		}
	}

	opts := options.Find().SetLimit(query.Limit).SetSkip((query.Page - 1) * query.Limit)
	result, err := staffCol.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	for result.Next(context.Background()) {
		var staff models.Staff
		err = result.Decode(&staff)
		if err != nil {
			log.Fatal(err)
		}
		staffs = append(staffs, staff)
	}

	return staffs, err
}

func GetStaffByEmail(ctx context.Context, email string) bool {
	var (
		staffCol = database.StaffCollection()
		filter   = bson.D{{"email", email}}
		staff    models.Staff
	)

	err := staffCol.FindOne(ctx, filter).Decode(&staff)

	if err != nil {
		return false
	}
	return true
}

func GetStaffByEmailAndPassword(ctx context.Context, email string, password string) bool {
	var (
		staffCol = database.StaffCollection()
		filter   = bson.D{{"email", email}, {"password", password}}
		staff    models.Staff
	)

	err := staffCol.FindOne(ctx, filter).Decode(&staff)

	if err != nil {
		return false
	}
	return true
}
