package db

import (
	"GranicnaPolicija/data"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type GPRepoDb struct {
	logger *log.Logger
	client *mongo.Client
}

//======DataBase Connection==========

func NewGPRepoDB(ctx context.Context, logger *log.Logger) (*GPRepoDb, error) {

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

	return &GPRepoDb{
		client: client,
		logger: logger,
	}, nil
}

func (gr *GPRepoDb) Disconnect(ctx context.Context) error {
	err := gr.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (gr *GPRepoDb) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := gr.client.Ping(ctx, readpref.Primary())
	if err != nil {
		gr.logger.Println(err)
	}

	// Print available databases
	databases, err := gr.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		gr.logger.Println(err)
	}
	fmt.Println(databases)
}

//======Repo==========

func (gr GPRepoDb) Login(jmbg string, password string) (data.GPolicajac, error) {
	gr.logger.Println("Logging in . . .")
	var result data.GPolicajac

	coll := gr.getGPCollection()

	filter := bson.D{{"jmbg", jmbg}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	//Checking password
	resultBool := gr.CheckPasswordHash(password, result.Password)
	if resultBool == false {
		return result, errors.New("wrong username or password")
	}

	if err != nil {
		gr.logger.Println(err)
		return result, errors.New("wrong username or password")
	}

	return result, nil
}

func (gr GPRepoDb) NewGPolicajac(gpolicajac *data.GPolicajac) bool {
	gr.logger.Println("Registering granicni policajac...")
	coll := gr.getGPCollection()
	id := uuid.New()
	gpolicajac.Id = id.String()
	hashedPass, err := gpolicajac.HashPassword(gpolicajac.Password)
	gpolicajac.Password = hashedPass
	rand.Seed(time.Now().UnixNano())
	cCode := rand.Intn(999999-100001) + 100000
	gpolicajac.CCode = cCode

	user, err := gpolicajac.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		gr.logger.Println(err)
		return false
	}

	gr.logger.Printf("Registered user with _id: %v\n", result.InsertedID)
	return true
}

func (gr GPRepoDb) GetPolicajac(id string) (data.GPolicajac, error) {
	gr.logger.Println("Getting policajac...")
	var result data.GPolicajac
	coll := gr.getGPCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		gr.logger.Println(err)
		return result, errors.New("Couldnt find policajca")
	}

	return result, nil
}

//======Provera gradjanina

func (gr GPRepoDb) CreateProvera(provera *data.ProveraGradjanina) bool {
	gr.logger.Println("Creating provera...")
	coll := gr.getProveraCollection()
	id := uuid.New()
	provera.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := provera.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		gr.logger.Println(err)
		return false
	}

	gr.logger.Printf("Created prelazak with _id: %v\n", result.InsertedID)
	return true
}

func (gr GPRepoDb) GetProvera(gradjanin *data.Gradjanin) (data.ProveraGradjanina, error) {
	gr.logger.Println("Getting provera...")
	var result data.ProveraGradjanina
	coll := gr.getProveraCollection()
	filter := bson.D{{"gradjanin", gradjanin}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		gr.logger.Println(err)
		return result, errors.New("Couldnt find provera")
	}

	return result, nil
}

func (gr GPRepoDb) GetProvere() data.ProvereG {
	gr.logger.Println("Getting provere gradjanina...")
	coll := gr.getProveraCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		gr.logger.Println(err)
	}

	var results []*data.ProveraGradjanina
	if err = cursor.All(context.TODO(), &results); err != nil {
		gr.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			gr.logger.Println(err)
		}
		gr.logger.Printf("%s\n", output)
	}
	return results
}

func (gr GPRepoDb) GetProvereByStatus(status string) data.ProvereG {
	gr.logger.Println("Getting provere gradjanina...")
	coll := gr.getProveraCollection()
	filter := bson.D{{"status", status}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		gr.logger.Println(err)
	}

	var results []*data.ProveraGradjanina
	if err = cursor.All(context.TODO(), &results); err != nil {
		gr.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			gr.logger.Println(err)
		}
		gr.logger.Printf("%s\n", output)
	}
	return results
}

//======Prelazak Granice

func (gr GPRepoDb) CreatePrelazak(prelazak *data.PrelazakGranice) bool {
	gr.logger.Println("Creating prelazak...")
	coll := gr.getPrelazakCollection()
	id := uuid.New()
	prelazak.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := prelazak.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		gr.logger.Println(err)
		return false
	}

	gr.logger.Printf("Created prelazak with _id: %v\n", result.InsertedID)
	return true
}

func (gr GPRepoDb) GetPrelasci() data.PrelasciGranice {
	gr.logger.Println("Getting prelasci granice...")
	coll := gr.getPrelazakCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		gr.logger.Println(err)
	}

	var results []*data.PrelazakGranice
	if err = cursor.All(context.TODO(), &results); err != nil {
		gr.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			gr.logger.Println(err)
		}
		gr.logger.Printf("%s\n", output)
	}
	return results
}

func (gr GPRepoDb) GetPrelasciByPrelaz(prelaz string) data.PrelasciGranice {
	gr.logger.Println("Getting prelasci granice...")
	coll := gr.getPrelazakCollection()
	filter := bson.D{{"g_prelaz", prelaz}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		gr.logger.Println(err)
	}

	var results []*data.PrelazakGranice
	if err = cursor.All(context.TODO(), &results); err != nil {
		gr.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			gr.logger.Println(err)
		}
		gr.logger.Printf("%s\n", output)
	}
	return results
}

func (gr GPRepoDb) GetPrelazak(id string) (data.PrelazakGranice, error) {
	gr.logger.Println("Getting prelazak...")
	var result data.PrelazakGranice
	coll := gr.getPrelazakCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		gr.logger.Println(err)
		return result, errors.New("Couldnt find prelazak")
	}

	return result, nil
}

//======Krivicna prijava

func (gr GPRepoDb) CreatePrijava(prijava *data.KrivicnaPrijava) bool {
	gr.logger.Println("Creating prijava...")
	coll := gr.getPrijavaCollection()
	id := uuid.New()
	prijava.Id = id.String()
	rand.Seed(time.Now().UnixNano())

	user, err := prijava.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		gr.logger.Println(err)
		return false
	}

	gr.logger.Printf("Created prijava with _id: %v\n", result.InsertedID)
	return true
}

func (gr GPRepoDb) GetPrijave() data.KrivicnePrijave {
	gr.logger.Println("Getting krivicne prijave...")
	coll := gr.getPrijavaCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		gr.logger.Println(err)
	}

	var results []*data.KrivicnaPrijava
	if err = cursor.All(context.TODO(), &results); err != nil {
		gr.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			gr.logger.Println(err)
		}
		gr.logger.Printf("%s\n", output)
	}
	return results
}

func (gr GPRepoDb) GetPrijava(id string) (data.KrivicnaPrijava, error) {
	gr.logger.Println("Getting prijava...")
	var result data.KrivicnaPrijava
	coll := gr.getPrijavaCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		gr.logger.Println(err)
		return result, errors.New("Couldnt find prijava")
	}

	return result, nil
}

//======Kolekcije==========

func (gr *GPRepoDb) getGPCollection() *mongo.Collection {
	db := gr.client.Database("myGpDB")
	collection := db.Collection("policajac")
	return collection
}

func (gr *GPRepoDb) getProveraCollection() *mongo.Collection {
	db := gr.client.Database("myGpDB")
	collection := db.Collection("provere")
	return collection
}

func (gr *GPRepoDb) getPrelazakCollection() *mongo.Collection {
	db := gr.client.Database("myGpDB")
	collection := db.Collection("prelasci")
	return collection
}

func (gr *GPRepoDb) getPrijavaCollection() *mongo.Collection {
	db := gr.client.Database("myGpDB")
	collection := db.Collection("prijave")
	return collection
}

func (gr GPRepoDb) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
