package mongo

import (
	// "context"
	// "time"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongClients map[string]*mongo.Client

func InitClient(name string, cli *mongo.Client) {
	if cli == nil {
		panic("set mongoClient err, mongoClient = nil")
	}
	if name == "" {
		panic("set mongoClient err, name = '' ")
	}

	mongClients[name] = cli
}

func GetClient(name string) *mongo.Client {

	return mongClients[name]
}

func init() {
	mongClients = map[string]*mongo.Client{}
}

//************** conn example ********************
// func conn_1() {
// 	// Replace the uri string with your MongoDB deployment's connection string.
// 	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"

// 	uri := "mongodb://gzy:gzy@49.232.190.114:27717/admin?w=majority"
// 	ctx := context.Background()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer func() {
// 	// 	if err = client.Disconnect(ctx); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()

// 	InitClient("name", client)
// }

// func conn_2() {
// 	AppName := "ggg"
// 	auth := &options.Credential{
// 		AuthSource:  "admin",
// 		Username:    "gzy",
// 		Password:    "gzy",
// 		PasswordSet: true,
// 	}
// 	maxPoolSize := uint64(10)
// 	minPoolSize := uint64(2)
// 	conTimeout := time.Second * time.Duration(10)
// 	opt := &options.ClientOptions{
// 		AppName:        &AppName,
// 		Auth:           auth,
// 		ConnectTimeout: &conTimeout,
// 		MaxPoolSize:    &maxPoolSize,
// 		MinPoolSize:    &minPoolSize,
// 		Hosts:          []string{"49.232.190.114:27717"},
// 	}
// 	ctx := context.TODO()
// 	client, err := mongo.NewClient(opt)
// 	if err != nil {
// 		panic(err)
// 	}
// 	client.Connect(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer func() {
// 	// 	if err = client.Disconnect(ctx); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()
// 	InitClient("name", client)
// }
