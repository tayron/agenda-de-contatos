package util

import (
	"os"
	"path/filepath"
)

// ObterCaminhoDiretorioAplicacao -
func ObterCaminhoDiretorioAplicacao() string {
	ambienteAplicacao := os.Getenv("AMBIENTE")

	if ambienteAplicacao == "desenvolvimento" {
		return ObterCaminhoDiretorioAplicacaoWeb()
	}

	return ObterCaminhoDiretorioAplicacaoLinux()
}

// ObterCaminhoDiretorioAplicacao -
func ObterCaminhoDiretorioAplicacaoWeb() string {
	caminho, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return caminho
}

// ObterCaminhoDiretorioAplicacaoLinux
func ObterCaminhoDiretorioAplicacaoLinux() string {
	caminho, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return caminho
}
