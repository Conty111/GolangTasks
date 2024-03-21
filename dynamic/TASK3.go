package dynamic

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	pizzas := map[int]int{s: cs, m: cm, l: cl}
	pizzas_coeff := make(map[int]float32, 3)
	best_pizza := []float32{float32(cs) / float32(s), float32(s)}
	for key, value := range pizzas {
		tmp := float32(value) / float32(key)
		if tmp < best_pizza[0] {
			best_pizza[0] = tmp
			best_pizza[1] = float32(key)
		}
		pizzas_coeff[key] = tmp
	}
	res := (x / int(best_pizza[1])) * pizzas[int(best_pizza[1])]
	if x%int(best_pizza[1]) == 0 {
		return res
	}
	tmp := res
	res += pizzas[int(best_pizza[1])]
	for square, cost := range pizzas {
		var pizza_count int
		tmp_sm_ostalos := x % int(best_pizza[1])
		for ; tmp_sm_ostalos > 0; pizza_count++ {
			tmp_sm_ostalos -= square
		}
		current_cost := tmp + cost*pizza_count
		if current_cost < res {
			res = current_cost
		}
	}
	return res
}
