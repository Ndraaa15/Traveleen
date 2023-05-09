package mysql

import "gin/src/entity"

func (db *DB) RunMigration() error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Ecotourism{},
		&entity.Article{},
		&entity.Comment{},
		&entity.Trash{},
		&entity.Cart{},
		&entity.CartProduct{},
		&entity.Purchase{},
	); err != nil {
		return err
	}

	return nil
}
