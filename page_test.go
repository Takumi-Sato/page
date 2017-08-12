package page

import (
  "testing"
)

func TestPage(t *testing.T){
	
	if _, err := Get("https://google.co.jp"); err != nil {
	    t.Errorf("Incorrect.") 
	}
	
	/*
	if _, err := Get("https://hoge.foo"); err != nil {
	    t.Errorf("Incorrect.") 
	}
	*/
}