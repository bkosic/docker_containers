package main

import (
  "fmt"
  "log"
  "github.com/gocql/gocql"
  "net"
  "net/http"
  "net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
 resp.Write([]byte("<h1>Hello, 3FS</h1>\n<p>Behold my Go web app.</p>"))
}

func main() {

  listener, _ := net.Listen("tcp", "127.0.0.1:9000")
  srv := new(FastCGIServer)
  fcgi.Serve(listener, srv)

	// // connect to the cluster
	// cluster := gocql.NewCluster("db-cassandra")
  // cluster.ProtoVersion = 4 // cluster.ProtoVersion = 4
	// cluster.Keyspace = "projekt1"
	// session, _ := cluster.CreateSession()
	// defer session.Close()
  //
  // // Create keyspace projekt1 if doesn't exits
  // if err := session.Query("CREATE KEYSPACE IF NOT EXISTS projekt1 WITH replication = {'class': 'SimpleStrategy','replication_factor' : 3}").Exec(); err != nil {
  //   log.Fatal(err)
  // }
  //
  // // Create table if doesn't exits
  // if err := session.Query("CREATE TABLE IF NOT EXISTS projekt1.test(cas timestamp)").Exec(); err != nil {
  //   log.Fatal(err)
  // }
  //
	// // insert current timestamp
	// if err := session.Query("INSERT INTO test (cas) VALUES (toTimestamp(now())").Exec(); err != nil {
	// 	log.Fatal(err)
	// }
  //
	// // Use select to get the timestamp we just entered
	// var cas string
  //
	// // Select and show the change
	// iter := session.Query("SELECT cas FROM test").Iter()
	// for iter.Scan(&cas) {
	// 	fmt.Println(cas)
	// }
	// if err := iter.Close(); err != nil {
	// 	log.Fatal(err)
	// }
}
