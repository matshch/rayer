package main

type Material interface {
	Scatter(r Ray, hit Hit) (*Ray, *Color)
}
