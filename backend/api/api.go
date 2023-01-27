package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ryogrid/SamehadaDB/samehada"
)

type ListItem struct {
	Id   int32  `json:"id"`
	Item string `json:"item"`
	//Done bool   `json:"done"`
	Done int32 `json:"done"`
}

//var db *sql.DB
var db *samehada.SamehadaDB
var dbLock sync.Mutex
var err error

var nextId int32 = 0

func SetupPostgres() {
	//// db, err = sql.Open("postgres", "postgres://postgres:password@postgres/todo?sslmode=disable")
	//
	//// when running locally
	//dburl := os.Getenv("DATABASE_URL")
	////db, err = sql.Open("postgres", "postgres://postgres:password@localhost/todo?sslmode=disable")
	//db, err = sql.Open("postgres", dburl)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//if err = db.Ping(); err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//log.Println("connected to postgres")

	db = samehada.NewSamehadaDB("./demo", 10*1024) // buffer pool capacity max is 10MB
	dbLock.Lock()
	db.ExecuteSQL("CREATE table list (id int, item char(256), done int);")
	dbLock.Unlock()
}

// CRUD: Create Read Update Delete API Format

// List all todo items
func TodoItems(c *gin.Context) {
	// Use SELECT Query to obtain all rows
	//rows, err := db.Query("SELECT * FROM list")
	dbLock.Lock()
	err, rows := db.ExecuteSQL("SELECT * FROM list")
	dbLock.Unlock()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
	}

	// Get all rows and add into items
	items := make([]ListItem, 0)

	if len(rows) != 0 {
		//defer rows.Close()
		//for rows.Next() {
		for _, row := range rows {
			// Individual row processing
			item := ListItem{}
			//if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
			//	fmt.Println(err.Error())
			//	c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			//}
			//item.Item = strings.TrimSpace(item.Item)
			item.Id = row[0].(int32)
			item.Item = strings.TrimSpace(row[1].(string))
			item.Done = row[2].(int32)
			items = append(items, item)
		}
	}

	// Return JSON object of all rows
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// Create todo item and add to DB
func CreateTodoItem(c *gin.Context) {
	item := c.Param("item")

	// Validate item
	if len(item) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an item"})
	} else {
		// Create todo item
		var TodoItem ListItem

		dbLock.Lock()
		TodoItem.Id = nextId
		nextId++
		TodoItem.Item = item
		//TodoItem.Done = false
		TodoItem.Done = 0

		// Insert item to DB
		err, _ := db.ExecuteSQL(fmt.Sprintf("INSERT INTO list(id, item, done) VALUES(%d, '%s', %d);", TodoItem.Id, TodoItem.Item, TodoItem.Done))
		dbLock.Unlock()
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})

		}

		// Log message
		log.Println("created todo item", item)

		// Return success response
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusCreated, gin.H{"items": &TodoItem})
	}
}

// Update todo item
func UpdateTodoItem(c *gin.Context) {
	idParam := c.Param("id")
	doneParam := c.Param("done")

	fmt.Printf("doneParam: %s\n", doneParam)

	// Validate id and done
	if len(idParam) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an id"})
		return
	} else if len(doneParam) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter a done state"})
		return
	} else {
		var done int32
		if doneParam == "true" {
			done = 1
		} else if doneParam == "false" {
			done = 0
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter a done state"})
			return
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an id"})
			return
		}

		// Find and update the todo item
		//var exists bool
		dbLock.Lock()
		//err := db.QueryRow("SELECT * FROM list WHERE id=$1;", id).Scan(&exists)
		err, rows := db.ExecuteSQL(fmt.Sprintf("SELECT * FROM list WHERE id=%d;", id))
		dbLock.Unlock()
		if err != nil || len(rows) == 0 {
			fmt.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			//_, err := db.Query("UPDATE list SET done=$1 WHERE id=$2;", done, id)
			dbLock.Lock()
			err, _ := db.ExecuteSQL(fmt.Sprintf("UPDATE list SET done=%d WHERE id=%d;", done, id))
			dbLock.Unlock()
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
				return
			}

			// Log message
			log.Println("updated todo item", id, done)

			// Return success response
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
			c.JSON(http.StatusOK, gin.H{"message": "successfully updated todo item", "todo": id})
		}
	}
}

// Delete todo item
func DeleteTodoItem(c *gin.Context) {
	idParam := c.Param("id")

	// Validate id
	if len(idParam) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an id"})
	} else {
		// Find and delete the todo item
		//var exists bool
		//err := db.QueryRow("SELECT * FROM list WHERE id=$1;", id).Scan(&exists)

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an id"})
			return
		}

		dbLock.Lock()
		err, rows := db.ExecuteSQL(fmt.Sprintf("SELECT * FROM list WHERE id=%d;", id))
		dbLock.Unlock()
		if err != nil && len(rows) == 0 {
			fmt.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			//_, err = db.Query("DELETE FROM list WHERE id=$1;", id)
			dbLock.Lock()
			err, rows = db.ExecuteSQL(fmt.Sprintf("DELETE FROM list WHERE id=%d;", id))
			dbLock.Unlock()
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			}

			// Log message
			log.Println("deleted todo item", id)

			// Return success response
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
			c.JSON(http.StatusOK, gin.H{"message": "successfully deleted todo item", "todo": id})
		}
	}
}

// Add Filter API
