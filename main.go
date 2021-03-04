package main

// import("~/../golang/golang/tpl")
import("test")
import("github.com/flosch/pongo2")
import("fmt")
import("net/http")

func queryHandle(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hello all !!! request: %s\n\n", req)
	// var mainPage = tpl.MainPage.ExecuteWriter(_, w)
	var mainPage = pongo2.Must(pongo2.FromFile("templates/main.html"))
	mainPage.ExecuteWriter(nil, w)
}

func hello(w http.ResponseWriter, req *http.Request) {
	var names []string;

	names = append(names, "John")
	names = append(names, "Joe")
	names = append(names, "Silver")

	fmt.Fprintf(w, "Hello all !!! request: %s\n\n", req)

	for _, name := range names {
		fmt.Fprintf(w, "Hello %s !!!\n", name)
	}
}

func main() {
	test.Boo()
	test.Boo2()
	// http.HandleFunc("/", queryHandle)
	// http.HandleFunc("/hello", hello)
	// http.ListenAndServe(":8090", nil)
}
