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
    templateId varchar(255) NOT NULL,
    testResultId varchar(255),
    startedAt DATETIME NOT NULL,
    FOREIGN KEY (templateId) REFERENCES templates (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);

CREATE TABLE users(
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE TABLE player_test_results (
    playerId varchar(255) NOT NULL,
    testResultId varchar(255) NOT NULL,
    waveDied INTEGER NOT NULL check(waveDied > 0),
    diedTo varchar(255) NOT NULL,
    PRIMARY KEY (playerId, testResultId),
    FOREIGN KEY (playerId) REFERENCES players (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id)
);

CREATE TABLE components (
    name varchar(255) PRIMARY KEY,
    type varchar(255) NOT NULL
);

CREATE TABLE player_components (
    playerId varchar(255),
    component varchar(255),
    count INTEGER NOT NULL check(count >= 0),
    PRIMARY KEY (playerId, component),
    FOREIGN KEY (playerId) REFERENCES player (id),
    FOREIGN KEY (component) REFERENCES components (name)
);
