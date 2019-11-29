package controllers

import (
	"github.com/gocql/gocql"
	"github.com/revel/revel"
	"log"
	"strconv"
)


type User struct {
	*revel.Controller
}

func (c User) AllUser() revel.Result{
	cluster := gocql.NewCluster("192.168.109.137")
	cluster.Keyspace = "squeat_db"
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	data := make(map[string]interface{})
	stuff := make(map[string]interface{})
	i := 0

	iter := session.Query(`SELECT * FROM user_by_id`).Consistency(gocql.One).Iter()
	for iter.MapScan(stuff){
		data[strconv.Itoa(i)] = stuff
		i++
		stuff = make(map[string]interface{})
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}

func (c User) UserById() revel.Result{
	cluster := gocql.NewCluster("192.168.109.137")
	cluster.Keyspace = "squeat_db"
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	param := make(map[string]interface{})
	c.Params.BindJSON(&param)

	data := make(map[string]interface{})

	if err := session.Query(`SELECT * FROM user_by_id WHERE id = ?`, param["id"]).Consistency(gocql.One).MapScan(data); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}

func (c User) UserByMail() revel.Result{
	cluster := gocql.NewCluster("192.168.109.137")
	cluster.Keyspace = "squeat_db"
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	param := make(map[string]interface{})
	c.Params.BindJSON(&param)

	data := make(map[string]interface{})

	if err := session.Query(`SELECT * FROM user_by_mail WHERE mail = ?`, param["mail"]).Consistency(gocql.One).MapScan(data); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}

func (c User) UserByUsername(username string) revel.Result{
	cluster := gocql.NewCluster("192.168.109.137")
	cluster.Keyspace = "squeat_db"
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	param := make(map[string]interface{})
	c.Params.BindJSON(&param)

	data := make(map[string]interface{})

	if err := session.Query(`SELECT * FROM user_by_username WHERE username = ?`, param["username"]).Consistency(gocql.One).MapScan(data); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}

