package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
	}
	defer db.Close()

	rows, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao consultar o banco de dados!"))
		return
	}
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario

		if erro := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	// TESTANDO ESCREVER O BODY DO HTTP NUM ARQUIVO .TXT
	//arquivo, err := os.OpenFile("JsonEncoder.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	//if err != nil {
	//	fmt.Println("Erro ao abrir arquivo:", err)
	//	return
	//}
	//defer arquivo.Close() // Garante o fechamento do arquivo
	//
	//if erro := json.NewEncoder(arquivo).Encode(usuarios); erro != nil {
	//	w.Write([]byte("Erro ao converter usuário para JSON"))
	//	return
	//}

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuário para JSON!"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
