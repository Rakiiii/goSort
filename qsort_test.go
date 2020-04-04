package gosort
import (
	"fmt"
	"testing"
)

func TestQuicksortFloat64(t *testing.T){
	fmt.Println("Start TestQuicksortFloat64")
	unsorted := []float64{2.0,6.0,1.0,4.0}
	sorted := []float64{1.0,2.0,4.0,6.0}
	unsortedAfter := QuicksortFloat64(unsorted)
	errFlag := true  
	
	for i:=0 ; i < len(unsorted);i++{
		if sorted[i] != unsorted[i]{
			t.Error("Wrong value at pos:",i," expected:",sorted[i]," found:",unsortedAfter[i])
			errFlag = false
		}
	}
	if errFlag{
		fmt.Println("TestQuicksortFloat64=[ok]")
	}
}

func TestQuicksortInt(t *testing.T){
	fmt.Println("Start TestQuicksortInt")
	unsorted := []int{2,6,1,4}
	sorted := []int{1,2,4,6}
	unsortedAfter := QuicksortInt(unsorted)
	errFlag := true
	for i:=0 ; i < len(unsorted);i++{
		if sorted[i] != unsortedAfter[i]{
			t.Error("Wrong value at pos:",i," expected:",sorted[i]," found:",unsortedAfter[i])
			errFlag = false
		}
	}
	if errFlag{
		fmt.Println("TestQuicksortInt=[ok]")
	}
}