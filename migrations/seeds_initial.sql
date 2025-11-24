SET search_path TO portfolio;

-- ============================================================================
-- ABOUT TEXT
-- ============================================================================
INSERT INTO about_text (title, content) VALUES
('Desenvolvedor Full Stack, Engenheiro de IA e Automações, Especialista em Redes, Engenheiro de Software, Entusiasta de Linux e servidores, Apaixonado por progamação', 'Muito prazer! Sou um apaixonado por programação e tecnologia, e atualmente meu foco principal é construir microsserviços e soluções inteligentes usando Python e Go. Gosto de desenvolver sistemas distribuídos que sejam leves, eficientes e escaláveis, integrando recursos de IA para automatizar processos, analisar dados e entregar respostas mais rápidas e inteligentes.

No dia a dia, trabalho criando APIs de alta performance, serviços orientados a eventos e integrações com modelos de IA. Utilizo ferramentas como Redis, PostgreSQL, MinIO e Docker para estruturar serviços modernos, bem organizados e preparados para lidar com alto volume de dados e requisições. Minha formação em Redes também me ajuda muito na parte de infraestrutura, segurança e operação em ambientes Linux.

Hoje, estou focado em desenvolver soluções modulares e práticas, unindo engenharia de software, automação e inteligência artificial. Meu objetivo é sempre criar sistemas que realmente resolvam problemas, sejam fáceis de evoluir e consigam acompanhar o ritmo das novas tecnologias.');

-- ============================================================================
-- SKILLS
-- ============================================================================
INSERT INTO skills (title, description) VALUES
-- Linguagens de Programação
('C', 'Linguagem de programação de baixo nível para sistemas e aplicações de alto desempenho'),
('Python', 'Linguagem versátil para desenvolvimento backend, scripts, IA e automação'),
('C#', 'Linguagem orientada a objetos utilizada principalmente em desenvolvimento de jogos com Unity'),
('Go', 'Linguagem moderna para desenvolvimento de APIs REST e microserviços de alta performance'),
('PHP', 'Linguagem para desenvolvimento web e sistemas'),
('C++', 'Extensão de C com recursos de programação orientada a objetos'),

-- Bancos de Dados
('PostgreSQL', 'Sistema de gerenciamento de banco de dados relacional open-source'),
('MySQL', 'Sistema de gerenciamento de banco de dados relacional'),
('Redis', 'Banco de dados em memória para cache e comunicação assíncrona entre serviços'),
('MongoDB', 'Sistema de gerenciamento de banco de dados orientado a documentos'),

-- Frameworks e Bibliotecas
('Flask', 'Framework web minimalista em Python'),
('Unity', 'Motor de desenvolvimento de jogos e experiências VR/AR'),
('OpenCV', 'Biblioteca de visão computacional para processamento de imagens'),
('CrewAI', 'Framework para orquestração de agentes de IA'),
('LangChain', 'Framework para desenvolvimento de aplicações com LLMs'),

-- Cloud e Armazenamento
('MinIO', 'Servidor de armazenamento de objetos compatível com S3'),
('S3', 'Serviço de armazenamento de objetos da Amazon Web Services'),

-- DevOps e Infraestrutura
('Docker', 'Plataforma de containerização de aplicações'),
('Git', 'Sistema de controle de versão distribuído'),
('Linux', 'Sistema operacional open-source'),
('Proxmox', 'Plataforma de virtualização open-source'),

-- Redes e Servidores
('Redes de Computadores', 'Conhecimento em infraestrutura, protocolos e configuração de redes'),
('Servidores', 'Configuração e gerenciamento de servidores'),
('Switches e Roteadores', 'Configuração de dispositivos de rede'),

-- IA e Machine Learning
('OpenAI API', 'Integração com modelos de linguagem da OpenAI'),
('Anthropic API', 'Integração com Claude e outros modelos da Anthropic'),
('Google AI', 'Integração com serviços de IA do Google'),
('DeepSeek', 'Integração com modelos de IA DeepSeek'),
('Cohere', 'Integração com plataforma Cohere de IA'),

-- Hardware e IoT
('ESP32', 'Microcontrolador para projetos de IoT'),
('Meta Quest 2', 'Óculos de realidade virtual'),
('Robô NAO V6', 'Robô humanoide programável'),

-- Ferramentas
('Blender', 'Software de modelagem 3D e animação'),
('PyHanko', 'Biblioteca Python para assinatura digital de PDFs'),

-- Metodologias
('Metodologias Ágeis', 'Práticas ágeis de desenvolvimento de software'),
('APIs REST', 'Desenvolvimento de interfaces de programação RESTful'),
('Microserviços', 'Arquitetura de software baseada em serviços independentes');

-- ============================================================================
-- EXPERIENCES
-- ============================================================================
INSERT INTO experiences (title, function, description, initial_date, end_date, actual) VALUES
('Universidade Federal de Santa Maria - UFSM', 'Desenvolvedor',
'Atuei no desenvolvimento de ambientes virtuais interativos aplicados ao ensino e treinamento, com foco na redução de riscos de acidentes e na diminuição de custos operacionais. Utilizei óculos de realidade virtual Meta Quest 2 para criar experiências imersivas que simulavam cenários de aprendizagem realistas.

Também participei da integração entre o Meta Quest 2 e o robô humanoide NAO V6, que possui mobilidade, articulações e sensores de visão. Essa integração permitiu que o usuário controlasse o robô diretamente a partir do VR, acompanhando em tempo real o que a câmera do NAO visualizava. O projeto envolveu conceitos de teleoperação, comunicação entre dispositivos, sincronização de movimentos e aplicação de técnicas de realidade mista.',
'2023-01-01', '2024-12-31', false),

('Universidade Federal de Santa Maria - UFSM', 'Desenvolvedor',
'Desenvolvi sistemas completos para controle de acesso e gerenciamento dos ambientes do CTISM, envolvendo reconhecimento facial, auditoria de acessos, integração com dispositivos físicos e gerenciamento de usuários. O projeto resultou no meu Trabalho de Conclusão de Curso, “Aperfeiçoando Sistemas de Controle de Acesso Utilizando Reconhecimento Facial”, onde explorei técnicas de visão computacional, bases de dados biométricas e otimização de performance para ambientes reais.

Além disso, contribuiu para uma publicação na ERRC/WRseg 2024, em Rio Grande/RS, apresentando novos métodos e soluções aplicadas ao controle de acesso físico. Esse trabalho uniu redes, segurança, desenvolvimento de software, automação e integração com dispositivos embarcados.',
'2024-01-01', '2025-01-01', false),

('Nasverdes Tecnologia LTDA', 'Desenvolvedor Back-End',
'Atuo como desenvolvedor backend com foco em arquitetura de microsserviços, APIs de alta performance e soluções inteligentes utilizando principalmente Go e Python. Desenvolvo integrações avançadas com agentes de inteligência artificial, utilizando tecnologias como CrewAI, LangChain e modelos das principais empresas do setor, como OpenAI, Anthropic, DeepSeek, Google e Cohere.

Crio pipelines completos de processamento e extração de informações a partir de documentos PDF, DOCX, arquivos estruturados e não estruturados, convertendo-os em dados limpos, validados e otimizados para consumo por agentes de IA. Esses pipelines envolvem OCR, parsing, saneamento de dados, chunking inteligente e enriquecimento de contexto.

Utilizo Redis para comunicação assíncrona entre serviços, filas distribuídas e troca de mensagens; PostgreSQL para armazenamento relacional e consultas eficientes; e S3/MinIO para gestão de arquivos e objetos. Trabalho ainda com conteinerização (Docker), versionamento com Git e práticas modernas de observabilidade e escalabilidade.

Minha atuação envolve desde a arquitetura dos serviços, definição das integrações, criação de soluções de IA assistida até a entrega de sistemas robustos, seguros e altamente escaláveis para uso em produção.',
'2025-01-01', NULL, true);



-- ============================================================================
-- SKILLS RELATIONS - EXPERIENCES
-- ============================================================================
-- Experiência 1: VR/AR Platform (2023-2024)
INSERT INTO skills_relations (parent_id, module, skill_id, revelance)
SELECT 1, 'experiences', id, 1 FROM skills WHERE title IN ('C#', 'Python', 'Unity', 'Blender', 'Meta Quest 2', 'Robô NAO V6', 'Redes de Computadores', 'Servidores');

-- Experiência 2: Controle de Acesso (2024-2025)
INSERT INTO skills_relations (parent_id, module, skill_id, revelance)
SELECT 2, 'experiences', id, 1 FROM skills WHERE title IN ('C++', 'Python', 'PHP', 'MySQL', 'OpenCV', 'Flask', 'ESP32', 'Servidores');

-- Experiência 3: Nasverdes (2025-atual)
INSERT INTO skills_relations (parent_id, module, skill_id, revelance)
SELECT 3, 'experiences', id, 1 FROM skills WHERE title IN ('Go', 'Python', 'CrewAI', 'LangChain', 'Redis', 'PostgreSQL', 'PyHanko', 'MinIO', 'S3', 'Docker', 'Git', 'OpenAI API', 'Anthropic API', 'DeepSeek', 'Google AI', 'Cohere', 'APIs REST', 'Microserviços');


-- ============================================================================
-- ARTICLES
-- ============================================================================
INSERT INTO articles (type, title, description, local, published_at, revelance) VALUES
('article', 'Aperfeiçoando Sistemas de Controle de Acesso Utilizando Reconhecimento Facial',
'Trabalho de Conclusão de Curso sobre sistemas de controle de acesso utilizando reconhecimento facial desenvolvido no CTISM/UFSM.',
'Universidade Federal de Santa Maria - UFSM', '2024-12-01', 1),

('article', 'Publicação na ERRC/WRseg 2024',
'Artigo científico sobre controle de acesso e gerenciamento de ambientes físicos apresentado na ERRC/WRseg 2024.',
'Rio Grande/RS', '2024-11-01', 2);