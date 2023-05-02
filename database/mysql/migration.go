package mysql

import "gin/src/entity"

func (db *DB) RunMigration() error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Ecotourism{},
		&entity.Comment{},
		&entity.Trash{},
	); err != nil {
		return err
	}

	return nil
}
