package main

import (
	"fmt"
	"time"
)

const MAX_SPAREPARTS = 100
const MAX_CUSTOMERS = 100
const MAX_TRANSACTIONS = 100

type StructSparePart struct {
	ID         int
	Name       string
	Price      float64
	Stock      int
	UsageCount int
}

type StructCustomer struct {
	ID   int
	Name string
}

type StructTransaction struct {
	ID         int
	CustomerID int
	Date       time.Time
	ServiceFee float64
	SpareParts [MAX_SPAREPARTS]int
	TotalPrice float64
}

// Data global array statis
var ArrSpareParts [MAX_SPAREPARTS]StructSparePart
var ArrCustomers [MAX_CUSTOMERS]StructCustomer
var ArrTransactions [MAX_TRANSACTIONS]StructTransaction

// Sinta -
// Menampilkan spare parts
func DisplaySpareParts() {
	fmt.Println("List of Spare Parts:")
	fmt.Println("---------------------------------------------------")
	for _, part := range ArrSpareParts {
		if part.ID != 0 {
			fmt.Printf("ID         : %03d\n", part.ID)
			fmt.Printf("Name       : %s\n", part.Name)
			fmt.Printf("Price      : %.2f\n", part.Price)
			fmt.Printf("Stock      : %d\n", part.Stock)
			fmt.Printf("UsageCount : %d\n", part.UsageCount)
			fmt.Println("---------------------------------------------------")
		}
	}
}

// Menambahkan spare part baru
func AddSparePart(newSparePart StructSparePart) bool {
	for i := 0; i < MAX_SPAREPARTS; i++ {
		if ArrSpareParts[i].ID == 0 {
			ArrSpareParts[i] = newSparePart
			return true
		}
	}
	return false
}

// Mengubah data spare part
func UpdateSparePart(id int, updatedPart StructSparePart) bool {
	for i := 0; i < MAX_SPAREPARTS; i++ {
		if ArrSpareParts[i].ID == id {
			ArrSpareParts[i] = updatedPart
			return true
		}
	}
	return false
}

// Menghapus spare part
func DeleteSparePart(id int) bool {
	for i := 0; i < MAX_SPAREPARTS; i++ {
		if ArrSpareParts[i].ID == id {
			ArrSpareParts[i] = StructSparePart{}
			return true
		}
	}
	return false
}

// Menampilkan pelanggan
func DisplayCustomers() {
	fmt.Println("List of Customers:")
	fmt.Println("---------------------------------------------------")
	for _, customer := range ArrCustomers {
		if customer.ID != 0 {
			fmt.Printf("ID   : %03d\n", customer.ID)
			fmt.Printf("Name : %s\n", customer.Name)
			fmt.Println("---------------------------------------------------")
		}
	}
}

// Menambahkan pelanggan baru
func AddCustomer(newCustomer StructCustomer) bool {
	for i := 0; i < MAX_CUSTOMERS; i++ {
		if ArrCustomers[i].ID == 0 {
			ArrCustomers[i] = newCustomer
			return true
		}
	}
	return false
}

// Mengubah data pelanggan
func UpdateCustomer(id int, updatedCustomer StructCustomer) bool {
	for i := 0; i < MAX_CUSTOMERS; i++ {
		if ArrCustomers[i].ID == id {
			ArrCustomers[i] = updatedCustomer
			return true
		}
	}
	return false
}

// Menghapus pelanggan
func DeleteCustomer(id int) bool {
	for i := 0; i < MAX_CUSTOMERS; i++ {
		if ArrCustomers[i].ID == id {
			ArrCustomers[i] = StructCustomer{}
			return true
		}
	}
	return false
}

// Menampilkan transaksi dalam format daftar
func DisplayTransactions() {
	fmt.Println("List of Transactions:")
	fmt.Println("---------------------------------------------------")
	for _, transaction := range ArrTransactions {
		if transaction.ID != 0 {
			fmt.Printf("ID         : %03d\n", transaction.ID)
			fmt.Printf("CustomerID : %03d\n", transaction.CustomerID)
			fmt.Printf("Date       : %s\n", transaction.Date.Format("2006-01-02"))
			fmt.Printf("ServiceFee : %.2f\n", transaction.ServiceFee)
			fmt.Printf("TotalPrice : %.2f\n", transaction.TotalPrice)
			fmt.Println("---------------------------------------------------")
		}
	}
}

// Menambahkan transaksi baru
func AddTransaction(newTransaction StructTransaction) bool {
	for i := 0; i < MAX_TRANSACTIONS; i++ {
		if ArrTransactions[i].ID == 0 {
			ArrTransactions[i] = newTransaction
			return true
		}
	}
	return false
}

// Mengubah data transaksi
func UpdateTransaction(id int, updatedTransaction StructTransaction) bool {
	for i := 0; i < MAX_TRANSACTIONS; i++ {
		if ArrTransactions[i].ID == id {
			ArrTransactions[i] = updatedTransaction
			return true
		}
	}
	return false
}

// Menghapus transaksi
func DeleteTransaction(id int) bool {
	for i := 0; i < MAX_TRANSACTIONS; i++ {
		if ArrTransactions[i].ID == id {
			ArrTransactions[i] = StructTransaction{}
			return true
		}
	}
	return false
}

// - Sinta

// Subprogram Perhitungan - Najwa
func CalculateServiceFee(serviceFee float64, spareParts [MAX_SPAREPARTS]int, ArrSpareParts [MAX_SPAREPARTS]StructSparePart) float64 {
	var totalPrice float64 = serviceFee

	// Menghitung total harga berdasarkan spare-parts yang digunakan
	for i := 0; i < MAX_SPAREPARTS; i++ {
		if spareParts[i] > 0 { // Jika spare part digunakan (jumlah lebih besar dari 0)
			totalPrice += float64(spareParts[i]) * ArrSpareParts[i].Price
		}
	}
	return totalPrice
}

// Searching - Wahyu Bagus S
func FindCustomersByServicePeriod(startDate, endDate time.Time) [MAX_CUSTOMERS]StructCustomer {
	var result [MAX_CUSTOMERS]StructCustomer
	index := 0

	for _, transaction := range ArrTransactions {
		if transaction.Date.After(startDate) && transaction.Date.Before(endDate) {
			for _, customer := range ArrCustomers {
				if customer.ID == transaction.CustomerID {
					result[index] = customer
					index++
				}
			}
		}
	}
	return result
}

// ke-2
func FindCustomersBySparePart(sparePartID int) [MAX_CUSTOMERS]StructCustomer {
	var result [MAX_CUSTOMERS]StructCustomer
	index := 0

	for _, transaction := range ArrTransactions {
		for i, qty := range transaction.SpareParts {
			if qty > 0 && i < len(ArrSpareParts) && ArrSpareParts[i].ID == sparePartID {
				for _, customer := range ArrCustomers {
					if customer.ID == transaction.CustomerID {
						result[index] = customer
						index++
					}
				}
			}
		}
	}
	return result
}

// SortSparePartsByUsage mengurutkan spare-part berdasarkan jumlah penggunaan (Selection Sort) - Anisa
func SortSparePartsByUsage(parts [MAX_SPAREPARTS]StructSparePart, order string) [MAX_SPAREPARTS]StructSparePart {
	sortedParts := parts
	for i := 0; i < MAX_SPAREPARTS-1; i++ {
		selectedIdx := i
		for j := i + 1; j < MAX_SPAREPARTS; j++ {
			if (order == "ascending" && sortedParts[j].UsageCount < sortedParts[selectedIdx].UsageCount) ||
				(order == "descending" && sortedParts[j].UsageCount > sortedParts[selectedIdx].UsageCount) {
				selectedIdx = j
			}
		}

		sortedParts[i], sortedParts[selectedIdx] = sortedParts[selectedIdx], sortedParts[i]
	}
	return sortedParts
}

// Sorting Customers By Name (Insertion sort) - Anisa
func SortCustomersByName(custs [MAX_CUSTOMERS]StructCustomer, order string) [MAX_CUSTOMERS]StructCustomer {
	sortedCusts := custs
	for i := 1; i < MAX_CUSTOMERS; i++ {
		key := sortedCusts[i]
		j := i - 1
		for j >= 0 && ((order == "ascending" && sortedCusts[j].Name > key.Name) ||
			(order == "descending" && sortedCusts[j].Name < key.Name)) {
			sortedCusts[j+1] = sortedCusts[j]
			j--
		}
		sortedCusts[j+1] = key
	}
	return sortedCusts
}

func sparePartMenu() {
	var choice int
	for {
		fmt.Println("\nSpare Part Menu:")
		fmt.Println("1. Display Spare Part Data")
		fmt.Println("2. Add Spare Part Data")
		fmt.Println("3. Update Spare Part Data")
		fmt.Println("4. Delete Spare Part Data")
		fmt.Println("5. Sort Spare Part Data")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplaySpareParts()
		case 2:
			var newPart StructSparePart
			fmt.Print("Enter Spare Part ID: ")
			fmt.Scan(&newPart.ID)
			fmt.Print("Enter Spare Part Name: ")
			fmt.Scan(&newPart.Name)
			fmt.Print("Enter Spare Part Price: ")
			fmt.Scan(&newPart.Price)
			fmt.Print("Enter Spare Part Stock: ")
			fmt.Scan(&newPart.Stock)
			if AddSparePart(newPart) {
				fmt.Println("Spare Part added successfully.")
			} else {
				fmt.Println("Failed to add Spare Part. Storage full.")
			}
		case 3:
			var id int
			var updatedPart StructSparePart
			fmt.Print("Enter Spare Part ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Spare Part Name: ")
			fmt.Scan(&updatedPart.Name)
			fmt.Print("Enter New Spare Part Price: ")
			fmt.Scan(&updatedPart.Price)
			fmt.Print("Enter New Spare Part Stock: ")
			fmt.Scan(&updatedPart.Stock)
			updatedPart.ID = id
			if UpdateSparePart(id, updatedPart) {
				fmt.Println("Spare Part updated successfully.")
			} else {
				fmt.Println("Spare Part not found.")
			}
		case 4:
			var id int
			fmt.Print("Enter Spare Part ID to delete: ")
			fmt.Scan(&id)
			if DeleteSparePart(id) {
				fmt.Println("Spare Part deleted successfully.")
			} else {
				fmt.Println("Spare Part not found.")
			}
		case 5:
			var order string
			fmt.Print("Enter order (ascending/descending): ")
			fmt.Scan(&order)
			SortSparePartsByUsage(ArrSpareParts, order)
		case 6:
			return
		default:
			fmt.Println("Invalid input.")
		}
	}
}

func customerMenu() {
	var choice int
	for {
		fmt.Println("\nCustomer Menu:")
		fmt.Println("1. Display Customer Data")
		fmt.Println("2. Add Customer Data")
		fmt.Println("3. Update Customer Data")
		fmt.Println("4. Delete Customer Data")
		fmt.Println("5. Search Customer Data")
		fmt.Println("6. Sort Customer Data")
		fmt.Println("7. Back to Main Menu")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayCustomers()
		case 2:
			var newCustomer StructCustomer
			fmt.Print("Enter Customer ID: ")
			fmt.Scan(&newCustomer.ID)
			fmt.Print("Enter Customer Name: ")
			fmt.Scan(&newCustomer.Name)
			if AddCustomer(newCustomer) {
				fmt.Println("Customer added successfully.")
			} else {
				fmt.Println("Failed to add customer. Storage full.")
			}
		case 3:
			var id int
			var updatedCustomer StructCustomer
			fmt.Print("Enter Customer ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Customer Name: ")
			fmt.Scan(&updatedCustomer.Name)
			updatedCustomer.ID = id
			if UpdateCustomer(id, updatedCustomer) {
				fmt.Println("Customer updated successfully.")
			} else {
				fmt.Println("Customer not found.")
			}
		case 4:
			var id int
			fmt.Print("Enter Customer ID to delete: ")
			fmt.Scan(&id)
			if DeleteCustomer(id) {
				fmt.Println("Customer deleted successfully.")
			} else {
				fmt.Println("Customer not found.")
			}
		case 5:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func transactionMenu() {
	var choice int
	for {
		fmt.Println("\nTransaction Menu:")
		fmt.Println("1. Display Transaction Data")
		fmt.Println("2. Add Transaction Data")
		fmt.Println("3. Update Transaction Data")
		fmt.Println("4. Delete Transaction Data")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayTransactions()
		case 2:
			var newTransaction StructTransaction
			fmt.Print("Enter Transaction ID: ")
			fmt.Scan(&newTransaction.ID)
			fmt.Print("Enter Customer ID: ")
			fmt.Scan(&newTransaction.CustomerID)
			fmt.Print("Enter Service Fee: ")
			fmt.Scan(&newTransaction.ServiceFee)
			newTransaction.Date = time.Now()
			fmt.Println("Enter Spare Part Usage (Enter ID and Quantity, 0 to stop):")
			for {
				var sparePartID, quantity int
				fmt.Print("Enter Spare Part ID: ")
				fmt.Scan(&sparePartID)
				if sparePartID == 0 {
					break
				}
				fmt.Print("Enter Quantity: ")
				fmt.Scan(&quantity)
				newTransaction.SpareParts[sparePartID] = quantity
			}
			newTransaction.TotalPrice = CalculateServiceFee(newTransaction.ServiceFee, newTransaction.SpareParts, ArrSpareParts) // <- Fixed function call with correct params
			if AddTransaction(newTransaction) {
				fmt.Println("Transaction added successfully.")
			} else {
				fmt.Println("Failed to add transaction. Storage full.")
			}
		case 3:
			var id int
			var updatedTransaction StructTransaction
			fmt.Print("Enter Transaction ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Customer ID: ")
			fmt.Scan(&updatedTransaction.CustomerID)
			fmt.Print("Enter New Service Fee: ")
			fmt.Scan(&updatedTransaction.ServiceFee)
			updatedTransaction.Date = time.Now()
			fmt.Println("Enter New Spare Part Usage (Enter ID and Quantity, 0 to stop):")
			for {
				var sparePartID, quantity int
				fmt.Print("Enter Spare Part ID: ")
				fmt.Scan(&sparePartID)
				if sparePartID == 0 {
					break
				}
				fmt.Print("Enter Quantity: ")
				fmt.Scan(&quantity)
				updatedTransaction.SpareParts[sparePartID] = quantity
			}
			updatedTransaction.TotalPrice = CalculateServiceFee(updatedTransaction.ServiceFee, updatedTransaction.SpareParts, ArrSpareParts) // <- Fixed function call with correct params
			if UpdateTransaction(id, updatedTransaction) {
				fmt.Println("Transaction updated successfully.")
			} else {
				fmt.Println("Transaction not found.")
			}
		case 4:
			var id int
			fmt.Print("Enter Transaction ID to delete: ")
			fmt.Scan(&id)
			if DeleteTransaction(id) {
				fmt.Println("Transaction deleted successfully.")
			} else {
				fmt.Println("Transaction not found.")
			}
		case 5:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func main() {
	ArrSpareParts[0] = StructSparePart{ID: 1, Name: "Oli Mesin", Price: 75000, Stock: 20, UsageCount: 0}
	ArrSpareParts[1] = StructSparePart{ID: 2, Name: "Kampas Rem", Price: 50000, Stock: 15, UsageCount: 0}
	ArrSpareParts[2] = StructSparePart{ID: 3, Name: "Filter Udara", Price: 30000, Stock: 10, UsageCount: 0}
	
	ArrCustomers[0] = StructCustomer{ID: 1, Name: "Sinta"}
	ArrCustomers[1] = StructCustomer{ID: 2, Name: "Nisa"}
	ArrCustomers[2] = StructCustomer{ID: 3, Name: "Najwa"}
	ArrCustomers[3] = StructCustomer{ID: 4, Name: "Wahyu"}
	
	ArrTransactions[0] = StructTransaction{
		ID:         1,
		CustomerID: 1,
		Date:       time.Now().AddDate(0, -1, 0),
		ServiceFee: 50000,
		SpareParts: [MAX_SPAREPARTS]int{1, 0, 2},
		TotalPrice: 155000,
	}
	ArrTransactions[1] = StructTransaction{
		ID:         2,
		CustomerID: 2,
		Date:       time.Now().AddDate(0, -2, 0),
		ServiceFee: 60000,
		SpareParts: [MAX_SPAREPARTS]int{0, 1, 1},
		TotalPrice: 140000,
	}
	
	var choice int
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Spare Part Data")
		fmt.Println("2. Customer Data")
		fmt.Println("3. Transaction Data")
		fmt.Println("4. Calculate Service Fee")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			sparePartMenu()
		case 2:
			customerMenu()
		case 3:
			transactionMenu()
		case 4:
			fmt.Print("Calculate Service Fee\n")
			var serviceFee float64
			fmt.Print("Input service fee: ")
			fmt.Scan(&serviceFee)

			var spareParts [MAX_SPAREPARTS]int 
			fmt.Println("Input spare parts used (quantity for each spare part):")
			for i := 0; i < MAX_SPAREPARTS; i++ {
				fmt.Printf("Spare Part %d: ", i+1)
				fmt.Scan(&spareParts[i])
			}

			totalPrice := CalculateServiceFee(serviceFee, spareParts, ArrSpareParts)
			fmt.Printf("Total Service Fee: %.2f\n", totalPrice)
		case 5:
			fmt.Println("See you next time!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}