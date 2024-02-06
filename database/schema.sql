CREATE TABLE players (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE templates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    playerId varchar(255) NOT NULL,
    data TEXT NOT NULL,
    name varchar(255) NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (playerId) REFERENCES players (id)
); 

CREATE TABLE test_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    waveSurvived INTEGER  NOT NULL check(waveSurvived > 0),
    moneyEarned DECIMAL(14, 2) NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE test_events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    environment varchar(255) NOT NULL,
    difficulty varchar(255) NOT NULL,
    templateId INTEGER NOT NULL,
    testResultId INTEGER,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (templateId) REFERENCES templates (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);

CREATE TABLE test_event_catastrophes (
    testEventId INTEGER NOT NULL,
    catastrophe varchar(255) NOT NULL,
    wave INTEGER NOT NULL check(wave > 0),
    FOREIGN KEY (testEventId) REFERENCES test_events (id)
);

CREATE TABLE test_event_players (
    testEventId INTEGER NOT NULL,
    playerId varchar(255) NOT NULL,
    FOREIGN KEY (testEventId) REFERENCES test_events (id),
    FOREIGN KEY (playerId) REFERENCES players (id)
);

CREATE TABLE users(
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE TABLE player_test_results (
    playerId varchar(255) NOT NULL,
    testResultId INTEGER NOT NULL,
    waveDied INTEGER NOT NULL check(waveDied > 0),
    diedTo varchar(255) NOT NULL,
    FOREIGN KEY (playerId) REFERENCES players (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);
