package main

import (
	"go-training/cfg"
	"go-training/lib/assignments"
	"go-training/lib/formletter"
)

func main() {
	assignments.A11()
	assignments.A14()
	assignments.A15()
	assignments.A16("Fred", "cloudy", "meatloaf")
	formletter.Letter(cfg.Name, cfg.Weather, cfg.Snack)
}
