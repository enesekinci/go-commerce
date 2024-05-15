package database

func IsExistInDB(table string, column string, value string) bool {

	var isExist int64

	DB.Table(table).Where(column+" = ?", value).Limit(1).Count(&isExist)

	return isExist > 0
}
