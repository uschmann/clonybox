package storage

func Set(key, value string) {
	if Has(key) {
		Db.Model(&Setting{}).Where("Key = ?", key).Update("Value", value)
	} else {
		Db.Create(&Setting{
			Key:   key,
			Value: value,
		})
	}
}

func Has(key string) bool {
	setting := Setting{}

	result := Db.Where("Key = ?", key).Find(&setting)

	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func Get(key string) (string, bool) {
	if !Has(key) {
		return "", false
	}

	setting := Setting{}

	Db.Where("Key = ?", key).First(&setting)

	return setting.Value, true
}
