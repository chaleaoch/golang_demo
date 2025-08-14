package repo

import "chaleaoch.com/golang_demo/internal/entity"

type mIPRepo struct {
}

func (r *mIPRepo) GetByType(mType string) ([]*entity.MIp, error) {
	return nil, nil
}

func (r *mIPRepo) GetMip() (*entity.MIp, error) {
	return nil, nil
}

func NewMipRepo() *mIPRepo {
	return &mIPRepo{}
}
