(function() {
    var players = {};
    var standings = [];
    var rounds = [];

    var WIN_POINTS = 3;
    var DRAW_POINTS = 1;
    var LOSE_POINTS = 0;

    var addPlayer = function(name) {
        players[name] = { name: name };
        standings.push(players[name]);
    };

    var addResults = function(roundIndex, player1, score1, player2, score2) {
        if (!rounds[roundIndex]) {
            rounds[roundIndex] = [];
        }

        rounds[roundIndex].push([
            { name: player1, score: score1 },
            { name: player2, score: score2 }
        ]);
    };

    var resetPlayers = function() {
        for (var index = 0; index < standings.length; ++index) {
            var player = standings[index];

            player.win = 0;
            player.points = 0;
            player.totalScore = 0;
            player.opponents = [];

            // Used to calculate the favored player
            player.lastTwoScore = 0;
        }
    };

    var calculateMatchResult = function(match, addLastTwo) {
        var player1 = players[match[0].name];
        var player2 = players[match[1].name];

        player1.opponents.push(player2);
        player2.opponents.push(player1);

        player1.totalScore += match[0].score;
        player2.totalScore += match[1].score;

        if (addLastTwo) {
            player1.lastTwoScore += match[0].score;
            player2.lastTwoScore += match[1].score;
        }

        if (match[0].score > match[1].score) {
            player1.points += WIN_POINTS;
            player2.points += LOSE_POINTS;
            player1.win += 1;
        } else if (match[0].score < match[1].score) {
            player1.points += LOSE_POINTS;
            player2.points += WIN_POINTS;
            player2.win += 1;
        } else {
            player1.points += DRAW_POINTS;
            player2.points += DRAW_POINTS;
        }
    };

    var calculateAllMatchResults = function() {
        for (var roundIndex = 0; roundIndex < rounds.length; ++roundIndex) {
            var round = rounds[roundIndex];
            var addLastTwo = (rounds.length - (roundIndex + 2)) <= 0;

            for (var matchIndex = 0; matchIndex < round.length; ++matchIndex) {
                calculateMatchResult(round[matchIndex], addLastTwo);
            }
        }
    };

    var calculateOpponentMatchWins = function() {
        var roundsLength = rounds.length * 1.0;
        for (var sIndex = 0; sIndex < standings.length; ++sIndex) {
            var player = standings[sIndex];
            var omw = 0;

            var opponents = player.opponents;

            // Calculate opponent match wins
            for (var oIndex = 0; oIndex < opponents.length; ++oIndex) {
                omw += opponents[oIndex].win;
            }

            var totalGamesPlayedByOpponents = roundsLength * roundsLength;
            player.omw = Math.round(omw * 100.0 / totalGamesPlayedByOpponents);
        }
    };

    var sortStandings = function() {
        standings.sort(function(player1, player2) {
            // Highest first
            var result = player2.points - player1.points;

            // Tie breaker 1 - oponent wins
            if (result === 0) {
                result = player2.omw - player1.omw;
            }

            // Tie breaker 2 - total score
            if (result === 0) {
                result = player2.totalScore - player1.totalScore;
            }

            // Tie breaker 3 - flip the coin!
            if (result === 0) {
                result = Math.round(Math.random()) === 1 ? 1 : -1;
            }

            return result;
        });
    };

    var calculateStandings = function() {
        resetPlayers();
        calculateAllMatchResults();
        calculateOpponentMatchWins();
        sortStandings();

        console.log('Current standings');
        console.log('-----------------');
        for (var index = 0; index < standings.length; ++index) {
            var player = standings[index];

            console.log(
                (index + 1) + '. ' +
                player.name + ' (' +
                'matchPoints: ' + player.points +
                ' opponentWin: ' + player.omw + '%' +
                ' totalScore: ' + player.totalScore +
                ')');
        }

        console.log('\nThis coming week');
        console.log('-----------------');
        for (var index = 0; index < standings.length; index += 2) {
            var player1 = standings[index];
            var player2 = standings[index + 1];

            if (!player2) {
                console.log(player1.name + ' bye week');
            }

            var favoredName = player1.name;
            var favoredAmount = player1.lastTwoScore - player2.lastTwoScore;
            favoredAmount /= 2;
            if (player2.lastTwoScore > player1.lastTwoScore) {
                favoredName = player2.name;
                favoredAmount *= -1;
            }

            console.log(
                player1.name + ' vs ' + player2.name + ' favored: ' +
                favoredName + ' by ' + favoredAmount);
        }
    };

    // -------------------------------------------------------------------------

    // Add the players
    addPlayer('Marcio');
    addPlayer('Patrick');
    addPlayer('Paul');
    addPlayer('Eliah');
    addPlayer('John');
    addPlayer('Dan');

    // Games
    addResults(0, 'Marcio', 79, 'Patrick', 26);
    addResults(0, 'Paul', 58, 'Eliah', 38);
    addResults(0, 'John', 45, 'Dan', 19);

    addResults(1, 'Paul', 68, 'Marcio', 67);
    addResults(1, 'Eliah', 79, 'John', 16);
    addResults(1, 'Patrick', 49, 'Dan', 36);

    addResults(2, 'Paul', 56, 'John', 52);
    addResults(2, 'Marcio', 68, 'Dan', 41);
    addResults(2, 'Eliah', 75, 'Patrick', 44);
    
    addResults(3, 'Paul', 56, 'Eliah', 28);
    addResults(3, 'Marcio', 93, 'John', 78);
    addResults(3, 'Dan', 62, 'Patrick', 28);

    addResults(4, 'Paul', 68, 'Marcio', 60);
    addResults(4, 'Eliah', 116, 'John', 47);
    addResults(4, 'Patrick', 88, 'Dan', 47);

    addResults(5, 'Marcio', 92, 'Paul', 51);
    addResults(5, 'Eliah', 89, 'Patrick', 56);
    addResults(5, 'John', 77, 'Dan', 32);

    addResults(6, 'Paul', 32, 'Marcio', 62);
    addResults(6, 'Eliah', 53, 'John', 44);
    addResults(6, 'Patrick', 40, 'Dan', 42);

    calculateStandings();
})();