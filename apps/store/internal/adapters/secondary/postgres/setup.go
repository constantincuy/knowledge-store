package postgres

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/ports"
)

func RunSetup(ctx context.Context, provider ports.DatabaseProvider) error {
	global, err := provider.GetDatabase("")
	defer global.Close()

	if err != nil {
		return err
	}

	// We ignore the error since we don't care if the DB already exists
	_, _ = global.ExecContext(ctx, `CREATE DATABASE kb_system`)

	sys, err := provider.GetDatabase("kb_system")
	defer sys.Close()

	if err != nil {
		return err
	}

	_, err = sys.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS knowledge_bases(name_id varchar(255) PRIMARY KEY)")

	if err != nil {
		return err
	}

	return nil
}
