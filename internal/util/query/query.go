package query

import (
	mongodb2 "myapp/module/mongodb"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myapp/internal/constant"
	"myapp/internal/util/parray"
	"myapp/internal/util/pstring"
	"myapp/internal/util/ptime"
)

// Query ...
type Query struct {
	Page           int64
	Limit          int64
	Active         string
	Keyword        string
	Status         string
	SortStr        string
	SortBSON       interface{}
	MultiStatus    string
	Supplier       string
	Inventory      string
	Inventories    string
	DeliveryStatus string
	User           string
	Province       string
	District       string
	Banned         string

	FromAt    time.Time
	ToAt      time.Time
	Timestamp time.Time
}

// AssignBanned ...
func (q Query) AssignBanned(cond bson.M) {
	b, err := strconv.ParseBool(q.Banned)
	if err == nil {
		cond["banned"] = b
	}
}

// AssignFromAtToAt ...
func (q Query) AssignFromAtToAt(cond bson.M) {
	if !q.FromAt.IsZero() && !q.ToAt.IsZero() {
		q.ToAt = ptime.TimeStartOfDayInHCM(q.ToAt.AddDate(0, 0, 1))
		q.FromAt = ptime.TimeStartOfDayInHCM(q.FromAt)
		(cond)["createdAt"] = bson.M{
			"$gte": q.FromAt,
			"$lt":  q.ToAt,
		}
	} else if !q.FromAt.IsZero() && q.ToAt.IsZero() {
		q.FromAt = ptime.TimeStartOfDayInHCM(q.FromAt)
		(cond)["createdAt"] = bson.M{
			"$gte": q.FromAt,
		}
	} else if q.FromAt.IsZero() && !q.ToAt.IsZero() {
		q.ToAt = ptime.TimeStartOfDayInHCM(q.ToAt.AddDate(0, 0, 1))
		(cond)["createdAt"] = bson.M{
			"$lt": q.ToAt,
		}
	}
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

// AssignUser ...
func (q Query) AssignUser(cond bson.M) {
	if q.User != "" {
		id, _ := primitive.ObjectIDFromHex(q.User)
		cond["user"] = id
	}
}

// AssignKeyword ...
func (q Query) AssignKeyword(cond bson.M) {
	if len(q.Keyword) >= constant.MinLenKeyword {
		cond["searchString"] = mongodb2.GenerateQuerySearchString(q.Keyword)
	}
}

// AssignDeliveryStatus ...
func (q Query) AssignDeliveryStatus(cond bson.M) {
	if q.DeliveryStatus != "" && q.DeliveryStatus != constant.QueryAll {
		cond["delivery.status"] = q.DeliveryStatus
	}
}

// AssignStatus ...
func (q Query) AssignStatus(cond bson.M) {
	if q.Status != "" && q.Status != constant.QueryAll {
		cond["status"] = q.Status
	}
}

// AssignMultiStatus ...
func (q Query) AssignMultiStatus(cond bson.M) {
	if q.MultiStatus != "" && q.MultiStatus != constant.QueryAll {
		cond["status"] = bson.M{
			"$in": pstring.Split(q.MultiStatus, ","),
		}
	}
}

// AssignOrderStatus ...
func (q Query) AssignOrderStatus(cond bson.M) {
	if q.Status != "" && q.Status != constant.QueryAll {
		(cond)["status"] = q.Status
	}
}

// AssignInventory ...
func (q Query) AssignInventory(cond bson.M) {
	if q.Inventory != "" {
		if inventoryID, isValid := mongodb2.NewIDFromString(q.Inventory); isValid {
			cond["inventory._id"] = inventoryID
		}
	}
}

// AssignInventories ...
func (q Query) AssignInventories(cond bson.M) {
	if q.Inventories != "" {
		inventories := parray.ConvertStringsToObjectIDs(pstring.Split(q.Inventories, ","))
		length := len(inventories)
		if length == 1 {
			cond["inventory"] = inventories[0]
		} else if length > 1 {
			cond["inventory"] = bson.M{
				"$in": inventories,
			}
		}
	}
}

// SetDefaultLimit ...
func (q *Query) SetDefaultLimit() {
	if q.Limit <= 0 {
		q.Limit = constant.Limit20
	}
}

// AssignSupplierMap ...
func (q *Query) AssignSupplierMap(m map[string]interface{}) {
	if q.Supplier != "" {
		supplierID, _ := mongodb2.NewIDFromString(q.Supplier)
		if !supplierID.IsZero() {
			m["supplier"] = supplierID
		}
	}
}

// AssignProvince ....
func (q *Query) AssignProvince(cond bson.M) {
	if q.Province != "" && q.Province != "all" {
		cond["province"] = q.Province
	}
}

// AssignDistrict ...
func (q *Query) AssignDistrict(cond bson.M) {
	if q.District != "" && q.District != "all" {
		cond["district"] = q.District
	}
}
