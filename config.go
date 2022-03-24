package main

type Config struct {
	Index struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"index" required:"true" split_words:"true"`
}
