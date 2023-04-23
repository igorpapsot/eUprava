package db

import (
	"Tuzilastvo/data"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"math/rand"
	"os"
	"time"
)

type RepoDb struct {
	logger *log.Logger
	client *mongo.Client
}

func (u RepoDb) GetPrijave() data.KrivicnePrijave {
	u.logger.Println("Getting krivicne prijave...")
	coll := u.getCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.KrivicnaPrijava
	if err = cursor.All(context.TODO(), &results); err != nil {
		u.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			u.logger.Println(err)
		}
		u.logger.Printf("%s\n", output)
	}
	return results
}

func (u RepoDb) CreatePrijava(p *data.KrivicnaPrijava) bool {
	u.logger.Println("Creating krivicna prijava...")
	coll := u.getCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created krivicna prijava with _id: %v\n", result.InsertedID)
	return true
}

func (u *RepoDb) ConfirmPrijava(prijava *data.KrivicnaPrijava) bool {
	//TODO implement me
	panic("implement me")
}

func (u *RepoDb) DeclinePrijava(prijava *data.KrivicnaPrijava) bool {
	//TODO implement me
	panic("implement me")
}

func (u *RepoDb) getCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("prijave")
	return collection
}

// NoSQL: Constructor which reads db configuration from environment
func NewRepoDB(ctx context.Context, logger *log.Logger) (*RepoDb, error) {

	uri := os.Getenv("MONGO_DB_URI")
	logger.Println(uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &RepoDb{
		client: client,
		logger: logger,
	}, nil
}

func (u *RepoDb) Disconnect(ctx context.Context) error {
	err := u.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *RepoDb) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := u.client.Ping(ctx, readpref.Primary())
	if err != nil {
		u.logger.Println(err)
	}

	// Print available databases
	databases, err := u.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		u.logger.Println(err)
	}
	fmt.Println(databases)
}
