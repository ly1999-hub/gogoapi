package mongodb
import(
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	database ="gogo"
	db *mongo.Database
	client *mongo.Client
)

func Connect(){
	client,err :=mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil{
		println(err)
	}
	ctx,cencel := context.WithTimeout(context.Background(),20*time.Second)
	defer cencel()

	err = client.Connect(ctx)

	if err !=nil{
		println(err)
	}else{
		println("connect ok")
	}
	db =client.Database(database)
}

func GetDB()*mongo.Database{
	return db
}