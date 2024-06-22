CREATE DATABASE IF NOT EXISTS wallet;

USE wallet;

CREATE TABLE IF NOT EXISTS clients(id varchar(255) primary key, name varchar(255), email varchar(255), created_at date);

CREATE TABLE IF NOT EXISTS accounts(id varchar(255) primary key, client_id varchar(255), balance int, created_at date);

CREATE TABLE IF NOT EXISTS transactions(id varchar(255) primary key, account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);

INSERT INTO `accounts` (`id`, `client_id`, `balance`, `created_at`) VALUES
('16c34238-8283-40b4-b536-fe70be50120a', '7a7a2587-7010-45ea-bf65-42ce4229be23', 1900, '2024-06-22'),
('f9baed52-ff45-4126-838c-4699d4aa5852', 'c0b39f93-0918-46fb-a14c-0f9a530f69aa', 2100, '2024-06-22');

INSERT INTO `clients` (`id`, `name`, `email`, `created_at`) VALUES
('7a7a2587-7010-45ea-bf65-42ce4229be23', 'teste1', 'teste1@gmail.com', '2024-06-22'),
('c0b39f93-0918-46fb-a14c-0f9a530f69aa', 'teste2', 'teste2@gmail.com', '2024-06-22');

INSERT INTO `transactions` (`id`, `account_id_from`, `account_id_to`, `amount`, `created_at`) VALUES
('730cd8ee-88c1-47c3-baba-63da25c9dfb6', '16c34238-8283-40b4-b536-fe70be50120a', 'f9baed52-ff45-4126-838c-4699d4aa5852', 100, '2024-06-22');