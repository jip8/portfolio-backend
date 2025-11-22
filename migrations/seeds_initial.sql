SET search_path TO portfolio;

-- ============================================================================
-- ABOUT TEXT
-- ============================================================================
INSERT INTO about_text (title, content) VALUES
('Desenvolvedor Full Stack, Desenvolvedor Backend, Especialista em Redes, Engenheiro de Software, Pesquisador em Controle de Acesso, Desenvolvedor Web, Desenvolvedor Mobile', 'Possuo sólida experiência em programação, com domínio nas linguagens C, Python e C#, além de atuação no desenvolvimento de APIs REST e microserviços utilizando Go e Python. Tenho conhecimento em bancos de dados como PostgreSQL e MySQL, e experiência com tecnologias como Redis, MinIO (S3) e ferramentas de versionamento com Git.

Durante a graduação em Tecnologia em Redes de Computadores, adquiri amplo conhecimento em redes, infraestrutura, configuração de dispositivos de roteamento e switches, servidores, uso de contêineres com Docker, além de familiaridade com diferentes protocolos de rede e metodologias ágeis.

Atualmente, atuo também com integração e orquestração de modelos de linguagem (LLMs) e agentes de IA, utilizando ferramentas como CrewAI e LangChain, com integrações em plataformas como OpenAI, Anthropic, Cohere, DeepSeek e Google.

Tenho experiência prática com Linux, virtualização com Proxmox, e um nível intermediário de inglês. Estou em constante evolução, buscando novos aprendizados e desafios na área de desenvolvimento e inovação tecnológica.');

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
'Atuei no desenvolvimento de ambientes virtuais interativos com o objetivo de reduzir riscos de acidentes e custos associados a práticas reais, utilizando óculos de realidade virtual, mais especificamente o Meta Quest 2.

Além disso, participei da integração do óculos Meta Quest 2 com o robô humanoide NAO V6, que possui mobilidade, articulações e câmeras. Esse trabalho possibilitou que o usuário do VR controlasse o robô e visualizasse a sua visão em tempo real.',
'2023-01-01', '2024-12-31', false),

('Universidade Federal de Santa Maria - UFSM', 'Desenvolvedor',
'Atuei no desenvolvimento de sistemas para o controle de acesso e gerenciamento dos ambientes do CTISM. Esse projeto também resultou no meu Trabalho de Conclusão de Curso, "Aperfeiçoando Sistemas de Controle de Acesso Utilizando Reconhecimento Facial", e também na publicação de um artigo na ERRC/WRseg 2024 em Rio Grande/RS.',
'2024-01-01', '2025-01-01', false),

('Nasverdes Tecnologia LTDA', 'Desenvolvedor Back-End',
'Atuo como desenvolvedor backend com foco em APIs RESTful e microserviços, utilizando principalmente Go e Python. Desenvolvo soluções voltadas para a integração e orquestração de agentes de inteligência artificial, empregando ferramentas como CrewAI e LangChain, além de realizar integrações com APIs da OpenAI, Anthropic, DeepSeek, Google e Cohere.

Também trabalho com processamento e manipulação de documentos (PDFs, DOCs, arquivos estruturados e não estruturados), alimentando agentes de IA com dados validados, extraídos e organizados. Utilizo Redis para comunicação assíncrona entre serviços, PostgreSQL para dados relacionais, Git para versionamento e S3 para armazenamento de documentos.',
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