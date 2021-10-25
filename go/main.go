package main

import (

	"database/sql"
	"fmt"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "."
	dbname   = "shf_web"
  )


type get struct {
 
    FieldA string `form:"name"`
}




func sendInformation(c *gin.Context) {
	
  
	var b get
    c.Bind(&b)
 
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	panic(err)
	}
	defer db.Close()


	rows, err := db.Query("SELECT input FROM shf_web1 where hash_code=$1",b.FieldA)
	if err != nil {
	  // handle this error better than this
	  panic(err)
	}

	defer rows.Close()

	for rows.Next() {
	  var input string
	
	  err = rows.Scan(&input)
	  if err != nil {
		// handle this error
		panic(err)
	  }
	  c.JSON(200, gin.H{"input":input })
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
	  panic(err)
	}

}
 

func setInformationFromDatabase( c *gin.Context ) {
	var input string = c.PostForm("name")

	if (len(input) <8 ) {
    c.JSON(200, gin.H{"error": "this input is lower than 8 character ." })	
  } else {
	  
  hash_code := sha256.Sum256([]byte(input))
  fmt.Printf("%x\n", hash_code)
  
 resualt_hash:= fmt.Sprintf("%x\n", hash_code)


  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
"password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
db, err := sql.Open("postgres", psqlInfo)
if err != nil {
panic(err)
}
defer db.Close()

sqlStatement := `
INSERT INTO shf_web1 (input,hash_code)
VALUES ($1, $2)`
_, err = db.Exec(sqlStatement, input,resualt_hash)
if err != nil {
panic(err)
}
c.JSON(200, gin.H{"input":resualt_hash })

}

}


func main() {
    r := gin.Default()
	//give information and set into database
	r.POST("/post", setInformationFromDatabase)

	//give information from database and send for front
    r.GET("/get", sendInformation)


    r.Run(":2060")


}