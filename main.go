package main 

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	//Price string `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fname string `json:"firstname"`
	Sname string `json:"lastname"`
}

var books []Book


func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)


}



func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag :=0
	for _,item := range books {
		if item.ID == params["id"]{
			flag=1
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	if flag==0 {
		fmt.Fprintf(w,"The book does not exist")
	}

}


func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag :=0
	for index,item := range books {
		if item.ID == params["id"] {
			flag=1
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	if flag==0 {
		fmt.Fprintf(w,"The book does not exist")
	}
}



func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag :=0
	for index,item := range books {
		if item.ID == params["id"] {
			flag=1
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			break
		}
	}
	if flag==0 {
		fmt.Fprintf(w,"The book does not exist")
	}
}


func main () {
	r := mux.NewRouter()

	books=append(books, Book{ID: "1", Title: "Abc", Author: &Author{Fname: "R", Sname: "B"}})
	books=append(books, Book{ID: "2", Title: "Xyz", Author: &Author{Fname: "A", Sname: "N"}})
	r.HandleFunc("/createBook",createBook).Methods("POST")
	r.HandleFunc("/getBooks",getBooks).Methods("GET")
	r.HandleFunc("/getBook/{id}",getBook).Methods("GET")
	r.HandleFunc("/updateBook/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/deleteBook/{id}",deleteBook).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}


