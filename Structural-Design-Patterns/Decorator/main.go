package main

import "fmt"

///////////// Compoennt Interface ////////////////

type pizza interface {
	getPrice() int
}

/////////////// Implementations //////////////////

///////////////// Components /////////////////////

// peppy paneer
type peppyPaneer struct{}

func (p *peppyPaneer) getPrice() int {
	return 20
}

// veggie mania
type veggieMania struct{}

func (p *veggieMania) getPrice() int {
	return 15
}

/////////////////// Decorators ///////////////////

// Same interface as component plus embedding of the component!

// tomato
type tomatoTopping struct {
	pizza pizza // embed pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// cheese
type cheeseTopping struct {
	pizza pizza // embed pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

//////////////// Main ////////////////////////////

func main() {
	///////////// Veggie Mania Pizza ///////////////
	veggiePizza := &veggieMania{}
	veggiePizzaWithCheese := &cheeseTopping{pizza: veggiePizza}
	veggiePizzaWithCheeseAndTomato := &tomatoTopping{pizza: veggiePizzaWithCheese}

	fmt.Printf("Price of Veggie Mania pizza with Cheese Topping is %d.\n", veggiePizzaWithCheese.getPrice())
	fmt.Printf("Price of Veggie Mania pizza with Tomato and Cheese Topping is %d.\n", veggiePizzaWithCheeseAndTomato.getPrice())

	///////////// Peppy Paneer Pizza ///////////////
	peppyPaneerPizza := &peppyPaneer{}
	peppyPaneerPizzaWithCheese := &cheeseTopping{pizza: peppyPaneerPizza}
	peppyPaneerPizzaWithCheeseAndTomato := &tomatoTopping{pizza: peppyPaneerPizzaWithCheese}

	fmt.Printf("Price of Peppy Paneer with Cheese Topping is %d.\n", peppyPaneerPizzaWithCheese.getPrice())
	fmt.Printf("Price of Peppy Paneer with Tomato and Cheese Topping is %d.\n", peppyPaneerPizzaWithCheeseAndTomato.getPrice())
}
