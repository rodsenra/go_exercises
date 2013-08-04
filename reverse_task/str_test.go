package str
 
import (
	"testing"
)
 
func TestReverse(t *testing.T) {
	input := "世界"
	//input := "ROMA"	
        output := Reverse(input)
	expected := "界世"
	//expected := "AMOR"
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
