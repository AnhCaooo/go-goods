package testcontainer

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Represent test container structure
type Container struct {
	Ctr testcontainers.Container // container
	URI string                   // URI of test container
}

// Setup receives test image then prepare test container for that image and return that container.
// After setup container, just proceed as normal as you deal with real container such as:
// connect to it & disconnect after the whole test suite.
//
// Usage example:
//
//	func TestMain(m *testing.M) {
//		var code = 1
//		defer func() { os.Exit(code) }()
//
//		ctx := context.Background()
//
//		mongoContainer, err := Setup(ctx, mongoDB)
//		if err != nil {
//			log.Fatal("fail to setup mongodb container")
//		}
//
//		client, err := mongo.Connect(options.Client().ApplyURI(mongoContainer.URI))
//		if err != nil {
//			log.Fatal("Could not connect to mongodb, fallback to mem repository implementations", err)
//		}
//
//		defer func() {
//			if err = client.Disconnect(ctx); err != nil {
//				panic(err)
//			}
//		}()
//
//		code = m.Run()
//	}
func Setup(ctx context.Context, image Image) (*Container, error) {
	cont, uri, err := prepareContainer(ctx, image)
	if err != nil {
		return nil, err
	}
	return &Container{
		Ctr: cont,
		URI: uri,
	}, nil
}

// Represent test image structure
type Image struct {
	Name string // name of test image. Example: "mongo"
	Port string // port of test image. Example: "27017"
}

// prepareContainer prepare test container for current image
func prepareContainer(ctx context.Context, image Image) (testcontainers.Container, string, error) {
	req := testcontainers.ContainerRequest{
		Image:        image.Name,
		ExposedPorts: []string{image.Port + "/tcp"},
		WaitingFor:   wait.ForListeningPort(nat.Port(image.Port)),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, "", err
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(image.Port))
	if err != nil {
		return nil, "", err
	}

	var uri string
	switch image {
	case MongoDB:
		uri = fmt.Sprintf("mongodb://%s:%s", hostIP, mappedPort.Port())
	default:
		return nil, "", errors.New("TestContainers: unsupported image: " + image.Name)
	}

	log.Printf("TestContainers: container %s is now running at %s\n", req.Image, uri)
	return container, uri, nil
}
