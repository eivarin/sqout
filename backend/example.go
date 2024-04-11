package main

import (
	"context"
	"sqout/endpoints"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	collection := client.Database("sqout").Collection("modules")
	c.Set("collection", collection)
}

func main() {

	r := gin.Default()

	r.Use(initDB)
	endpoints.SetupModulesRoutes(r)

	// r.GET("/ping", func(c *gin.Context) {
	// 	// Open the YAML file
	// 	file, err := os.Open("modules/ping/config.yaml")
	// 	if err != nil {
	// 		log.Fatalf("Error opening file: %v", err)
	// 	}
	// 	defer file.Close()

	// 	// Decode YAML data
	// 	var config Config
	// err = yaml.NewDecoder(file).Decode(&config)
	// if err != nil {
	// 	log.Fatalf("Error decoding YAML: %v", err)
	// }

	// 	// Extract the value of exe: command-name
	// 	commandName := config.Exe.CommandName
	// 	fmt.Println("Command name:", commandName)

	// 	fmt.Println("Flags:")
	// 	var length = len(config.Exe.Flags)
	// 	var args []string = make([]string, length)
	// 	for i := 0; i < length; i++ {
	// 		args[i] = config.Exe.Flags[i].Value
	// 		fmt.Printf("- Name: %s, Type: %s, Default: %v\n", config.Exe.Flags[i].Name, config.Exe.Flags[i].Type, config.Exe.Flags[i].Value)
	// 	}
	// 	// how to implement a pipe ?
	// 	c1 := exec.Command(commandName, args...)
	// 	c2 := exec.Command("modules/ping/parse.sh")

	// 	b := new(strings.Builder)

	// 	c2.Stdin, _ = c1.StdoutPipe()
	// 	c2.Stdout = b
	// 	_ = c2.Start()
	// 	_ = c1.Run()
	// 	_ = c2.Wait()

	// 	// print the output
	// 	fmt.Println("Output: ", b.String())

	// 	str := b.String()

	// 	// Define a slice of structs to hold the parsed data
	// 	var data []map[string]string

	// 	// Unmarshal the string into the slice of structs
	// 	err = json.Unmarshal([]byte(str), &data)
	// 	if err != nil {
	// 		fmt.Println("Error Unmarshal:", err)
	// 		return
	// 	}

	// 	// Print the JSON data
	// 	jsonData, err := json.MarshalIndent(data, "", "  ")
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	fmt.Println(string(jsonData))

	// 	res, _ := collection.InsertOne(context.Background(), config)
	// 	id := res.InsertedID
	// 	fmt.Println("Inserted ID:", id)

	// 	// POST the json data to the mongoDB

	// 	// Return the JSON data
	// 	c.JSON(http.StatusOK, data)
	//	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
