CREATE TABLE difficulties (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL
);

CREATE TABLE environments (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL
);

CREATE TABLE players (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE templates (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    playerId varchar(255) NOT NULL,
    templateName VARCHAR(255) NOT NULL,
    name varchar(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (playerId) REFERENCES players (id)
); 

CREATE TABLE test_results (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    waveSurvived SMALLINT UNSIGNED NOT NULL check(waveSurvived > 0),
moneyEarned DECIMAL(14, 2) NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE test_events (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    environmentId INTEGER NOT NULL,
    templateId INTEGER NOT NULL,
    difficultyId INTEGER NOT NULL,
    testResultId INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (environmentId) REFERENCES environments (id),
    FOREIGN KEY (templateId) REFERENCES templates (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);

CREATE TABLE catastrophes (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL
);

CREATE TABLE test_event_catastrophes (
    testEventId INTEGER NOT NULL,
    catastropheId INTEGER NOT NULL,
    wave SMALLINT UNSIGNED NOT NULL check(wave > 0),
    FOREIGN KEY (testEventId) REFERENCES test_events (id),
    FOREIGN KEY (catastropheId) REFERENCES catastrophes (id)
);


CREATE TABLE test_event_players (
    testEventId INTEGER NOT NULL,
    playerId varchar(255) NOT NULL,
    FOREIGN KEY (testEventId) REFERENCES test_events (id),
    FOREIGN KEY (playerId) REFERENCES players (id)
);



CREATE TABLE users(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);


CREATE TABLE player_test_results (
    playerId varchar(255) NOT NULL,
    testResultId INTEGER NOT NULL,
    waveDied SMALLINT UNSIGNED NOT NULL check(waveDied > 0),
    diedTo INTEGER NOT NULL,
    FOREIGN KEY (playerId) REFERENCES players (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id),
    FOREIGN KEY (diedTo) REFERENCES catastrophes (id)
);

