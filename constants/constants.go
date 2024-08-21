package constants

import "time"

const (
	TestMode  = "test"
	DebugMode = "debug"

	JWTRefreshTokenExpireDuration = time.Hour * 24
	JWTAccessTokenExpireDuration  = time.Minute * 30
	ContextTimeoutDuration        = time.Second * 7

	CustomerRoleInSignup = "customer_in_signup"
	CustomerRole         = "customer"
	AdminRole            = "admin"

	UzLang = "uz"
	RuLang = "ru"

	VerifyCodeLength = 6

	CasbinConfigPath    = "configs/rbac_model.conf"
	MiddlewareRolesPath = "configs/models.csv"

	FirebaseReturnURL = "https://firebasestorage.googleapis.com/v0/b/phleybo.appspot.com/o/"

	Success = "success"
)

const (
	IncomeTransactionID = iota
	ExpenseTransactionID
	TransferTransactionID
)
