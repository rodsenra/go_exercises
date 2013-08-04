
#textfiles downloader

Construa um programa que faça download de arquivos a partir de uma das URLs listadas no <http://www.textfiles.com/> (por exemplo, <http://www.textfiles.com/programming/)>. Faça sua solução explorar o suporte a concorrência e paralelismo da linguagem.

**Sugestão:** use um conjunto de workers, uma solução semelhante à ideia dos elevadores e das pessoas.

* para requisições HTTP, utilize a função [Get](http://golang.org/pkg/net/http/#Get), do pacote [net/http](http://golang.org/pkg/net/http/).
* para definir flags na linha de comando, utilize o pacote [flag](http://golang.org/pkg/flag/).

Existem duas formas de definir uma flag: usando uma referência para uma variável que já existia antes ou criando a variável na declaração da flag:

```go
// cria a variável antes e referencia na flag
// Quando o usuário passar -name Francisco, o valor da variável name será "Francisco".
var name string
flag.StringVar(&name, "name", "", "your name")

// age é um ponteiro para inteiro, para acessar seu valor, é necessário usar o operador *.
// 30 é o valor padrão dessa flag
age := flag.Int("age", 30, "your age")

// acessando valores das flags
fmt.Println(name, *age)
```

Para mais detalhes, veja a explicação no [Overview do pacote flag](http://golang.org/pkg/flag/#pkg-overview).

O seu programa deve ser invocável a partir da linha de comando e funcionar da seguinte forma:

```
% ./download -d <diretório-de-destino> -u <url> -w <workers>
```

Exemplo:

```
% ./download -d /Users/f/arquivos -u http://www.textfiles.com/programming/ -w 40
```
