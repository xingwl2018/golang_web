package rule

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 核心控制逻辑

func ruleCreateOneHandler(ctx iris.Context) {
	var param PostRuleParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.ReadJSON(make_response.MakeResponse(http.StatusNotFound, pkg.ErrorBodyJson, true))
		return
	}
	if !param.notNull() {
		ctx.ReadJSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorBodyIsNull, true))
		return
	}
	if err := param.Valid(); err != nil {
		ctx.ReadJSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var rule v1.RuleForExchangeOrCoupon
	rule = v1.RuleForExchangeOrCoupon{
		Question: param.Question,
		Answer:   param.Answer,
		Type:     param.Type,
	}
	pkg.MyDatabase.InsertOne(&rule)
	ctx.JSON(make_response.MakeResponse(http.StatusOK, rule.Serializer(), false))
}

func rulePatchOneHandler(ctx iris.Context) {

	var param PostRuleParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	id, _ := ctx.Params().GetInt("rule_id")
	var rule v1.RuleForExchangeOrCoupon
	if ok, _ := pkg.MyDatabase.Where("id = ?", id).Get(&rule); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	session := pkg.MyDatabase.NewSession()
	session.Begin()
	if param.Answer != "" {
		rule.Answer = param.Answer
		if _, dbError := session.ID(rule.ID).Cols("answer").Update(&rule); dbError != nil {
			session.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.Question != "" {
		rule.Question = param.Question
		if _, dbError := session.ID(rule.ID).Cols("question").Update(&rule); dbError != nil {
			session.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	rule.Type = param.Type
	if _, dbError := session.ID(rule.ID).Cols("type").Update(&rule); dbError != nil {
		session.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	session.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, rule.Serializer(), false))
}

func ruleGetAllHandler(ctx iris.Context) {
	p := ctx.URLParamDefault("return", "all_list")
	var param GetRuleParam
	param.Return = p
	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var rules []v1.RuleForExchangeOrCoupon
	if dbError := pkg.MyDatabase.OrderBy("id").Desc("id").Find(&rules); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if param.Return == "all_list" {
		var results []v1.RuleForExchangeOrCouponSerializer
		for _, i := range rules {
			results = append(results, i.Serializer())
		}
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}
	if param.Return == "all_count" {
		var count = make(map[string]int)
		count["count"] = len(rules)
		ctx.JSON(make_response.MakeResponse(http.StatusOK, count, false))
		return
	}

}

func ruleGetOneHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("rule_id")

	var rule v1.RuleForExchangeOrCoupon
	if ok, _ := pkg.MyDatabase.ID(id).Get(&rule); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, rule.Serializer(), false))

}
