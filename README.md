No repositório abaixo tem um guia de padrão de organização do projeto que a comunidade GO utiliza como boas práticas da linguagem.

```
https://github.com/golang-standards/project-layout
```

A pasta internal, é somente para a sua aplicação e não deve ser compartilhada com outros projetos. Essa pasta não pode ficar disponível para outras pessoas importarem. Tudo que você não quer que seja importado por outras aplicações, deve ficar dentro dessa pasta.

A pasta pkg é o oposto da internal, tudo que você quer que seja importado por outras aplicações, deve ficar dentro dessa pasta.

A pasta cmd é onde fica o seu projeto, o build, o main.go

A pasta configs é onde vai ficar as configurações do seu projeto. Padrões do projeto, configurações de banco de dados, templates, boot da aplicação, etc.

A pasta test armazena os testes da aplicação, stubs, documentações do teste.

A pasta api tem documentações da API, swagger, openapi, etc.

O foco dessa pasta é estudar a estrutura de API com GO, a parte de domain e infraestrutura não será abordada nessa etapa.