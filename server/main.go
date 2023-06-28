package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// this holds the arguments passed to json rpc service
type Args struct {
	Name string
}

// twitterprofile holds
type TwitterProfile struct {
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Followers string `json:"followers,omitempty"`
	Following string `json:"following,omitempty"`
}

type JSONServer struct{}

//twitterprofile detail
func (t *JSONServer) TwitterProfileDetail(r *http.Request, args *Args, reply *TwitterProfile) error {
	var twitterProfiles []TwitterProfile
	// Read JSON file and load data
	raw, err := ioutil.ReadFile("./twitterProfile.json")
	if err != nil {
		log.Fatal("error: ", err)
		os.Exit(1)
	}

	//Unmarshall json raw data into twitter profiles array
	marshallerr := jsonparse.Unmarshal(raw, &twitterProfiles)
	if marshallerr != nil {
		log.Println("error:", marshallerr)
		os.Exit(1)
	}

	// Iterate over each twitter profile to find the given twitter profile
	for _, twitterProfile := range twitterProfiles {
		if twitterProfile.Name == args.Name {
			*reply = twitterProfile
			break
		}
	}

	return nil

}

/*

the RPC time server legacy

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// fill reply pointer to sent the data back
	*reply = time.Now().Unix()
	return nil
}
*/

func main() {
	// create a new RPC server
	// timeserver := new(TimeServer) no library
	// Create a new RPC server
	s := rpc.NewServer() // Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":9000", r)
}
