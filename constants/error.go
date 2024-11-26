package constants

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	// ErrNotFound ...
	ErrNotFound = Sentinel("not found")
	// ErrXozmakConstantNotExists ...
	ErrXozmakConstantNotExists = Sentinel("no company constant exists")
	//ErrXozmakAlreadyExists ...
	ErrXozmakAlreadyExists = Sentinel("the xozmak already exists")
	// ErrProductConstantNotExists ...
	ErrProductConstantNotExists = Sentinel("no product constant exists")
	//ErrProductAlreadyExists ...
	ErrProductAlreadyExists = Sentinel("the product already exists")
	// ErrRowsAffectedIsZero indicates that sql command didn't work
	ErrRowsAffectedIsZero = Sentinel("no rows affected after sql command")
	//ErrTransaction ...
	ErrTransaction = Sentinel("transaction start error")
)

const (
	// PGForeignKeyViolationCode is used to check foriegn key violation in database
	PGForeignKeyViolationCode = "23503"
	// PGUniqueKeyViolationCode is used to check unique key violation in database
	PGUniqueKeyViolationCode = "23505"
)
