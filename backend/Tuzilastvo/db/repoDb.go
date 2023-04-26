package db

import (
	"Tuzilastvo/data"
	"context"
	"encoding/json"
	"errors"
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

func (u RepoDb) GetPrijave(javne bool) data.KrivicnePrijave {
	u.logger.Println("Getting krivicne prijave...")
	coll := u.getPrijaveCollection()
	filter := bson.D{}
	if javne {
		filter = bson.D{{"privatnost", true}}
	}
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

func (u *RepoDb) GetPrijava(id string) (data.KrivicnaPrijava, error) {
	u.logger.Println("Getting prijava...")
	var result data.KrivicnaPrijava
	coll := u.getPrijaveCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldnt find prijava")
	}

	return result, nil
}

func (u RepoDb) CreatePrijava(p *data.KrivicnaPrijava) bool {
	u.logger.Println("Creating krivicna prijava...")
	coll := u.getPrijaveCollection()
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

func (u RepoDb) CreateTuzilastvo(p *data.Tuzilastvo) bool {
	u.logger.Println("Creating tuzilastvo ...")
	coll := u.getTuzilastvoCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created tuzilastvo with _id: %v\n", result.InsertedID)
	return true
}

func (u *RepoDb) GetTuzilastvo(id string) (data.Tuzilastvo, error) {
	u.logger.Println("Getting tuzilastva...")
	var result data.Tuzilastvo
	coll := u.getTuzilastvoCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldnt find tuzilastvo")
	}

	return result, nil
}

func (u RepoDb) GetTuzilastva() data.Tuzilastva {
	u.logger.Println("Getting tuzilastva ...")
	coll := u.getTuzilastvoCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.Tuzilastvo
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

func (u *RepoDb) ConfirmPrijava(id string) bool {
	coll := u.getPrijaveCollection()
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"status", data.PRIHVACENA}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		u.logger.Println(err)
		return false
	}
	return true
}

func (u *RepoDb) DeclinePrijava(id string) bool {
	coll := u.getPrijaveCollection()
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"status", data.ODBACENA}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		u.logger.Println(err)
		return false
	}
	return true
}

func (u RepoDb) GetOptuznice() data.Optuznice {
	u.logger.Println("Getting optuznice...")
	coll := u.getOptuzniceCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.Optuznica
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

func (u *RepoDb) GetOptuznica(prijavaId string) (data.Optuznica, error) {
	u.logger.Println("Getting optuznica...")
	var result data.Optuznica
	coll := u.getOptuzniceCollection()
	filter := bson.D{{"krivicnaPrijava.id", prijavaId}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldnt find optuznica")
	}

	return result, nil
}

func (u RepoDb) CreateOptuznica(p *data.Optuznica) bool {
	u.logger.Println("Creating optuznica...")
	coll := u.getPrijaveCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created optuznica with _id: %v\n", result.InsertedID)
	return true
}

func (u *RepoDb) getPrijaveCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("prijave")
	return collection
}

func (u *RepoDb) getTuzilastvoCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("tuzilastvo")
	return collection
}

func (u *RepoDb) getOptuzniceCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("optuznice")
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
