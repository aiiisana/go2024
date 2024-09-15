package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

func ToJSON(p Product) (string, error) {
    data, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func FromJSON(jsonStr string) (Product, error) {
    var p Product
    err := json.Unmarshal([]byte(jsonStr), &p)
    if err != nil {
        return Product{}, err
    }
    return p, nil
}

func main() {
    p := Product{Name: "Laptop", Price: 999.99, Quantity: 10}

    jsonStr, err := ToJSON(p)
    if err != nil {
        fmt.Println("Error converting to JSON:", err)
        return
    }
    fmt.Println("JSON:", jsonStr)

    decodedProduct, err := FromJSON(jsonStr)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }
    fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}

// How do you work with JSON in Go?
// encoding using json.Marshal to convert a struct to JSON or decoding using json.Unmarshal

// What role do struct tags play in JSON encoding/decoding?
// define how struct fields are named in JSON and control encoding/decoding

// How do you handle errors that may occur during JSON encoding/decoding?
// check the error returned by json.Marshal/Unmarshal to handle issues
