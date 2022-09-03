
import "go.mongodb.org/mongo-driver/mongo"

clientOptions := options.Client().
    ApplyURI("mongodb+srv://vasu:vasu@cluster0.3gnlkhi.mongodb.net/?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
