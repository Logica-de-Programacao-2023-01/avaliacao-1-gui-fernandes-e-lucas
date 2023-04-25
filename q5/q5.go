package q5

import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	x := strings.Split(s, "")
	i := 0
	for i < len(x) {
		x[i] = strings.ToLower(x[i])
		i++
	}

	newSlice := []string{}
	vogais := "aeiou"
	for _, j := range x {
		if !strings.ContainsAny(j, vogais) {
			newSlice = append(newSlice, j)
		}

	}
	i = 0
	newS := ""
	for i < len(newSlice) {
		newS += "."
		newS += newSlice[i]

		i++
	}

	return newS
}
