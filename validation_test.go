package main

import (
    "testing"
    "fmt"
)

func TestCheckMsg(t *testing.T) {
    msg := "kerfuffle this has a curse word"
    want := "**** this has a curse word"
    resp, err := checkMsg(msg)
    if !(resp == want) || err != nil {
        t.Errorf(`CheckMsg("kerfuffle this has a curse word") = %q, %v, want %#q`, resp, err, want)
    }
    fmt.Println(resp)
}

func TestCheckMsgCase(t *testing.T) {
    msg := "KERFUFFLE this has a curse word"
    want := "**** this has a curse word"
    resp, err := checkMsg(msg)
    if !(resp == want)|| err != nil {
        t.Errorf(`CheckMsg("KERFUFFLE this has a curse word") = %q, %v, want %s`, resp, err, want)
    }
    fmt.Println(resp)
}

func TestCheckMsgAllUpCase(t *testing.T) {
    msg := "KERFUFFLE! I AM MAD AT THE WORLD SO HEAR ME sharbert this has a FORNAX word"
    want := "KERFUFFLE! I AM MAD AT THE WORLD SO HEAR ME **** this has a **** word"
    resp, err := checkMsg(msg)
    if !(resp == want)|| err != nil {
        t.Errorf(`CheckMsg("KERFUFFLE! I AM MAD AT THE WORLD SO HEAR ME sharbert this has a FORNAX word") = %q, %v, want %s`, resp, err, want)
    }
    fmt.Println(resp)
}
