<h1 align="center">
    Maratona Full Cycle 3.0 by <img width="120" height="40" src="https://github.com/trainningjava/Maratona-Full-Cycle-3.0/blob/master/public/assets/img/curso1.png">
</h1>

<img src="https://github.com/trainningjava/Maratona-Full-Cycle-3.0/blob/master/public/assets/img/buffalo.png" alt="buffalo" width="50" height="50"> <h1>Hello Full Cycle utilizando Buffalo</h1>


<h2>Informações do desafio</h2>


Esse segundo desafio consiste em criar uma página web com o conteúdo "Hello Full Cycle" utilizando a linguagem Golang e o framework Buffalo.
<br>
Quando o usuário acessar o projeto no endpoint /hello, ele deverá ver a mensagem "Hello Full Cycle".
<br>
Para fazer a entrega do desafio, gere uma imagem docker de sua aplicação funcionando e a disponibilize em sua conta no Docker Hub.
<br><br>

Quando alguém executar: <b>docker run -p 3000:3000 <i>seuuser/suaimagem</i></b>, a aplicação deverá estar disponível na porta 3000.
<br><br>

<h2>Uso</h2>

docker pull  "Nome da Imagem"
<br>
Executa o download da imagem
<br><br>

docker images<br>
Lista as imagens 
<br>
Parametros:
<ul>
  <li>REPOSITORY - repositório</li>
  <li>TAG - a tag utilizada no repositório (é setado pelo mantenedor)</li>
  <li>IMAGE ID - o ID da Imagem</li>
  <li>CREATED - quando essa imagem foi criada</li>
  <li>SIZE - tamanho dessa imagem</li>
</ul>
<br><br>
docker docker run -p 3000:3000 "IMAGE ID"<br>
Roda a aplicação dentro da imagem<br><br>

<h2>Resultado</h2>
<br>
Abra o browser no endereço <b>http://localhost:3000/hello</b>
<br>
Será mostrado a seguinte mensagem no browser:
<br>
<h1><b>Hello Full Cycle</b></h1><br>

