package app

import (
	"utaite"

	"github.com/favclip/ucon"
	"github.com/favclip/ucon/swagger"
)

func setupAPI(swPlugin *swagger.Plugin) {
	var hInfo *swagger.HandlerInfo

	{
		// Member
		s := &utaite.MemberService{}
		tag := swPlugin.AddTag(&swagger.Tag{
			Name:        "Member",
			Description: "utaite Member",
		})

		hInfo = swagger.NewHandlerInfo(s.Get)
		ucon.Handle("GET", "/api/member/{id}", hInfo)
		hInfo.Description, hInfo.Tags = "get member", []string{tag.Name}

		hInfo = swagger.NewHandlerInfo(s.List)
		ucon.Handle("GET", "/api/member", hInfo)
		hInfo.Description, hInfo.Tags = "get member list", []string{tag.Name}
	}
}
