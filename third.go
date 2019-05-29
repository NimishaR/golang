package main

import ("fmt"
 		"net/http"
 		"html/template")

type NewsPage struct{
	Title string
	News string
}

func news_handler(w http.ResponseWriter, r *http.Request){
	p :=NewsPage{"News ZZZZZZZZZ","some news"}
	t, _ := template.ParseFiles("News.html")
	t.Execute(w,p)
}

//method: we have to associate functions to structs to make them behave as method
func (n NewsPage) getnews() string {
   return n.Title;
}

func index_handler(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "<h1>I'm Here</h1> \n <button> Hello </button>")
}

func home_handler(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "<button> Hello </button>")
}

func main(){
	ne :=NewsPage{"News ZZZZZZZZZ","some news"}
	fmt.Println(ne.getnews());
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/home",home_handler)
	http.HandleFunc("/news",news_handler)
	http.ListenAndServe(":8080",nil)
} 		