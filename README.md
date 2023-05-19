# CLI Calculator in Golang

This is a simple command-line calculator application I made in Golang to practice the basics. It can perform the for basic arithmetic operations with integer numbers. The calculator provides a user-friendly menu for selecting the desired operation and accepts two integer inputs for performing calculations.

## Prerequisites
Before running the calculator application, ensure that you have the following requirements met:

- Go programming language installed ([Download Go](https://golang.org/dl/))

## Installation

1. Clone the repository to your local machine:

  ```
  git clone https://github.com/your-username/calculator.git
  ``` 
2. Change to the project repository:
  
  ```
  cd calculator
  ```
3. Install the required package:
  ```
  go get github.com/inancgumus/screen
  ```
## Usage

1. At a terminal, in the project repository, run the application:
  ```
  go run main.go
  ```
2. The menu will be displayed: 
  ```
  Choose a number for an operation:
  [1] - Addition
  [2] - Subtraction
  [3] - Multiplication
  [4] - Division
  [0] - Exit
  ```
4. Choose an operation by typing the number corresponding to the desired option and pressing ENTER.
5. Insert the two integers for the operation. After the second number is entered, the result will be displayed.
6. The menu will be displayed again, until you enter '0' to exit the application.

## Acknowledgments 

- [inancgumus/screen](https://github.com/inancgumus/screen) - Module used for clearing the screen.
