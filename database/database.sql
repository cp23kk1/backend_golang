CREATE TABLE `user` (
  `id` INT PRIMARY KEY,
  `email` VARCHAR,
  `role` VARCHAR,
  `display_name` VARCHAR,
  `active` BOOLEAN,
  `image` VARCHAR,
  `private_profile` VARCHAR
);

CREATE TABLE `score_board` (
  `id` INT PRIMARY KEY,
  `user_id` INT,
  `score` INT,
  `week` INT,
  `start_date` DATETIME,
  `end_date` DATETIME
);

CREATE TABLE `vocabulary` (
  `id` INT PRIMARY KEY,
  `word` VARCHAR,
  `meaning` VARCHAR,
  `pos` VARCHAR,
  `difficulty` ENUM
);

CREATE TABLE `sentence` (
  `id` INT PRIMARY KEY,
  `passage_id` INT,
  `sequence` INT,
  `text` VARCHAR,
  `meaning` VARCHAR
);

CREATE TABLE `vocabulary_related` (
  `vocabulary_id` INT,
  `sentence_id` INT,
  PRIMARY KEY (`vocabulary_id`, `sentence_id`)
);

CREATE TABLE `passage` (
  `id` INT PRIMARY KEY,
  `title` VARCHAR
);

CREATE TABLE `vocabulary_history` (
  `id` INT PRIMARY KEY,
  `user_id` INT,
  `vocabulary_id` INT,
  `game_id` VARCHAR,
  `correctness` BOOLEAN
);

CREATE TABLE `sentence_history` (
  `id` INT PRIMARY KEY,
  `user_id` INT,
  `sentence_id` INT,
  `game_id` VARCHAR,
  `correctness` BOOLEAN
);

CREATE TABLE `passage_history` (
  `id` INT PRIMARY KEY,
  `user_id` INT,
  `passage_id` INT,
  `game_id` VARCHAR,
  `correctness` BOOLEAN
);

ALTER TABLE `score_board` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `sentence` ADD FOREIGN KEY (`passage_id`) REFERENCES `passage` (`id`);

ALTER TABLE `vocabulary_related` ADD FOREIGN KEY (`vocabulary_id`) REFERENCES `vocabulary` (`id`);

ALTER TABLE `vocabulary_related` ADD FOREIGN KEY (`sentence_id`) REFERENCES `sentence` (`id`);

ALTER TABLE `vocabulary_history` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `vocabulary_history` ADD FOREIGN KEY (`vocabulary_id`) REFERENCES `vocabulary` (`id`);

ALTER TABLE `sentence_history` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `sentence_history` ADD FOREIGN KEY (`sentence_id`) REFERENCES `sentence` (`id`);

ALTER TABLE `passage_history` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `passage_history` ADD FOREIGN KEY (`passage_id`) REFERENCES `passage` (`id`);
