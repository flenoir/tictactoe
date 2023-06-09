package main

import "fmt"

// creation de la grille de jeu
var Grid = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

// variable de l'entrée utilisateur
var UserInput string

// variable de tour de jeu
var turn uint = 0

// definition des symboles associés aux joueurs
var playerSigns = []string{"X", "O"}

// fonction principale
func main() {
	fmt.Printf("Tic Tac Toe Game\n")
	printGrid()
	for checkWin() {
		var userSelection string = getUserInput()
		// verification si l'entrée valide
		if checkIfContainsUserInput(userSelection) {
			modifGrid(userSelection)
			printGrid()
			turn++ //incrémentation du tour
		} else {
			fmt.Printf("Your selection is invalid, please retry.\n")
			var userSelection string = getUserInput()
			modifGrid(userSelection)
			printGrid()
			turn++
		}
	}
}

// affichage de la grille de jeu
func printGrid() {
	fmt.Printf("%v  %v  %v\n", Grid[0], Grid[1], Grid[2])
	fmt.Printf("%v  %v  %v\n", Grid[3], Grid[4], Grid[5])
	fmt.Printf("%v  %v  %v\n", Grid[6], Grid[7], Grid[8])
}

// demande de la selection  de l'utilisateur
func getUserInput() string {
	if turn%2 == 0 {
		fmt.Println("Player 1, your turn. Choose a place on the grid using it's number : ")
		fmt.Scan(&UserInput)
		fmt.Printf("You selected %s\n", UserInput)
		return UserInput
	} else {
		fmt.Println("Player 2, your turn. Choose a place on the grid using it's number : ")
		fmt.Scan(&UserInput)
		fmt.Printf("You selected %s\n", UserInput)
		return UserInput
	}
}

// modification avec X ou O de la grille en fonction de l'input user
func modifGrid(UserInput string) {
	for index, gridCase := range Grid {
		if gridCase == UserInput && turn%2 == 0 {
			Grid[index] = "X"
		} else if gridCase == UserInput && turn%2 != 0 {
			Grid[index] = "O"
		}

	}
}

// verification si un joueur a gagné
func checkWin() bool {
	for _, value := range playerSigns {
		//verifie les lignes puis les colonnes et les diagonales
		if Grid[0] == value && Grid[1] == value && Grid[2] == value ||
			Grid[3] == value && Grid[4] == value && Grid[5] == value ||
			Grid[6] == value && Grid[7] == value && Grid[8] == value {
			fmt.Printf("C'est gagné !\n")
			return false
		} else if Grid[0] == value && Grid[3] == value && Grid[6] == value ||
			Grid[1] == value && Grid[4] == value && Grid[7] == value ||
			Grid[2] == value && Grid[5] == value && Grid[8] == value {
			fmt.Printf("C'est gagné !\n")
			return false
		} else if Grid[0] == value && Grid[4] == value && Grid[8] == value ||
			Grid[2] == value && Grid[4] == value && Grid[6] == value {
			fmt.Printf("C'est gagné !\n")
			return false
		}
	}
	return true
}

// fonction de verification de la présence de l'entrée utilisateur dans la valeurs possibles de la grille
func checkIfContainsUserInput(userChoice string) bool {
	for _, v := range Grid {
		fmt.Printf("The value is %v\n", v)
		if userChoice == v {
			return true
		}
	}
	return false
}
