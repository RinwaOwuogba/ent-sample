package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"

	// _ "github.com/mattn/go-sqlite3"

	"github.com/rinwaowuogba/ent-sample/db/ent"
	"github.com/rinwaowuogba/ent-sample/db/ent/car"
	"github.com/rinwaowuogba/ent-sample/db/ent/group"
	"github.com/rinwaowuogba/ent-sample/db/ent/user"
)

func runTest(c *ent.Client) {
	ctx := context.TODO()

	_, err := CreateUser(ctx, c)
	if err != nil {
		log.Fatalf("unable to continue test: %v", err)
	}

	_, err = QueryUser(ctx, c)
	if err != nil {
		log.Fatalf("unable to continue test: %v", err)
	}

	// u, err := CreateCars(ctx, c)
	// if err != nil {
	// 	log.Fatalf("unable to continue test: %v", err)
	// }

	// fmt.Println()
	// fmt.Println()

	// CreateGraph(ctx, c)
	// QueryGroupWithUsers(ctx, c)
	// QueryArielCars(ctx, c)
	// QueryGithub(ctx, c)
	// QueryCarUsers(ctx, u)
	// QueryCars(ctx, u)
	// if err != nil {
	// 	log.Fatalf("unable to continue test: %v", err)
	// }
}

func main() {
	// client, err := ent.Open(dialect.SQLite, "file:en	t?mode=memory&cache=shared&_fk=1")
	// connectionString := "host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
	// connectionString := "postgresql://postgres:postgres@127.0.0.1/test_db?sslmode=disable"

	// todo: replace with pgx
	connectionString := "host=127.0.0.1 port=5432 user=postgres dbname=test_db password=postgres sslmode=disable"
	client, err := ent.Open(dialect.Postgres, connectionString)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	runTest(client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)

	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	ford, err := a8m.QueryCars().
		Where(car.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)

	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}

	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner, %w", c.Model, err)
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}

	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}

	// Create the cars
	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		// Attach this car to Ariel.
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()).
		// Attach this car to Ariel.
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		// Attach this graph to Neta.
		SetOwner(neta).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Create the groups, and add their users in the creation.
	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")).
		QueryUsers().
		QueryCars().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	// Get "Ariel" from previous steps.
	a8m := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name("Ariel"),
		).
		OnlyX(ctx)
	cars, err := a8m.
		QueryGroups(). // (Group(Name=GitHub), Group(Name=GitLab),)
		QueryUsers().  // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
		QueryCars().   //
		Where(         //
			car.Not( //  Get Neta and Ariel cars, but filter out
				car.Model("Mazda"), //  those who named "Mazda"
			), //
		). //
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}
