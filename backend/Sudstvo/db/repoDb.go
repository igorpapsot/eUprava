package db

import (
	"Sudstvo/data"
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

func (u *RepoDb) GetOptuznica(id string) (data.Optuznica, error) {
	u.logger.Println("Getting optuznica...")
	var result data.Optuznica
	coll := u.getOptuzniceCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldnt find optuznica")
	}

	return result, nil
}

func (u RepoDb) CreatePoternica(p *data.Poternica) bool {
	u.logger.Println("Creating poternica...")
	coll := u.getPoterniceCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created poternica with _id: %v\n", result.InsertedID)
	return true
}

func (u RepoDb) GetPoternice() data.Poternice {
	u.logger.Println("Getting poternice...")
	coll := u.getPoterniceCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.Poternica
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

func (u *RepoDb) GetPoternica(id string) (data.Poternica, error) {
	u.logger.Println("Getting poternica...")
	var result data.Poternica
	coll := u.getPoterniceCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldn't find poternica")
	}

	return result, nil
}

//// ispraviti
//func (u RepoDb) CreateKonacnaPresuda(p *data.Sud) bool {
//	u.logger.Println("Creating konacna presuda ...")
//	coll := u.getKonacnaPresudaCollection()
//	id := uuid.New()
//	p.Id = id.String()
//	rand.Seed(time.Now().UnixNano())
//
//	user, err := p.ToBson()
//	result, err := coll.InsertOne(context.TODO(), user)
//	if err != nil {
//		u.logger.Println(err)
//		return false
//	}
//
//	u.logger.Printf("Created konacna presuda with _id: %v\n", result.InsertedID)
//	return true
//}

func (u RepoDb) CreateSud(p *data.Sud) bool {
	u.logger.Println("Creating sud ...")
	coll := u.getSudCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created sud with _id: %v\n", result.InsertedID)
	return true
}

func (u *RepoDb) GetSud(id string) (data.Sud, error) {
	u.logger.Println("Getting sud...")
	var result data.Sud
	coll := u.getSudCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldn't find sud")
	}

	return result, nil
}

func (u RepoDb) GetSudovi() data.Sudovi {
	u.logger.Println("Getting sudovi ...")
	coll := u.getSudCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.Sud
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

func (u RepoDb) CreateSudija(p *data.Sudija) bool {
	u.logger.Println("Creating sudija ...")
	coll := u.getSudijaCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created sudija with _id: %v\n", result.InsertedID)
	return true
}

func (u *RepoDb) GetSudija(jmbg string) (data.Sudija, error) {
	u.logger.Println("Getting sudija...")
	var result data.Sudija
	coll := u.getSudijaCollection()
	filter := bson.D{{"jmbg", jmbg}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldn't find sudija")
	}

	return result, nil
}

func (u RepoDb) GetSudije() data.Sudije {
	u.logger.Println("Getting sudije ...")
	coll := u.getSudijaCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.Sudija
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

func (u RepoDb) CreateRociste(p *data.Rociste) bool {
	u.logger.Println("Creating rociste...")
	coll := u.getRocisteCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Created rociste with _id: %v\n", result.InsertedID)
	return true
}

func (u RepoDb) GetRocista() data.Rocista {
	u.logger.Println("Getting rocista...")
	coll := u.getRocisteCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.Rociste
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

func (u *RepoDb) GetRociste(id string) (data.Rociste, error) {
	u.logger.Println("Getting rociste...")
	var result data.Rociste
	coll := u.getRocisteCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldn't find rociste")
	}

	return result, nil
}

func (u *RepoDb) Login(jmbg string, lozinka string) (data.Sudija, error) {
	u.logger.Println("Logging in...")
	var result data.Sudija
	coll := u.getSudijaCollection()
	filter := bson.D{{"jmbg", jmbg}, {"lozinka", lozinka}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		u.logger.Println(err)
		return result, errors.New("Couldn't log in")
	}

	return result, nil
}

func (u RepoDb) Register(p *data.Sudija) bool {
	u.logger.Println("Registering...")
	coll := u.getSudijaCollection()
	id := uuid.New()
	p.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Registered sudija with _id: %v\n", result.InsertedID)
	return true
}

//func (u *RepoDb) getKonacnaPresudaCollection() *mongo.Collection {
//	db := u.client.Database("myDB")
//	collection := db.Collection("konacna presuda")
//	return collection
//}

func (u *RepoDb) getRocisteCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("rociste")
	return collection
}

func (u *RepoDb) getPoterniceCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("poternice")
	return collection
}

func (u *RepoDb) getSudCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("sud")
	return collection
}

func (u *RepoDb) getOptuzniceCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("optuznice")
	return collection
}

func (u *RepoDb) getSudijaCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("sudije")
	return collection
}

// NoSQL: Constructor which reads db configuration from environment
func SudstvoRepoDB(ctx context.Context, logger *log.Logger) (*RepoDb, error) {

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
