package users

import (
    "context"
    "time"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// New initializes a new MongoDB client and retuns a UserInterface.
func NewClient(mongoURI string) UserInterface {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    return NewMaldonClient(client)
}

// DiscordUser is the struct for the Maldon Database User table.
type DiscordUser struct {
    ID              string `bson:"_id,omitempty"`
    DiscordUsername string `bson:"discordUsername"`
    DiscordDiscrim  string `bson:"discordDiscrim"`
}

// UserInterface is the interface for the Maldon Database User table Actions.
type UserInterface interface {
    CreateUser(ctx context.Context, user DiscordUser) (*mongo.InsertOneResult, error)
    GetUser(ctx context.Context, discordUsername, discordDiscrim string) (DiscordUser, error) // Corrected method name
    SearchUsers(ctx context.Context) ([]DiscordUser, error)
}

// MaldonClient wraps a MongoDB client and implements the UserInterface.
type MaldonClient struct {
    client *mongo.Client
}

// NewMaldonClient creates a new MaldonClient.
func NewMaldonClient(client *mongo.Client) *MaldonClient {
    return &MaldonClient{client: client}
}

// CreateUser creates a new user in the Maldon Database.
func (mc *MaldonClient) CreateUser(ctx context.Context, user DiscordUser) (*mongo.InsertOneResult, error) {
    collection := mc.client.Database("Maldon").Collection("users")
    return collection.InsertOne(ctx, user)
}

// GetUser gets a user from the Maldon Database, by discord username and discrim.
func (mc *MaldonClient) GetUser(ctx context.Context, discordUsername, discordDiscrim string) (DiscordUser, error) {
    var user DiscordUser
    collection := mc.client.Database("Maldon").Collection("users")

    filter := bson.M{"discordUsername": discordUsername, "discordDiscrim": discordDiscrim}
    err := collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        return DiscordUser{}, err
    }

    return user, nil
}

// SearchUsers gets all users from the Maldon Database.
// SearchUsers returns a list of all users in the database.
func (mc *MaldonClient) SearchUsers(ctx context.Context) ([]DiscordUser, error) {
    var users []DiscordUser
    collection := mc.client.Database("Maldon").Collection("users")

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user DiscordUser
        err := cursor.Decode(&user)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}
