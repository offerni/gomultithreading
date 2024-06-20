# Intructions 
Example run:
```
go run cmd/main.go 17055250
```
## Go Multithreading (PT-BR)
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:
- [x] https://brasilapi.com.br/api/cep/v1/{CEP}
- [x] http://viacep.com.br/ws/"{CEP}"/json/

Os requisitos para este desafio são:
- [x] Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- [x] O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- [x] Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Go Multithreading (EN)
In this challenge, you will need to use what we learned about Multithreading and APIs to fetch the fastest result from two different APIs.

The two requests will be made simultaneously to the following APIs:
- [x] https://brasilapi.com.br/api/cep/v1/{CEP}
- [x] http://viacep.com.br/ws/"{CEP}"/json/

The requirements for this challenge are:
- [x] Accept the API that delivers the fastest response and discard the slower response.
- [x] The result of the request should be displayed on the command line with the address details, as well as which API provided the response.
- [x] Limit the response time to 1 second. If this limit is exceeded, a timeout error should be displayed.