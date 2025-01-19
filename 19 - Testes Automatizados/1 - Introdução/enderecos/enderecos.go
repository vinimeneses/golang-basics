package enderecos

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if strings.Contains(primeiraPalavraDoEndereco, tipo) {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavraDoEndereco)
	}

	return "Tipo inválido"
}
