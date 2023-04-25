package bonus

import "sort"

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	maiorTorre := 1
	midTorre := 1
	num_torres := len(barLengths)
	sort.Ints(barLengths)
	elemento := 0
	for elemento < (len(barLengths) - 1) {
		if barLengths[elemento] == barLengths[elemento+1] {
			midTorre += 1
			if midTorre > maiorTorre && midTorre > 1 {
				midTorre = midTorre
			}
			num_torres -= 1

		} else {
			midTorre = 1
		}
		elemento++
	}

	return maiorTorre, num_torres
}
