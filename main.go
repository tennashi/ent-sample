package main

import (
	"context"
	"fmt"

	"github.com/tennashi/ent-sample/ent"
	"github.com/tennashi/ent-sample/ent/namespace"
	"github.com/tennashi/ent-sample/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func initClient() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

func createEntities(ctx context.Context, c *ent.Client) error {
	fmt.Println("=== Create users ===")

	userNames := []string{"user-a", "user-b", "user-c"}
	bulk := make([]*ent.UserCreate, len(userNames))
	for i, name := range userNames {
		bulk[i] = c.User.Create().SetID("user:" + fmt.Sprintf("%.3d", i+1)).SetName(name)
	}
	us, err := c.User.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	fmt.Println(us)

	fmt.Println("=== Create namespaces ===")

	nsa, err := c.Namespace.Create().SetID("namespace:001").SetName("namespace-a").
		AddMembers(us[0], us[1], us[2]).AddAdmins(us[0]).Save(ctx)
	if err != nil {
		return err
	}

	fmt.Println(nsa)

	nsb, err := c.Namespace.Create().SetID("namespace:002").SetName("namespace-b").
		AddMembers(us[0], us[1]).AddAdmins(us[0], us[1]).Save(ctx)
	if err != nil {
		return err
	}

	fmt.Println(nsb)

	nsc, err := c.Namespace.Create().SetID("namespace:003").SetName("namespace-c").
		AddMembers(us[2]).AddAdmins(us[0]).Save(ctx)
	if err != nil {
		return err
	}

	fmt.Println(nsc)

	return nil
}

func main() {
	c, err := initClient()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	ctx := context.Background()

	err = createEntities(ctx, c)
	if err != nil {
		panic(err)
	}

	fmt.Println("=== Query members of namespace-a ===")

	us, err := c.Namespace.Query().Where(namespace.NameEQ("namespace-a")).QueryMembers().All(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(us)

	fmt.Println("=== Query namespaces to which user-a belongs ===")

	nss, err := c.User.Query().Where(user.NameEQ("user-a")).QueryNamespaces().All(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(nss)
}
