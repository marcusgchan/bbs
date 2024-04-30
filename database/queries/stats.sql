-- name: GetMostRecentStats :many
SELECT AvgWave.version, AvgWave.avgWave, AvgMoney.avgMoneyEarned, MaxWave.maxWave, Count.numOfTestEvents, StartDate.startDate, EndDate.endDate
    FROM (
    SELECT test_events.version, CAST(AVG(player_test_results.wavesSurvived) as REAL) as avgWave
    FROM (
        SELECT DISTINCT value as version FROM versions
        ORDER BY value DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    INNER JOIN test_results ON test_events.testResultId = test_results.id
    JOIN player_test_results ON test_events.testResultId = player_test_results.testresultId
    GROUP BY test_events.version
) as AvgWave, (
    SELECT test_events.version, CAST(MAX(player_test_results.wavesSurvived) AS INTEGER) as maxWave
    FROM (
        SELECT DISTINCT value as version FROM versions
        ORDER BY value DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    INNER JOIN test_results ON test_events.testResultId = test_results.id
    JOIN player_test_results ON test_events.testResultId = player_test_results.testresultId
    GROUP BY test_events.version
) as MaxWave, (
    SELECT test_events.version, CAST(AVG(test_results.moneyEarned) as REAL) as avgMoneyEarned
    FROM (
        SELECT DISTINCT value as version 
        FROM versions
        ORDER BY value DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    INNER JOIN test_results ON test_events.testResultId = test_results.id
    GROUP BY test_events.version
) as AvgMoney, (
    SELECT test_events.version, COUNT(test_events.id) as numOfTestEvents
    FROM (
        SELECT DISTINCT value as version 
        FROM versions
        ORDER BY value DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Count, (
    SELECT test_events.version, CAST(MIN(test_events.startedAt) AS TEXT) as startDate
    FROM (
        SELECT DISTINCT value as version FROM versions
        ORDER BY value DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as StartDate, (
    SELECT test_events.version, CAST(MAX(test_results.endedAt) AS TEXT) as endDate
    FROM (
        SELECT DISTINCT value as version FROM versions
        ORDER BY value DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    INNER JOIN test_results ON test_events.testResultId = test_results.id
    GROUP BY test_events.version
) as EndDate
WHERE AvgWave.version = MaxWave.version
AND MaxWave.version = Count.version 
AND AvgMoney.version = Count.version
AND Count.version = StartDate.version 
AND StartDate.version = EndDate.version
ORDER BY AvgWave.version DESC;

-- name: GetStatsByVersion :one
SELECT AvgWave.version, AvgWave.avgWave, AvgMoney.avgMoneyEarned, MaxWave.maxWave, Count.numOfTestEvents, StartDate.startDate, EndDate.endDate
FROM (
    SELECT test_events.version, CAST(AVG(player_test_results.wavesSurvived) as REAL) as avgWave
    FROM test_events
    INNER JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.version = ?
) as AvgWave, (
    SELECT test_events.version, CAST(AVG(test_results.moneyEarned) as REAL) as avgMoneyEarned
    FROM test_events
    INNER JOIN test_results ON test_events.testResultId = test_results.id
    WHERE test_events.version = ?
) as AvgMoney, (
    SELECT test_events.version, CAST(MAX(player_test_results.wavesSurvived) AS INTEGER) as maxWave
    FROM test_events
    INNER JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.version = ?
) as MaxWave, (
    SELECT test_events.version, COUNT(test_events.id) as numOfTestEvents
    FROM test_events
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Count, (
    SELECT test_events.version, CAST(MIN(test_events.startedAt) AS TEXT) as startDate
    FROM test_events
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as StartDate, (
    SELECT test_events.version, CAST(MAX(test_results.endedAt) AS TEXT) as endDate
    FROM test_events
    INNER JOIN test_results ON test_events.testResultId = test_results.id
    WHERE test_events.version = ?
) as EndDate
WHERE AvgWave.version = MaxWave.version 
AND MaxWave.version = Count.version
AND AvgMoney.version = Count.version
AND Count.version = StartDate.version
AND StartDate.version = EndDate.version;
