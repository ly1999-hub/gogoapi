package service
import(
	"context"
	"myapp/internal/dao"
	"myapp/internal/constant"
	"go.mongodb.org/mongo-driver/bson"
)
type Permission struct{}

func (s Permission) Migration(ctx context.Context)error{
	var d= dao.Configuration{}

	update :=bson.M{
		"$set":bson.M{
			"permissions":constant.Permission{},
		},
	}
	return d.UpdateOne(ctx,bson.M{},update)
}