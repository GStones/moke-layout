package db_nosql

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-layout/internal/services/demo/db_nosql/demo"
)

type Database struct {
	logger *zap.Logger
	coll   diface.ICollection
}

func OpenDatabase(l *zap.Logger, coll diface.ICollection) Database {
	return Database{
		logger: l,
		coll:   coll,
	}
}

func (db *Database) LoadOrCreateDemo(ctx context.Context, id string) (*demo.Dao, error) {
	if dm, err := demo.NewDemoModel(ctx, id, db.coll); err != nil {
		return nil, err
	} else if err = dm.Load(); errors.Is(err, nerrors.ErrNotFound) {
		if dm, err = demo.NewDemoModel(ctx, id, db.coll); err != nil {
			return nil, err
		} else if err := dm.InitDefault(); err != nil {
			return nil, err
		} else if err = dm.Create(); err != nil {
			if err = dm.Load(); err != nil {
				return nil, err
			} else {
				return dm, nil
			}
		} else {
			return dm, nil
		}
	} else if err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}
