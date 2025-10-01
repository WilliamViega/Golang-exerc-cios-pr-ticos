package main 

func notaParaConceito (n float64) string {
	var nota = int(n) 
	swith nota {
	case 10:
		fallthrough 
	case 9: 
	return "A" 
	case 8, 7:
		return "B" 

	default: 
		return "Nota inválida" 
	}
}