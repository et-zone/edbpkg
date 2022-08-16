package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	LIMIT                              = int64(1000)
	UPSERT                             = true
	RER_NEW_DOC options.ReturnDocument = options.After
)

type M bson.M

type MCollection struct {
	col *mongo.Collection
}

//func NewMCollection(client *mongo.Client, DBName string, CollectionName string) *MCollection {
//	if client == nil {
//		//return nil,errors.New("NewMCollection err , init client is nil")
//		return nil
//	}
//	db := client.Database(DBName)
//	mc := &MCollection{
//		col: db.Collection(CollectionName),
//	}
//	return mc
//}

func (mc *MCollection) InsertOne(ctx context.Context, document interface{}) (bool, error) {
	_, err := mc.clone().InsertOne(ctx, document)
	if err != nil {
		return false, err
	}
	return true, err
}

//not Transaction 不支持事务特性
func (mc *MCollection) InsertAll(ctx context.Context, documents ...interface{}) (bool, error) {
	_, err := mc.clone().InsertMany(ctx, documents)
	if err != nil {
		return false, err
	}
	return true, err
}

func (mc *MCollection) FindOne(ctx context.Context, filter Filter, result interface{}) error {
	return mc.clone().FindOne(ctx, filter).Decode(result)
}

func (mc *MCollection) FindByID(ctx context.Context, id string, result interface{}) error {
	return mc.clone().FindOne(ctx, bson.M{"_id": id}).Decode(result)
}

//results is addr,default limit=1000
func (mc *MCollection) FindAll(ctx context.Context, filter Filter, results interface{}) error {
	cursor, err := mc.clone().Find(ctx, filter, &options.FindOptions{Limit: &LIMIT})
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

/*
- results is addr
- sortOpt=nil ==>no sort;
- skipOpt=0   ==>no skip
- sortOpt=0   ==>no sort ;sortOpt like bson.M{"age": 1}
**/
func (mc *MCollection) FindAllWithOptions(ctx context.Context, filter Filter, results interface{}, sortOpt interface{}, limitOpt int64, skipOpt int64) error {
	s := &options.FindOptions{
		Limit: &LIMIT,
	}
	if sortOpt != nil {
		s.Sort = sortOpt
	}

	if limitOpt > 0 {
		s.Limit = &limitOpt
	}
	if skipOpt > 0 {
		s.Skip = &skipOpt
	}
	cursor, err := mc.clone().Find(ctx, filter, s)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

func (mc *MCollection) UpdateByID(ctx context.Context, id string, update interface{}) (int64, error) {
	//UpdateOptions ==>Upsert, If true, a new document will be inserted if the filter does not match any documents in the collection
	r, err := mc.clone().UpdateByID(ctx, id, update, &options.UpdateOptions{Upsert: &UPSERT})
	if err != nil {
		return 0, err
	}
	return r.ModifiedCount, err
}

func (mc *MCollection) UpdateByCondition(ctx context.Context, filter Filter, update interface{}) (int64, error) {
	//UpdateOptions ==>Upsert, If true, a new document will be inserted if the filter does not match any documents in the collection
	r, err := mc.clone().UpdateMany(ctx, filter, update, &options.UpdateOptions{Upsert: &UPSERT})
	if err != nil {
		return 0, err
	}
	return r.ModifiedCount, err
}

//Inc
func (mc *MCollection) updateInc(ctx context.Context, filter Filter, update interface{}) (int64, error) {
	//UpdateOptions ==>Upsert, If true, a new document will be inserted if the filter does not match any documents in the collection
	r, err := mc.clone().UpdateMany(ctx, filter, update, &options.UpdateOptions{Upsert: &UPSERT})
	if err != nil {
		return 0, err
	}
	b := []byte{}
	r.UnmarshalBSON(b)

	fmt.Println(string(b))
	return r.ModifiedCount, err
}

// return new val
func (mc *MCollection) IncByID(ctx context.Context, id string, update interface{}, result interface{}) error {
	return mc.col.FindOneAndUpdate(ctx, Filter{"_id": id}, update, &options.FindOneAndUpdateOptions{ReturnDocument: &RER_NEW_DOC}).Decode(result)
}

// return new val ;only update one collection
func (mc *MCollection) IncOneByCondition(ctx context.Context, filter Filter, update interface{}, result interface{}) error {
	err := mc.col.FindOneAndUpdate(ctx, filter, update, &options.FindOneAndUpdateOptions{ReturnDocument: &RER_NEW_DOC}).Decode(result)
	return err
}

func (mc *MCollection) DeleteByID(ctx context.Context, id string) (int64, error) {
	r, err := mc.clone().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, err
	}
	return r.DeletedCount, err
}

func (mc *MCollection) DeleteByIDs(ctx context.Context, ids []string) (int64, error) {
	r, err := mc.clone().DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return 0, err
	}
	return r.DeletedCount, err
}

func (mc *MCollection) DeleteByCondition(ctx context.Context, filter Filter) (int64, error) {
	r, err := mc.clone().DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return r.DeletedCount, err
}

func (mc *MCollection) GetCount(ctx context.Context, filter interface{}) (int64, error) {

	count, err := mc.clone().CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, err
}

/*
[
  { $match : { score : { $gt : 70, $lte : 90 } } },
  { $group: { _id: null, count: { $sum: 1 } } }
]
**/
func (mc *MCollection) Aggregate(ctx context.Context, pipeline interface{}, results interface{}) error {
	cursor, err := mc.clone().Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

//use base MongoCollection
func (mc *MCollection) clone() *mongo.Collection {
	return mc.col
}
