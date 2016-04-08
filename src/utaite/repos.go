package utaite

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Repository interface {
	New(dst interface{}) (*datastore.Key, error)
	FindByName(dst interface{}, name string) error
	FindByID(dst interface{}, id int64) error
	All(dst interface{}) error
}

func init() {
	var _ Repository = &MemberRepository{}
}

type MemberRepository struct {
	ctx context.Context
}

func (repo *MemberRepository) key(id int64) *datastore.Key {
	return datastore.NewKey(repo.ctx, "Member", "", id, nil)
}

func (repo *MemberRepository) New(dst interface{}) (*datastore.Key, error) {
	m, ok := dst.(*Member)
	if !ok {
		return nil, fmt.Errorf("type mismatch: %T for %T data", &Member{}, dst)
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	k := datastore.NewIncompleteKey(repo.ctx, "Member", nil)
	k, err := datastore.Put(repo.ctx, k, m)
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (repo *MemberRepository) FindByName(dst interface{}, name string) error {
	return errors.New("Cannot find by name")
}

func (repo *MemberRepository) FindByID(dst interface{}, id int64) error {
	m, ok := dst.(*Member)
	if !ok {
		return fmt.Errorf("type mismatch: %T for %T data", &Member{}, dst)
	}

	k := repo.key(id)
	err := datastore.Get(repo.ctx, k, m)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MemberRepository) All(dst interface{}) error {
	var ms *[]*Member
	ms, ok := dst.(*[]*Member)
	if !ok {
		return fmt.Errorf("type mismatch: %T for %T data", ms, dst)
	}

	q := datastore.NewQuery("Member").Order("-UpdatedAt")
	_, err := q.GetAll(repo.ctx, ms)
	if err != nil {
		return err
	}
	return nil
}
