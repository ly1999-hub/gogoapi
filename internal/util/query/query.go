package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myapp/internal/constant"
	mongodb2 "myapp/internal/module/mongodb"
	"strconv"
)

// Query ...
type Query struct {
	Page     int64
	Limit    int64
	Active   string
	Keyword  string
	SortBSON interface{}
}

// GetFindOptionsUsingPage ...
func (q Query) GetFindOptionsUsingPage() *options.FindOptions {
	opts := options.Find()

	if q.Limit > 0 {
		opts.SetLimit(q.Limit).SetSkip(q.Limit * q.Page)
	}

	if q.SortBSON != nil {
		opts.SetSort(q.SortBSON)
	}
	return opts
}

// AssignActive ...
func (q Query) AssignActive(cond bson.M) {
	b, err := strconv.ParseBool(q.Active)

	if err == nil {
		cond["active"] = b
	}
}

// AssignKeyword ...
func (q Query) AssignKeyword(cond bson.M) {
	if len(q.Keyword) >= constant.MinLenKeyword {
		cond["searchString"] = mongodb2.GenerateQuerySearchString(q.Keyword)
	}
}

// SetDefaultLimit ...
func (q *Query) SetDefaultLimit() {
	if q.Limit <= 0 {
		q.Limit = constant.Limit20
	}
}
