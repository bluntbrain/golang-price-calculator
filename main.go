package main

import (
	"fmt"

	"ishanlakhwani.com/price-calculator/filemanager"
	"ishanlakhwani.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	// create done channel
	doneChans := make([]chan bool, len(taxRates))
	// create error channel
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {

		// add a channel for each item in taxRates, make keyword is used to make a channel
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error) // type here is important, error type

		// defining different name for each tax rate
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// go keyword is added to use goroutines, all processes will run in parallel
		// goroutines does not return values
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		// select is similar to switch but for channels, for different values emitted by a channel
		select {
		case err := <-errorChans[index]: // if error channel returns a value means an error has occured
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]: // if that particualr task is completed, print Done!
			fmt.Println("Done!")
		}
	}
}
