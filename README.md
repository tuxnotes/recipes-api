## golang code for insert users to mongodb

```golang
func main() {
   users := map[string]string{
       "admin":      "fCRmh4Q2J7Rseqkz",
       "packt":      "RE4zfHB35VPtTkbT",
       "mlabouardy": "L3nSFRcZzNQ67bcc",
   }
   ctx := context.Background()
   client, err := mongo.Connect(ctx,
       options.Client().ApplyURI(os.Getenv("MONGO_URI")))
   if err = client.Ping(context.TODO(),
          readpref.Primary()); err != nil {
       log.Fatal(err)
   }
   collection := client.Database(os.Getenv(
       "MONGO_DATABASE")).Collection("users")
   h := sha256.New()
   for username, password := range users {
       collection.InsertOne(ctx, bson.M{
           "username": username,
           "password": string(h.Sum([]byte(password))),
       })
   }
}
```
