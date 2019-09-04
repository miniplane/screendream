package downloader

import (
	"testing"
	"fmt"
)

func TestTimeconversion(t *testing.T) {
	fmt.Println("test")
	fmt.Println(parse_duration("1:01:14")) 	// should equal 3674
	fmt.Println(parse_duration("01:14")) 	// should equal 74
	fmt.Println(parse_duration("asd:asf")) 	// should equal 0 and print errors

}