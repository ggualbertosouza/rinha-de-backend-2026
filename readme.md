
# Load dataset
Nessa etapa estou lendo o arquivo, descompactando e por fim fazendo decode.

**Primeiros testes**
```
2026/05/26 22:27:44 server starting on 8080
2026/05/26 22:27:44 processing dataset
2026/05/26 22:27:44 open file took: 21.41µs
2026/05/26 22:27:44 decompress gz took: 27.61µs
2026/05/26 22:27:51 decode took: 7.270271159s
2026/05/26 22:27:51 total time took: 7.270358649s
2026/05/26 22:27:51 dataset loaded: 3000000 references
```

---

3. Endpoint `/fraud-score`
    3.1. Recebe json
    3.2. Converte em vetor
    3.3. Busca
    3.4. Calcula score
    3.5. Responde