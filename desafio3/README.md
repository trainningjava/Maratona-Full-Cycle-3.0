<h1 align="center">
    Maratona Full Cycle 3.0 by <img width="120" height="40" src="https://github.com/trainningjava/Maratona-Full-Cycle-3.0/blob/master/public/assets/img/curso1.png">
</h1>

<img src="https://github.com/trainningjava/Maratona-Full-Cycle-3.0/blob/master/public/assets/img/serverless.png" alt="docker" width="50" height="50"> <h1>Somando com endpoint

## Informações do desafio

Esse terceiro desafio vai te introduzir ao mundo serverless!

E por conta disso você terá que gerar um endpoint no seguinte 
formato: /soma?a={numero}&b={numero}.

Quando alguém acessar através do método get um json 
deve ser retornado no formato:

`{"resultado":valor}`

Para realizar o desafio, fique na liberdade para escolher entre as linguagens: javascript, golang ou python. Também utilize o framework Serverless para realizar essa tarefa e utilize o cloud provider que achar mais conveniente.


## Minhas referências

* https://www.serverless.com/framework/docs/getting-started/
* https://www.serverless.com/examples/openwhisk-node-simple/
* https://github.com/serverless/examples/blob/master/openwhisk-node-simple/hello_world.js
* https://www.serverless.com/framework/docs/providers/openwhisk/cli-reference/invoke/
* https://www.serverless.com/framework/docs/providers/openwhisk/guide/credentials/

## Executando

### Autenticar na IBM Cloud

```bash
docker-compose run serveless ibmcloud login
docker-compose run serveless ibmcloud api
docker-compose run serveless ibmcloud wsk property get --auth
```

### Acessando funcion localmente

```bash
docker-compose run serveless sls invoke local -f soma --data '{"a": 1, "b": 2 }'
```

