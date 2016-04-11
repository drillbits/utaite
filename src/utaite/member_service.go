package utaite

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type MemberService struct{}

type MemberListResp struct {
	List []*Member `json:"list"`
}

func (s *MemberService) Get(ctx context.Context, r *http.Request, req *IntIDRequest) (*Member, error) {
	appengine.AppID(appengine.NewContext(r))

	repo := MemberRepository{ctx: ctx}
	member := &Member{}
	err := repo.FindByID(member, req.ID)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, &HttpError{http.StatusNotFound, fmt.Sprintf("Member: %d not found", req.ID)}
		}
		return nil, err
	}

	return member, nil
}

func (s *MemberService) List(ctx context.Context, r *http.Request, opts *ListOpts) (*MemberListResp, error) {
	appengine.AppID(appengine.NewContext(r))

	repo := MemberRepository{ctx: ctx}
	members := make([]*Member, 0)
	err := repo.All(&members)
	if err != nil {
		return nil, err
	}

	return &MemberListResp{List: members}, nil
}
