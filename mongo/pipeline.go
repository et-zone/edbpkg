package mongo

// "go.mongodb.org/mongo-driver/bson"

type Pipeline []map[string]interface{}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

//pipM use pipM Type
func (p *Pipeline) Group(pipM PipM) *Pipeline {
	(*p) = append((*p), PipM{"$group": pipM})
	return p
}

//matchFilter use Filter Type ,if neet match ,the func must executed first
func (p *Pipeline) Match(matchFilter Filter) *Pipeline {
	(*p) = append((*p), Filter{"$match": matchFilter})
	return p
}

func (p *Pipeline) Values() Pipeline {
	return *p
}

type PipM map[string]interface{}

func NewPipM() *PipM {
	return &PipM{}
}

// pip := []bson.M{{"$group": bson.M{"_id": "$age", "total": bson.M{"$sum": 1}}}}
//groupField 需要分组的字段，groupFieldName，给予命名
func (pm *PipM) GroupField(groupFieldName, groupField string) *PipM {
	(*pm)[groupFieldName] = "$" + groupField
	return pm
}

//val like 1,'$age' ,选择任意字段求和
func (pm *PipM) Sum(sumName string, val interface{}) *PipM {
	(*pm)[sumName] = PipM{"$sum": val}
	return pm
}

//val like 1,'$age' ,选择任意字段求平均值
func (pm *PipM) Avg(avgName string, val interface{}) *PipM {
	(*pm)[avgName] = PipM{"$avg": val}
	return pm
}

//val like 1,'$age' ,选择任意字段求最小值
func (pm *PipM) Min(minName string, val interface{}) *PipM {
	(*pm)[minName] = PipM{"$min": val}
	return pm
}

//val like 1,'$age' ,选择任意字段求最大值
func (pm *PipM) Max(sumName string, val interface{}) *PipM {
	(*pm)[sumName] = PipM{"$max": val}
	return pm
}

func (pm *PipM) Values() PipM {
	return *pm
}
