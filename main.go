package main

import (
	"context"
	"fmt"
	"reflect"
	"cloud.google.com/go/firestore"
)

func main() {
	// Inicializar o cliente do Firestore
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "project-id")
	if err != nil {
		fmt.Println(err)
	}
	// Nome da coleção a ser copiada
	collectionName := "collection-id"
	// Recuperar todos os documentos na coleção
	docs, err := client.Collection(collectionName).Documents(ctx).GetAll()

	if err != nil {
		fmt.Println("Erro: ", err)
	}
	fmt.Println(reflect.TypeOf(docs))
	defer client.Close()

	for _, doc := range docs {

		for field := range doc.Data() {

			fmt.Println(doc.Data()[field])

		}

	}

}
