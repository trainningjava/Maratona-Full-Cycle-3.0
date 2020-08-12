<img src="https://maratonafc3-assets.s3-sa-east-1.amazonaws.com/maratonafc3-assets/uploads/news/images/ac814a57a7144f60a214e4e09f6b0b0f.png" alt="docker" width="50" height="50"> <h1>Hello Full Cycle utilizando Golang e Docker</h1>


<h2>Informações do desafio</h2>

                            
O primeiro desafio dessa maratona consiste em criar um "Hello Full Cycle" utilizando a linguagem Golang.
<br>
Basicamente quando o arquivo compilado for executado, deverá ser exibido: Hello Full Cycle.
<br>
Se tudo estiver funcionando de forma adequada, gere uma imagem docker que quando executada deva rodar o programa criado em Golang.
<br><br>
Faça o push da imagem no Docker Hub e informe a url da imagem na área de entrega do desafio abaixo.


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
docker run "IMAGE ID"<br>
Roda a aplicação dentro da imagem<br><br>

<h2>Resultado</h2>
<br>
Será mostrado a mensagem no terminal
<br>
<h1><b>Hello Full Cycle</b></h1><br>

