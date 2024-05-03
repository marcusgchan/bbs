CREATE TABLE players (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
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

CREATE TABLE versions (
    value varchar(10) PRIMARY KEY,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE test_events (
    id varchar(255) PRIMARY KEY,
    environment varchar(255) NOT NULL,
    difficulty varchar(255) NOT NULL,
    templateId varchar(255) NOT NULL,
    testResultId INTEGER,
    startedAt DATETIME NOT NULL,
    version varchar(10) NOT NULL,
    FOREIGN KEY (templateId) REFERENCES templates (id),
    FOREIGN KEY (testResultId) REFERENCES test_results (id),
    FOREIGN KEY (version) REFERENCES versions (value)
);

CREATE TABLE player_test_events (
    playerId varchar(255) NOT NULL,
    testEventId varchar(255) NOT NULL,
    FOREIGN KEY (playerId) REFERENCES players (id),
    FOREIGN KEY (testEventId) REFERENCES test_events (id),
    PRIMARY KEY (playerId, testEventId)
);

CREATE TABLE users(
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE TABLE player_test_results (
    playerId varchar(255) NOT NULL,
    testResultId INTEGER NOT NULL,
    wavesSurvived INTEGER NOT NULL check(wavesSurvived >= 0),
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
