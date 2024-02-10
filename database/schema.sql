CREATE TABLE players (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE templates (
    id varchar(255) PRIMARY KEY,
    playerId varchar(255) NOT NULL,
    data TEXT NOT NULL,
    name varchar(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    FOREIGN KEY (playerId) REFERENCES players (id)
); 

CREATE TABLE test_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    moneyEarned INTEGER NOT NULL,
    endedAt DATETIME NOT NULL
);

CREATE TABLE test_events (
    id varchar(255) PRIMARY KEY,
    environment varchar(255) NOT NULL,
    difficulty varchar(255) NOT NULL,
    templateId INTEGER NOT NULL,
    testResultId INTEGER,
    startedAt DATETIME NOT NULL,
    FOREIGN KEY (templateId) REFERENCES templates (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);

CREATE TABLE test_event_catastrophes (
    testEventId varchar(255) NOT NULL,
    catastrophe varchar(255) NOT NULL,
    wave INTEGER NOT NULL check(wave > 0),
    FOREIGN KEY (testEventId) REFERENCES test_events (id)
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
    PRIMARY KEY (playerId, testResultId),
    FOREIGN KEY (playerId) REFERENCES players (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);
