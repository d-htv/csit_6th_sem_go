package repositories

import (
	"class_proj_secb/common"
	"class_proj_secb/database/myquery"
	"class_proj_secb/models"
	"context"
)

func GetAllUsers(ctx context.Context) ([]*models.User, error){
	userQuery := myquery.User
	db := &common.MyApiServer.DB
	userQuery.UseDB(db)
	allUsers, err := userQuery.WithContext(ctx).Find()
	if err != nil{
		return nil, err
	}
	return allUsers, nil
}
func GetUserById(ctx context.Context, id uint64)(*models.User, error){
	userQuery := myquery.User
	userQuery.UseDB(&common.MyApiServer.DB)
	user,err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).First()
	if err != nil{
		return nil, err
	}
	return user, nil	
}

func CreateUser(ctx context.Context, userDto models.UserDto) (*models.User, error){
	user := &models.User{
		UserDto: userDto,
	}
	userQuery := myquery.User
	userQuery.UseDB(&common.MyApiServer.DB)
	err := userQuery.WithContext(ctx).Create(user)
	if err != nil{
		return nil, err
	}
	return user, nil
}

func UpdateUser(ctx context.Context, userDto models.UserDto, id uint64) (*models.User, error){
	user := &models.User{
		ID: id,
		UserDto: userDto,
	}
	userQuery := myquery.User
	userQuery.UseDB(&common.MyApiServer.DB)
	_, err := userQuery.WithContext(ctx).Where(myquery.User.ID.Eq(id)).Updates(user)
	if err != nil{
		return nil, err
	}
	return user, nil
}	

func DeleteUser(ctx context.Context, id uint64) error{
	// actual delete using unscoped
	// _, err := myquery.User.WithContext(ctx).Where(myquery.User.ID.Eq(id)).Unscoped().Delete()
	// fake delete only updated the deleted_at in table
	userQuery := myquery.User
	userQuery.UseDB(&common.MyApiServer.DB)
	_, err := userQuery.WithContext(ctx).Where(myquery.User.ID.Eq(id)).Delete()
	if err != nil{
		return err
	}
	return nil
}
// getting all data with fake deleted data
func GetAllUnscopedUsers(ctx context.Context) ([]*models.User, error){
	userQuery := myquery.User
	userQuery.UseDB(&common.MyApiServer.DB)
	allUsers, err := userQuery.WithContext(ctx).Unscoped().Find()
	if err != nil{
		return nil, err
	}
	return allUsers, nil
}