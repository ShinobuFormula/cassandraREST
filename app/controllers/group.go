package controllers

import (
	"github.com/gocql/gocql"
	"github.com/revel/revel"
	"log"
	"strconv"
)


type Group struct {
	*revel.Controller
}

func (c Group) AllGroup() revel.Result{
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

	iter := session.Query(`SELECT * FROM group_by_id`).Consistency(gocql.One).Iter()
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