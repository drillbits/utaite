package utaite

import (
	"math"
	"strings"
	"testing"
	"time"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestMemberRepositoryNew(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	repo := MemberRepository{ctx: ctx}
	member := &Member{Name: "test user"}

	k, err := repo.New(member)
	if err != nil {
		t.Fatal(err)
	}

	actual := &Member{}
	err = datastore.Get(ctx, k, actual)
	if err != nil {
		t.Fatal(err)
	}
	if actual == nil {
		t.Fatalf("new member is nil")
	}
	if actual.Name != member.Name {
		t.Fatalf("expected: `%v`, but got: `%v`", member.Name, actual.Name)
	}

	// Type mismatch
	_, err = repo.New(&Role{})
	if err == nil {
		t.Fatal("error expected")
	}
	if !strings.Contains(err.Error(), "type mismatch:") {
		t.Fatalf("expected: `%v`, but got: `%v`", "type mismatch: ...", err)
	}
}

func TestMemberRepositoryFindByName(t *testing.T) {
	repo := MemberRepository{}
	err := repo.FindByName(&struct{}{}, "")
	if err.Error() != "Cannot find by name" {
		t.Fatalf("expected: `%v`, but got: `%v`", "Cannot find by name", err)
	}
}

func TestMemberRepositoryFindByID(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	repo := MemberRepository{ctx: ctx}
	member := &Member{Name: "test user"}
	k, err := repo.New(member)
	if err != nil {
		t.Fatal(err)
	}
	id := k.IntID()

	actual := &Member{}
	err = repo.FindByID(actual, id)
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatalf("member: %d is nil", id)
	}
	if actual.Name != member.Name {
		t.Fatalf("expected: `%v`, but got: `%v`", member.Name, actual.Name)
	}

	// Not found
	err = repo.FindByID(actual, math.MaxInt64)
	if err != datastore.ErrNoSuchEntity {
		t.Fatalf("member: %d exists", math.MaxInt64)
	}

	// Type mismatch
	err = repo.FindByID(&Role{}, id)
	if err == nil {
		t.Fatal("error expected")
	}
	if !strings.Contains(err.Error(), "type mismatch:") {
		t.Fatalf("expected: `%v`, but got: `%v`", "type mismatch: ...", err)
	}
}

func TestMemberRepositoryAll(t *testing.T) {
	ctx, done, err := NewContextWithOptions(&aetest.Options{
		StronglyConsistentDatastore: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	t1, err := time.Parse(time.RFC3339, "2014-05-23T12:06:00+09:00")
	if err != nil {
		t.Fatal(err)
	}
	_, err = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Member", nil), &Member{
		Name:      "test1",
		UpdatedAt: t1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse(time.RFC3339, "2014-05-23T12:05:00+09:00")
	if err != nil {
		t.Fatal(err)
	}
	_, err = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Member", nil), &Member{
		Name:      "test2",
		UpdatedAt: t2,
	})
	if err != nil {
		t.Fatal(err)
	}

	repo := MemberRepository{ctx: ctx}
	var members []*Member
	err = repo.All(&members)
	if err != nil {
		t.Fatal(err)
	}

	if len(members) != 2 {
		t.Fatalf("expected: `%v`, but got: `%v`", 2, len(members))
	}
	if members[0].Name != "test1" {
		t.Fatalf("expected: `%v`, but got: `%v`", "test1", members[0].Name)
	}
	if members[1].Name != "test2" {
		t.Fatalf("expected: `%v`, but got: `%v`", "test1", members[1].Name)
	}

	// Type mismatch
	err = repo.All(members)
	if err == nil {
		t.Fatal("error expected")
	}
	if !strings.Contains(err.Error(), "type mismatch:") {
		t.Fatalf("expected: `%v`, but got: `%v`", "type mismatch: ...", err)
	}
}
