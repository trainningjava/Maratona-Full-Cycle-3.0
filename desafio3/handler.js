'use strict';

function soma(params) {
  if (!params.a || !params.b) {
    return {"resultado": 0, "mensagem": "Parametro invalido"};
  }
  return {"resultado": (params.a * 1) + (params.b * 1), "mensagem":"desafio3" };
}

exports.soma = soma;
