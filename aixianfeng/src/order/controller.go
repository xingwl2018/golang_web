package order

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"strconv"
)

// 核心控制逻辑

func getOneOrderHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("order_id")

	var order v1.Order
	if ok, _ := pkg.MyDatabase.ID(id).Get(&order); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	var account v1.Account
	if ok, _ := pkg.MyDatabase.ID(order.AccountId).Get(&account); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	order.Account = account
	ctx.JSON(make_response.MakeResponse(http.StatusOK, order.Serializer(), false))
}

func getAllOrderHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var (
		count   int64
		orders  []v1.Order
		dbError error
	)

	if count, dbError = pkg.MyDatabase.FindAndCount(&orders); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

	var resultsSerializer []v1.OrderSerializer
	for _, i := range orders {
		var account v1.Account
		pkg.MyDatabase.ID(i.AccountId).Get(&account)
		i.Account = account
		resultsSerializer = append(resultsSerializer, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, resultsSerializer, false))
}

func patchOrderHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("order_id")
	var param PatchOrderParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var order v1.Order
	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if ok, _ := tx.ID(id).Get(&order); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var account v1.Account
	tx.ID(order.AccountId).Get(&account)

	value := func(val string) int {
		for k, v := range v1.STATUS_MAP_EN {
			if val == v {
				return k
			}
		}
		return -1
	}
	order.Status = value(param.Status)
	order.Account = account
	if _, dbError := tx.ID(id).Cols("status").Update(&order); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, order.Serializer(), false))

}

func postOrderHandle(ctx iris.Context) {
	var param PostOrderParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var account v1.Account
	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if ok, _ := tx.ID(param.AccountId).Get(&account); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	var products []v1.Product
	if dbError := tx.In("id", param.ProductIds).Find(&products); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var total float64
	for _, i := range products {
		total += i.Price * i.Discount
	}
	t, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", total), 2)
	var order v1.Order
	order = v1.Order{
		ProductIds: param.ProductIds,
		Status:     0,
		AccountId:  int64(param.AccountId),
		Account:    account,
		Total:      t,
	}

	if _, dbError := tx.Insert(&order); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, order.Serializer(), false))

}
