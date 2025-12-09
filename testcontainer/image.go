package testcontainer

var (
	// Test image for MongoDB
	MongoDB = Image{
		Name: "mongo:latest",
		Port: "27017",
	}
)
