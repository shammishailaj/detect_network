package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Printf("\n\n")
	// get available network interfaces for
	// this machine
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range interfaces {

		fmt.Printf("Name : %v \n", i.Name)

		byNameInterface, err := net.InterfaceByName(i.Name)

		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println("Interface by Name : ", byNameInterface)

		addresses, err := byNameInterface.Addrs()

		if err != nil || len(addresses) <= 0 {
			fmt.Printf("No Addresses found for: %s\n", i.Name)
		}

		for k, v := range addresses {

			fmt.Printf("Interface Address #%v : %v\n", k, v.String())
		}
		fmt.Println("------------------------------------")

	}

	IPAddressByNetworkInteface("wlo1")
	fmt.Printf("\n\n")
}

func IPAddressByNetworkInteface(networkInterface string) (string, error) {
	interfaceDetails, interfaceDetailsErr := net.InterfaceByName(networkInterface)

	if interfaceDetailsErr != nil {
		fmt.Println(interfaceDetailsErr)
		return "", interfaceDetailsErr
	}

	addresses, adrressErr := interfaceDetails.Addrs()

	if adrressErr != nil {
		fmt.Println(adrressErr)
		return "", adrressErr
	} else {

		fmt.Printf("Interface Details for %s: ", networkInterface)
		for k, v := range addresses {
			fmt.Printf("\nAddress #%d: %s", k, v.String())
		}
		return "", nil
	}
}
