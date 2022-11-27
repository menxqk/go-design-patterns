package structural

import (
	"strings"
	"testing"
)

func TestPizza_GetIngredients(t *testing.T) {
	pizza := &Pizza{}
	pizzaResult, _ := pizza.GetIngredients()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient method of the pizza decorator it must return the text %s, not '%s", pizzaResult, expectedText)
	}
}

func TestMeat_GetIngredients(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.GetIngredients()
	if err == nil {
		t.Errorf("When calling GetIngredients on the meat decorator without an HasIngredients on its HasIngredients field must return an error, not a string with '%s'", meatResult)
	}
	meat = &Meat{&Pizza{}}
	meatResult, err = meat.GetIngredients()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(meatResult, "meat") {
		t.Errorf("When calling the GetIngredients of the meat decorator it must return a text with the word 'meat', not '%s'", meatResult)
	}
}
func TestOnion_GetIngredients(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.GetIngredients()
	if err == nil {
		t.Errorf("When calling GetIngredients on the onion decorator without an HasIngredients on its HasIngredients field must return an error, not a string with '%s'", onionResult)
	}
	onion = &Onion{&Pizza{}}
	onionResult, err = onion.GetIngredients()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("When calling the GetIngredients of the onion decorator it must return a text with the word 'onion', not '%s'", onionResult)
	}
}

func TestGetIngredients_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&Pizza{}}}
	pizzaResult, err := pizza.GetIngredients()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients: meat onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned string must contain the text '%s' but '%s' did not have it", expectedText, pizzaResult)
	}
}
