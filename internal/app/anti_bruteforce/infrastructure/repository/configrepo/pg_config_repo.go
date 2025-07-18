package configrepo

import (
	"context"
	"fmt"

	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/domain/valueobject"
	"github.com/gkarman/anti_bruteforce/internal/config"
	"github.com/jmoiron/sqlx"
	// регистрирует драйвер PostgreSQL для database/sql.
	_ "github.com/lib/pq"
)

type PgConfigRepo struct {
	db *sqlx.DB
}

func NewPgConfigRepo(cfg config.DBRepo) (*PgConfigRepo, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	return &PgConfigRepo{db: db}, nil
}

func (s *PgConfigRepo) Close(_ context.Context) error {
	return s.db.Close()
}

func (s *PgConfigRepo) AddCidrInBlacklist(ctx context.Context, cidr valueobject.CIDR) error {
	query := `
		INSERT INTO blacklist (cidr)
		VALUES (:cidr)
		ON CONFLICT (cidr) DO NOTHING
	`

	params := map[string]interface{}{
		"cidr": cidr.String(),
	}

	_, err := s.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return fmt.Errorf("failed to insert cidr into blacklist: %w", err)
	}

	return nil
}
