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

// Data global berupa array statis
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

// Wahyu Bagus S
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
func SortSparePartsByUsage(parts []StructSparePart, order string) []StructSparePart {
	sortedParts := make([]StructSparePart, len(parts))
	copy(sortedParts, parts)

	for i := 0; i < len(sortedParts)-1; i++ {
		selectedIdx := i
		for j := i + 1; j < len(sortedParts); j++ {
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
	// Implementasi insertion sort langsung di array input
	for i := 1; i < len(custs); i++ {
		key := custs[i]
		j := i - 1

		// Geser elemen selama valid dan sesuai urutan
		for j >= 0 && ((order == "ascending" && custs[j].Name > key.Name) ||
			(order == "descending" && custs[j].Name < key.Name)) {
			custs[j+1] = custs[j]
			j--
		}
		custs[j+1] = key
	}

	// Salin hasil ke array hasil
	var sortedCusts [MAX_CUSTOMERS]StructCustomer
	copy(sortedCusts[:], custs[:])

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
			sortedParts := SortSparePartsByUsage(ArrSpareParts[:], order)
			fmt.Println("Sorted Spare Parts by usage:")
			for _, part := range sortedParts {
				if part.ID != 0 {
					fmt.Printf("ID: %03d, Name: %s, UsageCount: %d\n", part.ID, part.Name, part.UsageCount)
				}
			}
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
			var search int
			fmt.Println("\nSearching Customer: ")
			fmt.Println("1. Search by Serviced period")
			fmt.Println("2. Search by Spare Part Usage")
			fmt.Println("3. Back to Main Menu")
			fmt.Print("Enter your choice: ")
			fmt.Scan(&search)

			switch search {
			case 1:
				var startDateStr, endDateStr string
				fmt.Print("Enter Start Date (YYYY-MM-DD): ")
				fmt.Scan(&startDateStr)
				fmt.Print("Enter End Date (YYYY-MM-DD): ")
				fmt.Scan(&endDateStr)
		
				startDate, _ := time.Parse("2006-01-02", startDateStr)
				endDate, _ := time.Parse("2006-01-02", endDateStr)
		
				result := FindCustomersByServicePeriod(startDate, endDate)
				fmt.Println("Customers serviced within the period: ")
				for _, customer := range result {
					if customer.ID != 0 {
						fmt.Printf("ID: %03d, Name: %s\n", customer.ID, customer.Name)
					}
				}
			case 2:
				var sparePartID int
				fmt.Print("Enter Spare Part ID: ")
				fmt.Scan(&sparePartID)

				result := FindCustomersBySparePart(sparePartID)
				fmt.Println("Customers who used the spare part:")
				for _, customer := range result {
					if customer.ID != 0 {
						fmt.Printf("ID: %03d, Name: %s\n", customer.ID, customer.Name)
					}
				}
			case 3:
				return
			default:
				fmt.Println("Invalid choice, try again.")			
			}
		case 6:
			var order string
			fmt.Print("Enter order (ascending/descending): ")
			fmt.Scan(&order)
			sortedCusts := SortCustomersByName(ArrCustomers, order)
			fmt.Println("Sorted Customers (alphabetically):")
			for _, cust := range sortedCusts {
				if cust.ID != 0 {
					fmt.Printf("ID: %03d, Name: %s\n", cust.ID, cust.Name)
				}
			}
		case 7:
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

			newTransaction.TotalPrice = CalculateServiceFee(newTransaction.ServiceFee, newTransaction.SpareParts, ArrSpareParts) 
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
			updatedTransaction.TotalPrice = CalculateServiceFee(updatedTransaction.ServiceFee, updatedTransaction.SpareParts, ArrSpareParts)
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
	ArrSpareParts[0] = StructSparePart{ID: 1, Name: "Oli Mesin", Price: 75000, Stock: 20, UsageCount: 5}
	ArrSpareParts[1] = StructSparePart{ID: 2, Name: "Kampas Rem", Price: 50000, Stock: 15, UsageCount: 3}
	ArrSpareParts[2] = StructSparePart{ID: 3, Name: "Filter Udara", Price: 30000, Stock: 10, UsageCount: 7}
	ArrSpareParts[3] = StructSparePart{ID: 4, Name: "Roda Motor", Price: 120000, Stock: 15, UsageCount: 3}

	ArrCustomers = [MAX_CUSTOMERS]StructCustomer{
		{ID: 1, Name: "Sinta"},
		{ID: 2, Name: "Nisa"},
		{ID: 3, Name: "Najwa"},
		{ID: 4, Name: "Wahyu"},
	}

	ArrTransactions[0] = StructTransaction{
		ID:         1,
		CustomerID: 1,
		Date: 		time.Date(2025, time.January, 3, 12, 0, 0, 0, time.UTC),
		ServiceFee: 50000,
		SpareParts: [MAX_SPAREPARTS]int{1, 0, 2},
		TotalPrice: 155000,
	}
	ArrTransactions[1] = StructTransaction{
		ID: 		2,
		CustomerID: 2,
		Date: 		time.Date(2024, time.February, 20, 12, 0, 0, 0, time.UTC),
		ServiceFee: 60000,
		SpareParts: [MAX_SPAREPARTS]int{2, 0, 3},
		TotalPrice: 140000,
	}

	ArrTransactions[2] = StructTransaction{
		ID:         3,
		CustomerID: 3,
		Date: 		time.Date(2024, time.November, 3, 12, 0, 0, 0, time.UTC),
		ServiceFee: 50000,
		SpareParts: [MAX_SPAREPARTS]int{1, 0, 2},
		TotalPrice: 155000,
	}
	ArrTransactions[3] = StructTransaction{
		ID: 		4,
		CustomerID: 4,
		Date: 		time.Date(2025, time.February, 2, 12, 0, 0, 0, time.UTC),
		ServiceFee: 60000,
		SpareParts: [MAX_SPAREPARTS]int{3, 2, 4},
		TotalPrice: 200000,
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
			fmt.Println("Calculate Service Fee")
			
			var serviceFee float64
			fmt.Print("Input service fee: ")
			fmt.Scan(&serviceFee)
		
			var spareParts [MAX_SPAREPARTS]int
			fmt.Println("Input spare parts used (enter -1 to finish):")
			for {
				var sparePartID int
				fmt.Print("Enter spare part ID: ")
				fmt.Scan(&sparePartID)
		
				if sparePartID == -1 {
					break
				}
		
				// Search spare part dengan ID
				var sparePartIndex int = -1
				for i := 0; i < MAX_SPAREPARTS; i++ {
					if ArrSpareParts[i].ID == sparePartID {
						sparePartIndex = i
						break
					}
				}
		
				if sparePartIndex != -1 {
					var quantity int
					fmt.Print("Enter quantity for Spare Part ", sparePartID, ": ")
					fmt.Scan(&quantity)
		
					if quantity >= 0 {
						spareParts[sparePartIndex] += quantity
					} else {
						fmt.Println("Invalid quantity. Please enter a non-negative number.")
					}
				} else {
					fmt.Println("Invalid spare part ID. Please try again.")
				}
			}
		
			// Calculate total service fee
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