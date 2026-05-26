Entendendo o projeto

1. Server HTTP
    1.1. Rota `/ready`
    1.2. Rota `/fraud-score`

2. Load dataset on init
    1.1. Lê arquivos ./resource
    1.2. Parser
    1.3. Otimiza estrutura
    1.4. Manter em memória
    1.5. Disponibiliza a API

!! Endpoint de `/ready` só vai informar que a API está pronta após informações já estarem em memória

Enquanto dataset não estiver pronto -> ready = false/503
Se estiver pronto -> ready = true/200

3. Endpoint `/fraud-score`
    3.1. Recebe json
    3.2. Converte em vetor
    3.3. Busca
    3.4. Calcula score
    3.5. Responde