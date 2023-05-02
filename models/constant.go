package models

const (
	APP_PORT string = "11010"

	DB_HOST string = "localhost"
	DB_PORT string = "5432"
	DB_USER string = "postgres"
	DB_PASS string = ""
	DB_NAME string = "job-portal"

	REDIS_CACHE_URL string = "localhost:6379"
)

type CachePrefix string

const (
	EnvKeySettingFeatureCacheGetListDresscodeCompanyEnabled         string = "SETTING_FEATURE_CACHE_GET_LIST_DRESSCODE_COMPANY_ENABLE"
	EnvKeySettingFeatureCacheGetListDresscodeCompanyDurationMinutes string = "SETTING_FEATURE_CACHE_GET_LIST_DRESSCODE_COMPANY_DURATION_MINUTES"
)

const (
	PrefixCompanyListDresscode CachePrefix = "company:companyDresscode:"
)
