package main

type conta struct{
	nome string
	itens map[string]float64
	gorjeta float64
}

func novaConta(nome string) conta {
	c:=conta {
		nome: nome,
		itens: map[string]float64{}
		gorjeta: 0.0,
	}
	return c
	}

	func (c conta) formatacao() string{
		fs: "Detalhes da conta \n"
		var total float64 = 0

		for k,v := range itens {
		fs += fmt.sprintf("%v....R$ %0.2f \n",k+":", v)
		total += v
		}
		
		f+= fmt.sprintf("%v.....%0.2f", "total", total)
		return fs
	}	
	
