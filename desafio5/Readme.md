<h1 align="center">
    Maratona Full Cycle 3.0 by <img width="120" height="40" src="https://github.com/trainningjava/Maratona-Full-Cycle-3.0/blob/master/public/assets/img/curso1.png">
</h1>

<img src="https://github.com/trainningjava/Maratona-Full-Cycle-3.0/blob/master/public/assets/img/nest.png" alt="Nest.js" width="50" height="50"> <h1>Endpoint com a programação das lives com Nest.js e SQLite</h1>


<h2>Informações do desafio</h2>

Esse quinto desafio que fará você ter contato com um dos frameworks javascript que mais cresce no mercado. O Nest.js.
<br>
Nesse desafio você terá que criar, utilizando o Nest.js, uma simples aplicação que listará o nome de todas as lives da nossa Maratona Full Cycle 3.0, sendo que os dados 
<br>
devem vir diretamente de um banco de dados SQLite.
<br><br>

Quando alguém acessar a aplicação através do endpoint /maratona, pela porta 3000, a listagem deverá aparecer.
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
Abra o browser no endereço <b>http://localhost:3000/maratona</b>
<br>
Será mostrado a seguinte mensagem no browser:
<br>
<h1><b>Lista de aulas no formato json</b></h1><br>

