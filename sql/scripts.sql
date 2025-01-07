CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(20) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

CREATE TABLE seguidores(
    usuario_id INT NOT NULL,
    seguidor_id INT NOT NULL,

    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE

    PRIMARY_KEY(usuario_id, seguidor_id)
);