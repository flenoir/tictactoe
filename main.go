package main

import "fmt"

// creation de la grille de jeu
var Grid = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var firstUserInput string
var secondUserInput string
var turn uint = 0

func main() {
	fmt.Printf("Tic Tac Toe Game\n")
	printGrid()
	var firstUserSelection = getFirstUserInput()
	modifGrid(firstUserSelection)
	printGrid()
	turn++
	fmt.Printf("%d\n", uint64(turn))
	var secondUserSelection = getSecondUserInput()
	modifGridSecondUser(secondUserSelection)
	printGrid()
	turn++
	fmt.Printf("%d\n", uint64(turn))
}

// affichage de la grille de jeu
func printGrid() {
	fmt.Printf("%v  %v  %v\n", Grid[0], Grid[1], Grid[2])
	fmt.Printf("%v  %v  %v\n", Grid[3], Grid[4], Grid[5])
	fmt.Printf("%v  %v  %v\n", Grid[6], Grid[7], Grid[8])
}

// demande de la selection du premier utilisateur
func getFirstUserInput() string {
	if turn%2 == 0 {
		fmt.Printf("Odd\n")
	} else {
		fmt.Printf("Even\n")
	}
	fmt.Println("User 1, your turn. Choose a place on the grid using it's number : ")
	fmt.Scan(&firstUserInput)
	fmt.Printf("You selected %s\n", firstUserInput)
	return firstUserInput
}

// modification avec X de la grille an fonction de l'input user 1
func modifGrid(firstUserInput string) {
	for index, gridCase := range Grid {
		if gridCase == firstUserInput {
			Grid[index] = "X"
		}

	}
}

// demande de la selection du second utilisateur
func getSecondUserInput() string {
	if turn%2 == 0 {
		fmt.Printf("Odd\n")
	} else {
		fmt.Printf("Even\n")
	}
	fmt.Println("User , your turn. Choose a place on the grid using it's number : ")
	fmt.Scan(&secondUserInput)
	fmt.Printf("You selected n°%s\n", secondUserInput)
	return secondUserInput
}

// modification avec O de la grille an fonction de l'input user 1
func modifGridSecondUser(secondUserInput string) {
	for index, gridCase := range Grid {
		if gridCase == secondUserInput {
			Grid[index] = "O"
		}

	}
}

// verification si gagné
// verification si entrée valide
// ajout d'une notion de tour avec une alternance de jeu si la valeur est paire ou impaire
