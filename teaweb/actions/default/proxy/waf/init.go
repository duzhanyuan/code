package waf

import (
	"github.com/TeaWeb/code/teaweb/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Prefix("/proxy/waf").
			Helper(new(helpers.UserMustAuth)).
			Helper(new(Helper)).
			Get("", new(IndexAction)).
			GetPost("/add", new(AddAction)).
			Post("/delete", new(DeleteAction)).
			Get("/detail", new(DetailAction)).
			GetPost("/update", new(UpdateAction)).
			Get("/rules", new(RulesAction)).
			GetPost("/group/add", new(GroupAddAction)).
			Get("/group", new(GroupAction)).
			Post("/group/delete", new(GroupDeleteAction)).
			Post("/group/on", new(GroupOnAction)).
			Post("/group/off", new(GroupOffAction)).
			Post("/group/move", new(GroupMoveAction)).
			GetPost("/group/rule/add", new(RuleAddAction)).
			Post("/group/rule/on", new(RuleOnAction)).
			Post("/group/rule/off", new(RuleOffAction)).
			Post("/group/rule/delete", new(RuleDeleteAction)).
			Post("/group/rule/move", new(RuleMoveAction)).
			GetPost("/group/rule/update", new(RuleUpdateAction)).
			EndAll()
	})
}