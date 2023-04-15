package repo

import "github.com/arifmr/movie-festival-app/pkg/entity"

type postgresql struct{ tc SQLTxConn }

type PostgreSQL interface {
	MoviesRepository
}

func NewPostgreSQL(txc SQLTxConn) PostgreSQL { return &postgresql{txc} }

const (
	_ entity.FlagIndex = iota

	ADMIN_FLAG_DISABLED      // is_active
	ADMIN_FLAG_ROLE_ROOT     // system
	ADMIN_FLAG_ROLE_ADMIN    // admin bprs
	ADMIN_FLAG_ROLE_BUSINESS // business
	ADMIN_FLAG_ROLE_CUSTOMER // customer
	ADMIN_FLAG_ROLE_FINANCE  // finance
	ADMIN_FLAG_PATH_XX       // path
)
