-- name: GetMostRecentStats :many
SELECT Avg.version, Avg.avgWave, Max.maxWave, Count.numOfTestEvents, StartDate.startDate, EndDate.endDate
    FROM (
    SELECT test_events.version, AVG(player_test_results.waveDied) as avgWave
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Avg, (
    SELECT test_events.version, MAX(player_test_results.waveDied) as maxWave
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Max, (
    SELECT test_events.version, COUNT(test_events.id) as numOfTestEvents
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Count, (
    SELECT test_events.version, MIN(test_events.startedAt) as startDate
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as StartDate, (
    SELECT test_events.version, MAX(test_results.endedAt) as endDate
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    JOIN test_results ON test_events.testResultId = test_results.id
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as EndDate
WHERE Avg.version = Max.version
AND Max.version = Count.version 
AND Count.version = StartDate.version 
AND StartDate.version = EndDate.version;

-- name: GetStatsByVersion :one
SELECT Avg.version, Avg.avgWave, Max.maxWave, Count.numOfTestEvents, StartDate.startDate, EndDate.endDate
FROM (
    SELECT test_events.version, AVG(player_test_results.waveDied) as avgWave
    FROM test_events
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Avg, (
    SELECT test_events.version, MAX(player_test_results.waveDied) as maxWave
    FROM test_events
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Max, (
    SELECT test_events.version, COUNT(test_events.id) as numOfTestEvents
    FROM test_events
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Count, (
    SELECT test_events.version, MIN(test_events.startedAt) as startDate
    FROM test_events
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as StartDate, (
    SELECT test_events.version, MAX(test_events.endedAt) as endDate
    FROM test_events
    JOIN test_results ON test_events.testResultId = test_result.id
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as EndDate
WHERE Avg.version = Max.version 
AND Max.version = Count.version
AND Count.version = StartDate.version
AND StartDate.version = EndDate.version;
