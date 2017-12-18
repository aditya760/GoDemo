package main
import "testing"


func Test_first(t *testing.T){
  if(!check(10)){
		t.Error("Unexpected Result")
	}
}

func Test_first1(t *testing.T){
  if(check(5)){
		t.Error("Unexpected Result")
	}
}