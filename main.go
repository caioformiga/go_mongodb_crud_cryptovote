package main

func main() {
	/*
		fmt.Printf("Iniciou chamada a ações de deleção...\n")
		// cria os parametros do filtro
		//filter := bson.M{"last_name": "formiga"}
		filter := bson.M{}
		controller.DeletarVariosRegistros(filter)
		fmt.Printf("Finalizou a chamada sem erros fatais!\n")
	*/

	/*
		fmt.Printf("Iniciou chamada a ações de busca...\n")
		// cria os parametros do filtro
		//filter := bson.M{"last_name": "campos"}

		// cria os parametros do filtro sem restrições
		filter := bson.M{}
		sliceObjEmployee := controller.BuscaMuitosRegistros(filter)
		fmt.Printf("sliceObjEmployee segunda call:%v!\n!\n", sliceObjEmployee)
		/*
				if sliceObjEmployee != nil {
					fmt.Printf("Foram localizados %d registros, sendo:\n", len(sliceObjEmployee))
					for index, value := range sliceObjEmployee {
						fmt.Printf("Posição [%d]: %+v\n", index, value)
					}
				}
				else {
					if len(sliceObjEmployee) == 0 {
						fmt.Printf("Não foi localizada NENHUM registro!\n")
					}
				}
				fmt.Printf("Finalizou a chamada sem erros fatais!\n")
	*/

	/*
		fmt.Printf("Iniciou chamada a ações de atualização...\n")
		// cria os parametros do filtro
		filter := bson.M{"last_name": "silva"}
		// cria os parametros de atualização
		newData := bson.M{
			"$set": bson.M{"age": -1},
		}
		objEmployee := controller.AtualizaUmRegistro(filter, newData)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Printf("NENHUM registro foi atualizado para o filtro %+v\n", filter)
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("Um registro foi localizado e atualizado %+v\n", objEmployee)
		}
		fmt.Printf("Finalizou a chamada sem erros fatais!\n")
	*/

	/*
		fmt.Printf("Iniciou chamada a ações de criação...\n")
		// criarUmRegistro()
		controller.CriarVariosRegistros(2)
		fmt.Printf("Finalizou a chamada sem erros fatais!\n")
	*/
}
