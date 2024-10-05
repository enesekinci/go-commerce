package database

import (
	"go-commerce/app/models"
	"go-commerce/core/constant"
)

func DropTables() {
	err := DB.Migrator().DropTable(
		&models.Role{},
		&models.User{},

		&models.Category{},
		&models.Brand{},

		//entities
		&models.Tax{},
		&models.TaxClass{},
		&models.Setting{},
		&models.Media{},
		&models.Language{},
		&models.Currency{},
		&models.Country{},
		&models.Zone{},
		&models.PaymentMethod{},
		&models.ShippingMethod{},
		//entities

		&models.VariantType{},
		&models.Variant{},

		&models.Tag{},

		&models.Attribute{},
		&models.AttributeCategory{},

		&models.OptionType{},
		&models.Option{},

		&models.ProductCategory{},
		&models.ProductAttribute{},
		&models.ProductVariant{},
		&models.ProductTag{},
		&models.ProductMedia{},
		&models.Product{},
	)
	if err != nil {
		panic("failed to drop tables")
	}
}
func SeedData() {

	roles := []models.Role{
		{
			Name: "Super Admin",
			Slug: "super-admin",
		},
		{
			Name: "Admin",
			Slug: "admin",
		},
		{
			Name: "User",
			Slug: "user",
		},
	}
	//
	users := []models.User{
		{
			Name:     "Enes Ekinci",
			Email:    "enes.eknc.96@gmail.com",
			Phone:    "5369501299",
			Password: "password",
			RoleID:   constant.SuperAdmin,
			Status:   constant.ACTIVE,
		},
		{
			Name:     "John Doe",
			Email:    "john@gmail.com",
			Phone:    "5555555555",
			Password: "password",
			RoleID:   constant.User,
			Status:   constant.ACTIVE,
		},
		{
			Name:     "Alex Doe",
			Email:    "alex@gmail.com",
			Phone:    "5545545454",
			Password: "password",
			RoleID:   constant.Admin,
			Status:   constant.INACTIVE,
		},
	}

	optionTypes := []string{
		"Field",
		"Textarea",
		"Dropdown",
		"Checkbox",
		"Radio",
		"Multiselect",
		"Date",
		"DateTime",
		"Time",
	}

	brands := []models.Brand{
		{
			Name:               "Apple",
			Description:        "Apple Inc. is an American multinational technology company that specializes in consumer electronics, computer software, and online services.",
			SeoUrl:             "apple",
			SeoMetaTitle:       "Apple",
			SeoMetaDescription: "Apple Inc. is an American multinational technology company that specializes in consumer electronics, computer software, and online services.",
			Status:             constant.ACTIVE,
			Logo:               "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Apple_logo_black.svg/505px-Apple_logo_black.svg.png",
			Banner:             "https://www.apple.com/v/home/eb/images/heroes/iphone-12-pro/iphone_12_pro__fjx2x2b4jw6u_large.jpg",
		},
		{
			Name:        "Samsung",
			Description: "Samsung Electronics Co., Ltd. is a South Korean multinational electronics company headquartered in the Yeongtong District of Suwon.",
			SeoUrl:      "samsung",
			Status:      constant.ACTIVE,
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/0/0e/Samsung_Logo.svg",
			Banner:      "https://www.samsung.com/us/explore/galaxy-s21-5g/assets/images/galaxy-s21-5g_hero_1.jpg"},
		{
			Name:        "Huawei",
			Description: "Huawei Technologies Co., Ltd. is a Chinese multinational technology company headquartered in Shenzhen, Guangdong.",
			SeoUrl:      "huawei",
			Status:      constant.ACTIVE,
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/2/26/Huawei_logo.svg",
			Banner:      "https://consumer.huawei.com/content/dam/huawei-cbg-site/common/mkt/pdp/phones/p40-pro-plus/images/hero-banner/p40-pro-plus-hero-banner-green.png"},
		{
			Name:        "Xiaomi",
			Description: "Xiaomi Corporation is a Chinese multinational electronics company founded in April 2010 and headquartered in Beijing.",
			SeoUrl:      "xiaomi",
			Status:      constant.ACTIVE,
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/4/41/Xiaomi_logo.svg",
			Banner:      "https://cdn.cnn.com/cnnnext/dam/assets/200407120159-xiaomi-mi-10-pro-5g-super-pano.jpg"},
		{
			Name:        "Oppo",
			Description: "Oppo Electronics Corp., commonly referred to as Oppo, is a Chinese consumer electronics and mobile communications company headquartered in Dongguan, Guangdong.",
			SeoUrl:      "oppo",
			Status:      constant.ACTIVE,
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/0/0b/Oppo_logo.svg",
			Banner:      "https://www.oppo.com/content/dam/oppo/common/mkt/v2-2/reno4-pro-5g/reno4-pro-5g-hero-banner-black.png"},
		{
			Name:        "Vestel",
			Description: "Vestel Elektronik Sanayi ve Ticaret A.Ş. is a Turkish home and professional appliances manufacturing company consisting of 18 companies specialised in electronics, major appliances and information technology.",
			SeoUrl:      "vestel",
			Status:      constant.ACTIVE,
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/4/4b/Vestel_logo.svg",
			Banner:      "https://www.vestel.com.tr/assets/images/hero-banner/hero-banner-vestel-1.jpg"},
		{
			Name:        "LG",
			Description: "LG Corporation is a South Korean multinational conglomerate corporation.",
			SeoUrl:      "lg",
			Status:      constant.ACTIVE,
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/6/6f/LG_logo_%282015%29.svg",
			Banner:      "https://www.lg.com/us/images/TVs/hero/hero-4k-oled-tv-77-inch-oled77gxpua-2020.jpg"},
	}

	for _, role := range roles {
		DB.Create(&role)
	}

	for _, user := range users {
		DB.Create(&user)
	}

	for _, optionType := range optionTypes {
		DB.Create(&models.OptionType{Name: optionType})
	}

	for _, brand := range brands {
		DB.Create(&brand)
	}
}
