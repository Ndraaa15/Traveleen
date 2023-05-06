package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type TrashInterface interface {
	Exchange(ctx context.Context, trash entity.Trash) (entity.Trash, error)
	GetHistory(ctx context.Context, userID uint) ([]entity.Trash, error)
	GetByCode(ctx context.Context, code string) (entity.Trash, error)
	Update(ctx context.Context, trash entity.Trash) (entity.Trash, error)
}

type Trash struct {
	sql      mysql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitTrash(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) TrashInterface {
	return &Trash{
		sql:      sql,
		supabase: supabase,
	}
}

func (r *Trash) Exchange(ctx context.Context, trash entity.Trash) (entity.Trash, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&trash).Error; err != nil {
		return trash, err
	}
	return trash, nil
}

func (r *Trash) GetHistory(ctx context.Context, userID uint) ([]entity.Trash, error) {
	var trashes []entity.Trash
	if err := r.sql.Debug().WithContext(ctx).Where("user_id = ?", userID).Find(&trashes).Error; err != nil {
		return trashes, err
	}

	return trashes, nil
}

func (r *Trash) GetByCode(ctx context.Context, code string) (entity.Trash, error) {
	var trash entity.Trash
	if err := r.sql.Debug().WithContext(ctx).Where("code = ?", code).First(&trash).Error; err != nil {
		return trash, err
	}

	return trash, nil
}

func (r *Trash) Update(ctx context.Context, trash entity.Trash) (entity.Trash, error) {
	if err := r.sql.Debug().WithContext(ctx).Model(&trash).Updates(trash).Error; err != nil {
		return trash, err
	}
	return trash, nil
}
