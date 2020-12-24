package mysql_models

type Areas struct {
	Id           int    `gorm:"column(id);auto"`
	ProvinceCode int    `gorm:"column(province_code)"`
	ProvinceName string `gorm:"column(province_name);size(20)"`
	CityCode     int    `gorm:"column(city_code)"`
	CityName     string `gorm:"column(city_name);size(20)"`
	DistrictCode int    `gorm:"column(district_code)"`
	DistrictName string `gorm:"column(district_name);size(20)"`
}

func (t *Areas) TableName() string {
	return "gin_areas"
}

// 获取当前收货地址
func GetAreasInfoById(id int, userId int64) (*Areas, error) {
	var areas Areas
	address, err := GetAddressInfoById(id, userId)
	err = db.Where("id= ?", address.AreasId).
		First(&areas).Error

	return &areas, err
}
