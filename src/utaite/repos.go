package utaite

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type MemberRepository struct {
	ctx context.Context
}

func (repo *MemberRepository) key(id int64) *datastore.Key {
	return datastore.NewKey(repo.ctx, "Member", "", id, nil)
}

func (repo *MemberRepository) New(dst *Member) (*datastore.Key, error) {
	dst.CreatedAt = time.Now()
	dst.UpdatedAt = time.Now()

	key := datastore.NewIncompleteKey(repo.ctx, "Member", nil)
	key, err := datastore.Put(repo.ctx, key, dst)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (repo *MemberRepository) Find(dst *Member, id int64) error {
	key := repo.key(id)
	err := datastore.Get(repo.ctx, key, dst)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MemberRepository) All(dst *[]*Member) error {
	q := datastore.NewQuery("Member").Order("-UpdatedAt")
	_, err := q.GetAll(repo.ctx, dst)
	if err != nil {
		return err
	}
	return nil
}
