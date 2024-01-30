DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);

DROP TABLE IF EXISTS test_events;
CREATE TABLE test_events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
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

DROP TABLE IF EXISTS difficulties;
CREATE TABLE difficulties (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name varchar(255) NOT NULL,
);

DROP TABLE IF EXISTS environments;
CREATE TABLE environments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name varchar(255) NOT NULL,
);

DROP TABLE IF EXISTS templates;
CREATE TABLE templates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    playerId varchar(255) NOT NULL,
    templateName VARCHAR(255) NOT NULL,
    name varchar(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (playerId) REFERENCES players (id)
); 

DROP TABLE IF EXISTS test_event_catastrophes;
CREATE TABLE test_event_catastrophes (
    testEventId INTEGER NOT NULL,
    catastropheId INTEGER NOT NULL,
    wave UNSIGNED SMALLINT NOT NULL check(wave > 0),
    FOREIGN KEY (testEventId) REFERENCES test_events (id),
    FOREIGN KEY (catastropheId) REFERENCES catastrophes (id)
);

DROP TABLE IF EXISTS test_event_players;
CREATE TABLE test_event_players (
    testEventId INTEGER NOT NULL,
    playerId INTEGER NOT NULL,
    FOREIGN KEY (testEventId) REFERENCES test_events (id),
    FOREIGN KEY (playerId) REFERENCES players (id)
);

DROP TABLE IF EXISTS players;
CREATE TABLE players (
    id STRING PRIMARY KEY,
    name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS test_results;
CREATE TABLE test_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    waveSurvived UNSIGNED SMALLINT NOT NULL check(waveSurvived > 0),
    moneyEarned DECIMAL(14, 2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
);

DROP TABLE IF EXISTS player_test_results;
CREATE TABLE player_test_results (
    playerId INTEGER NOT NULL,
    testResultId INTEGER NOT NULL,
    waveDied UNSIGNED SMALLINT NOT NULL check(waveDied > 0),
    diedTo INTEGER NOT NULL,
    FOREIGN KEY (playerId) REFERENCES players (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id),
    FOREIGN KEY (diedTo) REFERENCES catastrophes (id)
);

DROP TABLE IF EXISTS catastrophes;
CREATE TABLE catastrophes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name varchar(255) NOT NULL
);
