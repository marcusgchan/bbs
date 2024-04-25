-- name: GetMostRecentStats :many
SELECT Avg.version, Avg.avgWave, Max.maxWave, Count.numOfTestEvents
FROM (
    SELECT test_events.version, AVG(player_test_results.waveDied) as avgWave
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as V
    JOIN test_events ON test_events.version = V.version
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
    ) as V2
    JOIN test_events ON test_events.version = V2.version
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
    ) as V3
    JOIN test_events ON test_events.version = V3.version
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Count
WHERE Avg.version = Max.version AND Max.version = Count.version
