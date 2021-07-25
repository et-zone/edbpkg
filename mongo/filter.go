package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SORT_UP   = 1
	SORT_DOWN = -1
)

type Filter map[string]interface{}

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) Values() Filter {
	return *f
}

func (f *Filter) Equal(field string, val interface{}) *Filter {
	(*f)[field] = val
	return f
}

//sortFlag choose SORT_UP or SORT_DOWN
// func (f *Filter) Sort(field string, SORT_FLAG int) *Filter {
// 	switch SORT_FLAG {
// 	case SORT_UP:
// 		(*f)["$sort"] = bson.M{field: SORT_FLAG}
// 	case SORT_DOWN:
// 		(*f)["$sort"] = bson.M{field: SORT_FLAG}
// 	default:
// 	}

// 	return f
// }

func (f *Filter) In(field string, val []interface{}) *Filter {
	(*f)[field] = bson.M{"$in": val}
	return f
}

func (f *Filter) NotIn(field string, vals []interface{}) *Filter {
	(*f)[field] = bson.M{"$nin": vals}
	return f
}

//val is int or float
func (f *Filter) LessThan(field string, val interface{}) *Filter {
	(*f)[field] = bson.M{"$lt": val}
	return f
}

//val is int or float
func (f *Filter) LessEqualThan(field string, val interface{}) *Filter {
	(*f)[field] = bson.M{"$lte": val}
	return f
}

//val is int or float
func (f *Filter) GreatThan(field string, val interface{}) *Filter {
	(*f)[field] = bson.M{"$gt": val}
	return f
}

//val is int or float ; like age >=10
func (f *Filter) GreatEqualThan(field string, val interface{}) *Filter {
	(*f)[field] = bson.M{"$gte": val}
	return f
}

func (f *Filter) Like(field string, val string) *Filter {
	(*f)[field] = bson.M{"$regex": val}
	return f
}

//like a???
func (f *Filter) LikeFront(field string, val string) *Filter {
	(*f)[field] = bson.M{"$regex": "^" + val}
	return f
}

//like ???a
func (f *Filter) LikeBack(field string, val string) *Filter {
	(*f)[field] = bson.M{"$regex": val + "$"}
	return f
}

//
func (f *Filter) OR(filters ...Filter) *Filter {
	conds := []Filter{}
	for _, filter := range filters {
		conds = append(conds, filter)
	}
	(*f)["$or"] = conds

	return f
}

func (f *Filter) AND(filters ...Filter) *Filter {
	conds := []Filter{}
	for _, filter := range filters {
		conds = append(conds, filter)
	}
	(*f)["$and"] = conds

	return f
}

//************************ update *****************************

type Update map[string]interface{}

func NewUpdate() *Update {
	return &Update{}
}

func (u *Update) Set(val bson.M) *Update {
	(*u)["$set"] = val

	return u
}

// delete field
// no exception if field not exits
func (u *Update) UnSet(field ...string) *Update {
	m := bson.M{}
	for _, fiel := range field {
		m[fiel] = ""
	}
	(*u)["$unset"] = m

	return u
}

//updateNowTime
func (u *Update) CurrentDateUpdate(field string) *Update {
	(*u)["$currentDate"] = bson.M{field: true}
	return u
}

/*
 ++ or --
 cond like { quantity: -2, "metrics.orders": 1 }
 key-- {quantity: -2}
*/
func (u *Update) Inc(cond bson.M) *Update {
	(*u)["$inc"] = cond
	return u
}
func (u *Update) Values() Update {
	return *u
}
