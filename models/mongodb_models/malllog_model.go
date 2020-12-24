package mongodb_models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type MallLog struct {
	ID *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Pid string
	Uid string
	UserAgent string
	CreateTime  time.Time
}

// CollectionName 设置test Collection
func (this *MallLog)CollectionName() string {
	return "mall_log"
}

// 单条新增
func (m *MallLog)InsertOne() (insertId interface{}, err error) {

	Collection := MognoDb.Collection(m.CollectionName())
	insertResult, err := Collection.InsertOne(context.TODO(), &m)
	if err != nil {
		log.Fatal(err)
	}

	insertId = insertResult.InsertedID
	return
}
/*
// 新增多条
func (m *MallLog)InsertMany() (insertResult *mongo.InsertManyResult, err error) {
	data := []interface{}{
		Test{
			Username:"周杰伦1",
			Password:"123455",
			UpdateTime:time.Now(),
			CreateTime:time.Now() ,
		},
		Test{
			Username:"周杰伦2",
			Password:"123455",
			UpdateTime:time.Now(),
			CreateTime:time.Now() ,
		},
		Test{
			Username:"周杰伦3",
			Password:"123455",
			UpdateTime:time.Now(),
			CreateTime:time.Now() ,
		},
	}
	Collection := MognoDb.Collection(m.CollectionName())
	insertResult, err = Collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// 查询单条
func (m *MallLog)FindOne() (t Test, err error) {
	Collection := MognoDb.Collection(m.CollectionName())
	result := Collection.FindOne(context.TODO(), bson.M{"username": "周杰伦"})

	err = result.Decode(&t);

	return
}

// 查询多条
func (m *MallLog)FindAll() (results []Test, err error) {

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(5)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	Collection := MognoDb.Collection(m.CollectionName())
	cur, err := Collection.Find(context.TODO(), bson.M{"username": bson.M{"$eq": "周杰伦"}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Test
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

// 更新单条
func (m *MallLog)UpdateOne() (result *mongo.UpdateResult, err error) {
	// Update one
	Collection := MognoDb.Collection(m.CollectionName())
	result, err = Collection.UpdateOne(
		context.TODO(), bson.M{"username": "周杰伦"},
		bson.M{"$set": bson.M{"password": "UserName_changed"}});
	return
}

// 更新多条
func (m *MallLog)UpdateMany() (result *mongo.UpdateResult, err error) {
	// Update many
	Collection := MognoDb.Collection(m.CollectionName())
	result, err = Collection.UpdateMany(
		context.TODO(), bson.M{"username": primitive.Regex{Pattern: "周杰", Options: ""}},
		bson.M{"$set": bson.M{"password": "dddddd"}});
	return
}

// 删除单条
func (m *MallLog)DeleteOne() (result *mongo.DeleteResult, err error) {
	Collection := MognoDb.Collection(m.CollectionName())
	result, err = Collection.DeleteOne(context.TODO(), bson.M{"username": "周杰"});
	return
}

// 删除多条
func (m *MallLog)DeleteMany() (result *mongo.DeleteResult, err error) {
	Collection := MognoDb.Collection(m.CollectionName())
	result, err = Collection.DeleteMany(context.TODO(), bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}});
	return
}*/

