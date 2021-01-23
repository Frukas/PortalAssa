# PortalAssa

Um sistema em microserviços responsável por receber, identificar e criar chamados de atendimento a partir deles em um frontEnd para gerenciamento e geração de relatórios.
Backend em GO, frontend em react/Typescript, banco de dados MongoDB e conexão de microservice através do gRPC.

Os arquivos tem os seguintes propositos:

MailWatcher - Verificação de recebimento de novos email em uma conta de email, tratando e enviando os dados para outros serviço.

WebConnector - Responsável por gerenciar os email, gerar chamados e se comunicar com o frontend via websocket.

loginpage - Um sistema frontend simples para conectar ao sistema Microsoft.


