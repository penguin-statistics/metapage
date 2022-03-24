package main

type Config struct {
	Address string `required:"true" split_words:"true" default:":8060"`
	Index   struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"index" required:"true" split_words:"true"`
}
