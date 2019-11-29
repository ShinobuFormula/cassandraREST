package controllers

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/revel/revel"
	"log"
	"strconv"
)


type Meal struct {
	*revel.Controller
}

func (c Meal) AllMeal() revel.Result{
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

	iter := session.Query(`SELECT * FROM meal_by_id`).Consistency(gocql.One).Iter()
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


func (c Meal) MealById() revel.Result{
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

	if err := session.Query(`SELECT * FROM meal_by_id WHERE id = ?`, param["id"]).Consistency(gocql.One).MapScan(data); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}

func (c Meal) MealByGroup() revel.Result{
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

	if err := session.Query(`SELECT * FROM meal_by_group WHERE group = ?`, param["id"]).Consistency(gocql.One).MapScan(data); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}


func (c Meal) MealByName() revel.Result{
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
	stuff := make(map[string]interface{})
	i := 0

	iter := session.Query(`SELECT * FROM  meal_by_name WHERE id = ? AND name = ?`, param["id"], param["name"]).Consistency(gocql.One).Iter()
	for iter.MapScan(stuff) {
		fmt.Println(i)
		data[strconv.Itoa(i)] = stuff
		i++
		stuff = make(map[string]interface{})
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)

}

func (c Meal) MealByMember() revel.Result{
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
	stuff := make(map[string]interface{})
	i := 0

	iter := session.Query(`SELECT * FROM  meal_by_id WHERE members CONTAINS ?`, param["id"]).Consistency(gocql.One).Iter()
	for iter.MapScan(stuff) {
		fmt.Println(i)
		data[strconv.Itoa(i)] = stuff
		i++
		stuff = make(map[string]interface{})
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(data)
}