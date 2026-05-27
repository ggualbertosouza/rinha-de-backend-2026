
# Load dataset
Nessa etapa estou lendo o arquivo, descompactando e por fim fazendo decode.

**Primeiros testes**
```
2026/05/26 22:23:48 decompress gz took: 27.38µs
2026/05/26 22:23:55 decode took: 7.24995921s
2026/05/26 22:23:55 dataset loaded: 3000000 references
```

---

3. Endpoint `/fraud-score`
    3.1. Recebe json
    3.2. Converte em vetor
    3.3. Busca
    3.4. Calcula score
    3.5. Responde