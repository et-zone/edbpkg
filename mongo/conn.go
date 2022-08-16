package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"

	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	Cli     *mongo.Client
	Collect map[string]*MCollection
	mu      sync.Mutex
}

func (c *Client) Collection(keyName string) *MCollection {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Collect[keyName]
}

// keyName is unique name fo client
func (c *Client) AddCollection(keyName, DB, colleName string) error {
	if keyName == "" {
		return errors.New("keyName can not '' ")
	}
	if DB == "" {
		return errors.New("DB can not '' ")
	}
	if colleName == "" {
		return errors.New("colleName can not '' ")
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	db := c.Cli.Database(DB)
	mc := &MCollection{
		col: db.Collection(colleName),
	}
	c.Collect[keyName] = mc
	return nil
}

func New(ctx context.Context, uri string) (*Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &Client{
		client,
		map[string]*MCollection{},
		sync.Mutex{},
	}, err
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
